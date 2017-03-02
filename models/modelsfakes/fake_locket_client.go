// This file was generated by counterfeiter
package modelsfakes

import (
	"sync"

	"code.cloudfoundry.org/locket/models"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type FakeLocketClient struct {
	LockStub        func(ctx context.Context, in *models.LockRequest, opts ...grpc.CallOption) (*models.LockResponse, error)
	lockMutex       sync.RWMutex
	lockArgsForCall []struct {
		ctx  context.Context
		in   *models.LockRequest
		opts []grpc.CallOption
	}
	lockReturns struct {
		result1 *models.LockResponse
		result2 error
	}
	FetchStub        func(ctx context.Context, in *models.FetchRequest, opts ...grpc.CallOption) (*models.FetchResponse, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		ctx  context.Context
		in   *models.FetchRequest
		opts []grpc.CallOption
	}
	fetchReturns struct {
		result1 *models.FetchResponse
		result2 error
	}
	ReleaseStub        func(ctx context.Context, in *models.ReleaseRequest, opts ...grpc.CallOption) (*models.ReleaseResponse, error)
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct {
		ctx  context.Context
		in   *models.ReleaseRequest
		opts []grpc.CallOption
	}
	releaseReturns struct {
		result1 *models.ReleaseResponse
		result2 error
	}
	FetchAllStub        func(ctx context.Context, in *models.FetchAllRequest, opts ...grpc.CallOption) (*models.FetchAllResponse, error)
	fetchAllMutex       sync.RWMutex
	fetchAllArgsForCall []struct {
		ctx  context.Context
		in   *models.FetchAllRequest
		opts []grpc.CallOption
	}
	fetchAllReturns struct {
		result1 *models.FetchAllResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLocketClient) Lock(ctx context.Context, in *models.LockRequest, opts ...grpc.CallOption) (*models.LockResponse, error) {
	fake.lockMutex.Lock()
	fake.lockArgsForCall = append(fake.lockArgsForCall, struct {
		ctx  context.Context
		in   *models.LockRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("Lock", []interface{}{ctx, in, opts})
	fake.lockMutex.Unlock()
	if fake.LockStub != nil {
		return fake.LockStub(ctx, in, opts...)
	} else {
		return fake.lockReturns.result1, fake.lockReturns.result2
	}
}

func (fake *FakeLocketClient) LockCallCount() int {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return len(fake.lockArgsForCall)
}

func (fake *FakeLocketClient) LockArgsForCall(i int) (context.Context, *models.LockRequest, []grpc.CallOption) {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return fake.lockArgsForCall[i].ctx, fake.lockArgsForCall[i].in, fake.lockArgsForCall[i].opts
}

func (fake *FakeLocketClient) LockReturns(result1 *models.LockResponse, result2 error) {
	fake.LockStub = nil
	fake.lockReturns = struct {
		result1 *models.LockResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeLocketClient) Fetch(ctx context.Context, in *models.FetchRequest, opts ...grpc.CallOption) (*models.FetchResponse, error) {
	fake.fetchMutex.Lock()
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		ctx  context.Context
		in   *models.FetchRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("Fetch", []interface{}{ctx, in, opts})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub(ctx, in, opts...)
	} else {
		return fake.fetchReturns.result1, fake.fetchReturns.result2
	}
}

func (fake *FakeLocketClient) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeLocketClient) FetchArgsForCall(i int) (context.Context, *models.FetchRequest, []grpc.CallOption) {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return fake.fetchArgsForCall[i].ctx, fake.fetchArgsForCall[i].in, fake.fetchArgsForCall[i].opts
}

func (fake *FakeLocketClient) FetchReturns(result1 *models.FetchResponse, result2 error) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 *models.FetchResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeLocketClient) Release(ctx context.Context, in *models.ReleaseRequest, opts ...grpc.CallOption) (*models.ReleaseResponse, error) {
	fake.releaseMutex.Lock()
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct {
		ctx  context.Context
		in   *models.ReleaseRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("Release", []interface{}{ctx, in, opts})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		return fake.ReleaseStub(ctx, in, opts...)
	} else {
		return fake.releaseReturns.result1, fake.releaseReturns.result2
	}
}

func (fake *FakeLocketClient) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakeLocketClient) ReleaseArgsForCall(i int) (context.Context, *models.ReleaseRequest, []grpc.CallOption) {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return fake.releaseArgsForCall[i].ctx, fake.releaseArgsForCall[i].in, fake.releaseArgsForCall[i].opts
}

func (fake *FakeLocketClient) ReleaseReturns(result1 *models.ReleaseResponse, result2 error) {
	fake.ReleaseStub = nil
	fake.releaseReturns = struct {
		result1 *models.ReleaseResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeLocketClient) FetchAll(ctx context.Context, in *models.FetchAllRequest, opts ...grpc.CallOption) (*models.FetchAllResponse, error) {
	fake.fetchAllMutex.Lock()
	fake.fetchAllArgsForCall = append(fake.fetchAllArgsForCall, struct {
		ctx  context.Context
		in   *models.FetchAllRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("FetchAll", []interface{}{ctx, in, opts})
	fake.fetchAllMutex.Unlock()
	if fake.FetchAllStub != nil {
		return fake.FetchAllStub(ctx, in, opts...)
	} else {
		return fake.fetchAllReturns.result1, fake.fetchAllReturns.result2
	}
}

func (fake *FakeLocketClient) FetchAllCallCount() int {
	fake.fetchAllMutex.RLock()
	defer fake.fetchAllMutex.RUnlock()
	return len(fake.fetchAllArgsForCall)
}

func (fake *FakeLocketClient) FetchAllArgsForCall(i int) (context.Context, *models.FetchAllRequest, []grpc.CallOption) {
	fake.fetchAllMutex.RLock()
	defer fake.fetchAllMutex.RUnlock()
	return fake.fetchAllArgsForCall[i].ctx, fake.fetchAllArgsForCall[i].in, fake.fetchAllArgsForCall[i].opts
}

func (fake *FakeLocketClient) FetchAllReturns(result1 *models.FetchAllResponse, result2 error) {
	fake.FetchAllStub = nil
	fake.fetchAllReturns = struct {
		result1 *models.FetchAllResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeLocketClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	fake.fetchAllMutex.RLock()
	defer fake.fetchAllMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLocketClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ models.LocketClient = new(FakeLocketClient)
