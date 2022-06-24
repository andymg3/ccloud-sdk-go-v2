// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_scopes_service_quota_v1.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1 "github.com/confluentinc/ccloud-sdk-go-v2/service-quota/v1"
)

// ScopesServiceQuotaV1Api is a mock of ScopesServiceQuotaV1Api interface
type ScopesServiceQuotaV1Api struct {
	lockGetServiceQuotaV1Scope sync.Mutex
	GetServiceQuotaV1ScopeFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiGetServiceQuotaV1ScopeRequest

	lockGetServiceQuotaV1ScopeExecute sync.Mutex
	GetServiceQuotaV1ScopeExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiGetServiceQuotaV1ScopeRequest) (github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ServiceQuotaV1Scope, *net_http.Response, error)

	lockListServiceQuotaV1Scopes sync.Mutex
	ListServiceQuotaV1ScopesFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiListServiceQuotaV1ScopesRequest

	lockListServiceQuotaV1ScopesExecute sync.Mutex
	ListServiceQuotaV1ScopesExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiListServiceQuotaV1ScopesRequest) (github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ServiceQuotaV1ScopeList, *net_http.Response, error)

	calls struct {
		GetServiceQuotaV1Scope []struct {
			Ctx context.Context
			Id  string
		}
		GetServiceQuotaV1ScopeExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiGetServiceQuotaV1ScopeRequest
		}
		ListServiceQuotaV1Scopes []struct {
			Ctx context.Context
		}
		ListServiceQuotaV1ScopesExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiListServiceQuotaV1ScopesRequest
		}
	}
}

