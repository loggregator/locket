package handlers

import (
	"time"

	"code.cloudfoundry.org/bbs/db/sqldb/helpers"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/locket/db"
	"code.cloudfoundry.org/locket/expiration"
	"code.cloudfoundry.org/locket/metrics"
	"code.cloudfoundry.org/locket/models"
	"golang.org/x/net/context"
)

type locketHandler struct {
	logger lager.Logger

	db       db.LockDB
	exitCh   chan<- struct{}
	lockPick expiration.LockPick
	metrics  metrics.RequestMetrics
}

func NewLocketHandler(logger lager.Logger, db db.LockDB, lockPick expiration.LockPick, requestMetrics metrics.RequestMetrics, exitCh chan<- struct{}) *locketHandler {
	return &locketHandler{
		logger:   logger,
		db:       db,
		lockPick: lockPick,
		exitCh:   exitCh,
		metrics:  requestMetrics,
	}
}

func (h *locketHandler) exitIfUnrecoverable(err error) {
	if err != helpers.ErrUnrecoverableError {
		return
	}

	h.logger.Error("unrecoverable-error", err)

	select {
	case h.exitCh <- struct{}{}:
	default:
	}
}

func (h *locketHandler) monitorRequest(requestType string, f func() error) error {

	h.metrics.IncrementRequestsStartedCounter(requestType, 1)
	h.metrics.IncrementRequestsInFlightCounter(requestType, 1)
	defer h.metrics.DecrementRequestsInFlightCounter(requestType, 1)

	start := time.Now()

	err := f()

	h.metrics.UpdateLatency(requestType, time.Since(start))

	if err != nil && err != models.ErrLockCollision {
		h.metrics.IncrementRequestsFailedCounter(requestType, 1)
		h.exitIfUnrecoverable(err)
	} else {
		h.metrics.IncrementRequestsSucceededCounter(requestType, 1)
	}
	return err
}

func (h *locketHandler) Lock(ctx context.Context, req *models.LockRequest) (*models.LockResponse, error) {
	var (
		response *models.LockResponse
		err      error
	)

	err = h.monitorRequest("Lock", func() error {
		response, err = h.lock(ctx, req)
		return err
	})

	return response, err
}

func (h *locketHandler) Release(ctx context.Context, req *models.ReleaseRequest) (*models.ReleaseResponse, error) {
	var (
		response *models.ReleaseResponse
		err      error
	)

	err = h.monitorRequest("Release", func() error {
		response, err = h.release(ctx, req)
		return err
	})

	return response, err
}

func (h *locketHandler) Fetch(ctx context.Context, req *models.FetchRequest) (*models.FetchResponse, error) {
	var (
		response *models.FetchResponse
		err      error
	)

	err = h.monitorRequest("Fetch", func() error {
		response, err = h.fetch(ctx, req)
		return err
	})

	return response, err
}

func (h *locketHandler) FetchAll(ctx context.Context, req *models.FetchAllRequest) (*models.FetchAllResponse, error) {
	var (
		response *models.FetchAllResponse
		err      error
	)

	err = h.monitorRequest("FetchAll", func() error {
		response, err = h.fetchAll(ctx, req)
		return err
	})

	return response, err
}

func (h *locketHandler) lock(ctx context.Context, req *models.LockRequest) (*models.LockResponse, error) {
	logger := h.logger.Session("lock")
	logger.Debug("started")
	defer logger.Debug("complete")

	err := validate(req)
	if err != nil {
		logger.Error("invalid-request", err, lager.Data{"type": req.Resource.GetType(), "typeCode": req.Resource.GetTypeCode()})

		return nil, err
	}

	if req.TtlInSeconds <= 0 {
		logger.Error("failed-locking-lock", models.ErrInvalidTTL, lager.Data{
			"key":   req.Resource.Key,
			"owner": req.Resource.Owner,
		})
		return nil, models.ErrInvalidTTL
	}

	if req.Resource.Owner == "" {
		logger.Error("failed-locking-lock", models.ErrInvalidOwner, lager.Data{
			"key":   req.Resource.Key,
			"owner": req.Resource.Owner,
		})
		return nil, models.ErrInvalidOwner
	}

	lock, err := h.db.Lock(logger, req.Resource, req.TtlInSeconds)
	if err != nil {
		if err != models.ErrLockCollision {
			logger.Error("failed-locking-lock", err, lager.Data{
				"key":   req.Resource.Key,
				"owner": req.Resource.Owner,
			})
		}
		return nil, err
	}

	h.lockPick.RegisterTTL(logger, lock)

	return &models.LockResponse{}, nil
}

func (h *locketHandler) release(ctx context.Context, req *models.ReleaseRequest) (*models.ReleaseResponse, error) {
	logger := h.logger.Session("release")
	logger.Debug("started")
	defer logger.Debug("complete")

	err := h.db.Release(logger, req.Resource)
	if err != nil {
		return nil, err
	}

	return &models.ReleaseResponse{}, nil
}

func (h *locketHandler) fetch(ctx context.Context, req *models.FetchRequest) (*models.FetchResponse, error) {
	logger := h.logger.Session("fetch")
	logger.Debug("started")
	defer logger.Debug("complete")

	lock, err := h.db.Fetch(logger, req.Key)
	if err != nil {
		return nil, err
	}

	return &models.FetchResponse{
		Resource: lock.Resource,
	}, nil
}

func (h *locketHandler) fetchAll(ctx context.Context, req *models.FetchAllRequest) (*models.FetchAllResponse, error) {
	logger := h.logger.Session("fetch-all")
	logger.Debug("started")
	defer logger.Debug("complete")

	err := validate(req)
	if err != nil {
		logger.Error("invalid-request", err, lager.Data{"type": req.GetType(), "typeCode": req.GetTypeCode()})
		return nil, err
	}

	locks, err := h.db.FetchAll(logger, models.GetType(&models.Resource{Type: req.Type, TypeCode: req.TypeCode}))
	if err != nil {
		return nil, err
	}

	var responses []*models.Resource
	for _, lock := range locks {
		responses = append(responses, lock.Resource)
	}

	return &models.FetchAllResponse{
		Resources: responses,
	}, nil
}

func validate(req interface{}) error {
	var reqType string
	var reqTypeCode models.TypeCode

	switch incomingReq := req.(type) {
	case *models.LockRequest:
		reqType = incomingReq.Resource.GetType()
		reqTypeCode = incomingReq.Resource.GetTypeCode()
	case *models.FetchAllRequest:
		reqType = incomingReq.GetType()
		reqTypeCode = incomingReq.GetTypeCode()
	default:
		return nil
	}

	if _, found := models.TypeCode_name[int32(reqTypeCode)]; !found {
		return models.ErrInvalidType
	}

	if reqTypeCode == models.UNKNOWN {
		if reqType != models.PresenceType && reqType != models.LockType {
			return models.ErrInvalidType
		} else {
			return nil
		}
	}

	if reqType != "" {
		if reqTypeCode == models.LOCK && reqType != models.LockType {
			return models.ErrInvalidType
		}

		if reqTypeCode == models.PRESENCE && reqType != models.PresenceType {
			return models.ErrInvalidType
		}
	}

	return nil
}