// GetServiceQuotaV1Scope mocks base method by wrapping the associated func.
func (m *ScopesServiceQuotaV1Api) GetServiceQuotaV1Scope(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiGetServiceQuotaV1ScopeRequest {
	m.lockGetServiceQuotaV1Scope.Lock()
	defer m.lockGetServiceQuotaV1Scope.Unlock()

	if m.GetServiceQuotaV1ScopeFunc == nil {
		panic("mocker: ScopesServiceQuotaV1Api.GetServiceQuotaV1ScopeFunc is nil but ScopesServiceQuotaV1Api.GetServiceQuotaV1Scope was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.GetServiceQuotaV1Scope = append(m.calls.GetServiceQuotaV1Scope, call)

	return m.GetServiceQuotaV1ScopeFunc(ctx, id)
}

// GetServiceQuotaV1ScopeCalled returns true if GetServiceQuotaV1Scope was called at least once.
func (m *ScopesServiceQuotaV1Api) GetServiceQuotaV1ScopeCalled() bool {
	m.lockGetServiceQuotaV1Scope.Lock()
	defer m.lockGetServiceQuotaV1Scope.Unlock()

	return len(m.calls.GetServiceQuotaV1Scope) > 0
}

// GetServiceQuotaV1ScopeCalls returns the calls made to GetServiceQuotaV1Scope.
func (m *ScopesServiceQuotaV1Api) GetServiceQuotaV1ScopeCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockGetServiceQuotaV1Scope.Lock()
	defer m.lockGetServiceQuotaV1Scope.Unlock()

	return m.calls.GetServiceQuotaV1Scope
}

// GetServiceQuotaV1ScopeExecute mocks base method by wrapping the associated func.
func (m *ScopesServiceQuotaV1Api) GetServiceQuotaV1ScopeExecute(r github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiGetServiceQuotaV1ScopeRequest) (github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ServiceQuotaV1Scope, *net_http.Response, error) {
	m.lockGetServiceQuotaV1ScopeExecute.Lock()
	defer m.lockGetServiceQuotaV1ScopeExecute.Unlock()

	if m.GetServiceQuotaV1ScopeExecuteFunc == nil {
		panic("mocker: ScopesServiceQuotaV1Api.GetServiceQuotaV1ScopeExecuteFunc is nil but ScopesServiceQuotaV1Api.GetServiceQuotaV1ScopeExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiGetServiceQuotaV1ScopeRequest
	}{
		R: r,
	}

	m.calls.GetServiceQuotaV1ScopeExecute = append(m.calls.GetServiceQuotaV1ScopeExecute, call)

	return m.GetServiceQuotaV1ScopeExecuteFunc(r)
}

// GetServiceQuotaV1ScopeExecuteCalled returns true if GetServiceQuotaV1ScopeExecute was called at least once.
func (m *ScopesServiceQuotaV1Api) GetServiceQuotaV1ScopeExecuteCalled() bool {
	m.lockGetServiceQuotaV1ScopeExecute.Lock()
	defer m.lockGetServiceQuotaV1ScopeExecute.Unlock()

	return len(m.calls.GetServiceQuotaV1ScopeExecute) > 0
}

// GetServiceQuotaV1ScopeExecuteCalls returns the calls made to GetServiceQuotaV1ScopeExecute.
func (m *ScopesServiceQuotaV1Api) GetServiceQuotaV1ScopeExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiGetServiceQuotaV1ScopeRequest
} {
	m.lockGetServiceQuotaV1ScopeExecute.Lock()
	defer m.lockGetServiceQuotaV1ScopeExecute.Unlock()

	return m.calls.GetServiceQuotaV1ScopeExecute
}

// ListServiceQuotaV1Scopes mocks base method by wrapping the associated func.
func (m *ScopesServiceQuotaV1Api) ListServiceQuotaV1Scopes(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiListServiceQuotaV1ScopesRequest {
	m.lockListServiceQuotaV1Scopes.Lock()
	defer m.lockListServiceQuotaV1Scopes.Unlock()

	if m.ListServiceQuotaV1ScopesFunc == nil {
		panic("mocker: ScopesServiceQuotaV1Api.ListServiceQuotaV1ScopesFunc is nil but ScopesServiceQuotaV1Api.ListServiceQuotaV1Scopes was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.ListServiceQuotaV1Scopes = append(m.calls.ListServiceQuotaV1Scopes, call)

	return m.ListServiceQuotaV1ScopesFunc(ctx)
}

// ListServiceQuotaV1ScopesCalled returns true if ListServiceQuotaV1Scopes was called at least once.
func (m *ScopesServiceQuotaV1Api) ListServiceQuotaV1ScopesCalled() bool {
	m.lockListServiceQuotaV1Scopes.Lock()
	defer m.lockListServiceQuotaV1Scopes.Unlock()

	return len(m.calls.ListServiceQuotaV1Scopes) > 0
}

// ListServiceQuotaV1ScopesCalls returns the calls made to ListServiceQuotaV1Scopes.
func (m *ScopesServiceQuotaV1Api) ListServiceQuotaV1ScopesCalls() []struct {
	Ctx context.Context
} {
	m.lockListServiceQuotaV1Scopes.Lock()
	defer m.lockListServiceQuotaV1Scopes.Unlock()

	return m.calls.ListServiceQuotaV1Scopes
}

// ListServiceQuotaV1ScopesExecute mocks base method by wrapping the associated func.
func (m *ScopesServiceQuotaV1Api) ListServiceQuotaV1ScopesExecute(r github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiListServiceQuotaV1ScopesRequest) (github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ServiceQuotaV1ScopeList, *net_http.Response, error) {
	m.lockListServiceQuotaV1ScopesExecute.Lock()
	defer m.lockListServiceQuotaV1ScopesExecute.Unlock()

	if m.ListServiceQuotaV1ScopesExecuteFunc == nil {
		panic("mocker: ScopesServiceQuotaV1Api.ListServiceQuotaV1ScopesExecuteFunc is nil but ScopesServiceQuotaV1Api.ListServiceQuotaV1ScopesExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiListServiceQuotaV1ScopesRequest
	}{
		R: r,
	}

	m.calls.ListServiceQuotaV1ScopesExecute = append(m.calls.ListServiceQuotaV1ScopesExecute, call)

	return m.ListServiceQuotaV1ScopesExecuteFunc(r)
}

// ListServiceQuotaV1ScopesExecuteCalled returns true if ListServiceQuotaV1ScopesExecute was called at least once.
func (m *ScopesServiceQuotaV1Api) ListServiceQuotaV1ScopesExecuteCalled() bool {
	m.lockListServiceQuotaV1ScopesExecute.Lock()
	defer m.lockListServiceQuotaV1ScopesExecute.Unlock()

	return len(m.calls.ListServiceQuotaV1ScopesExecute) > 0
}

// ListServiceQuotaV1ScopesExecuteCalls returns the calls made to ListServiceQuotaV1ScopesExecute.
func (m *ScopesServiceQuotaV1Api) ListServiceQuotaV1ScopesExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_service_quota_v1.ApiListServiceQuotaV1ScopesRequest
} {
	m.lockListServiceQuotaV1ScopesExecute.Lock()
	defer m.lockListServiceQuotaV1ScopesExecute.Unlock()

	return m.calls.ListServiceQuotaV1ScopesExecute
}

// Reset resets the calls made to the mocked methods.
func (m *ScopesServiceQuotaV1Api) Reset() {
	m.lockGetServiceQuotaV1Scope.Lock()
	m.calls.GetServiceQuotaV1Scope = nil
	m.lockGetServiceQuotaV1Scope.Unlock()
	m.lockGetServiceQuotaV1ScopeExecute.Lock()
	m.calls.GetServiceQuotaV1ScopeExecute = nil
	m.lockGetServiceQuotaV1ScopeExecute.Unlock()
	m.lockListServiceQuotaV1Scopes.Lock()
	m.calls.ListServiceQuotaV1Scopes = nil
	m.lockListServiceQuotaV1Scopes.Unlock()
	m.lockListServiceQuotaV1ScopesExecute.Lock()
	m.calls.ListServiceQuotaV1ScopesExecute = nil
	m.lockListServiceQuotaV1ScopesExecute.Unlock()
}
