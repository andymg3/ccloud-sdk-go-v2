// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: api_regions_srcm_v2.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2 "github.com/confluentinc/ccloud-sdk-go-v2/srcm/v2"
)

// RegionsSrcmV2Api is a mock of RegionsSrcmV2Api interface
type RegionsSrcmV2Api struct {
	lockGetSrcmV2Region sync.Mutex
	GetSrcmV2RegionFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiGetSrcmV2RegionRequest

	lockGetSrcmV2RegionExecute sync.Mutex
	GetSrcmV2RegionExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiGetSrcmV2RegionRequest) (github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.SrcmV2Region, *net_http.Response, error)

	lockListSrcmV2Regions sync.Mutex
	ListSrcmV2RegionsFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiListSrcmV2RegionsRequest

	lockListSrcmV2RegionsExecute sync.Mutex
	ListSrcmV2RegionsExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiListSrcmV2RegionsRequest) (github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.SrcmV2RegionList, *net_http.Response, error)

	calls struct {
		GetSrcmV2Region []struct {
			Ctx context.Context
			Id  string
		}
		GetSrcmV2RegionExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiGetSrcmV2RegionRequest
		}
		ListSrcmV2Regions []struct {
			Ctx context.Context
		}
		ListSrcmV2RegionsExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiListSrcmV2RegionsRequest
		}
	}
}

// GetSrcmV2Region mocks base method by wrapping the associated func.
func (m *RegionsSrcmV2Api) GetSrcmV2Region(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiGetSrcmV2RegionRequest {
	m.lockGetSrcmV2Region.Lock()
	defer m.lockGetSrcmV2Region.Unlock()

	if m.GetSrcmV2RegionFunc == nil {
		panic("mocker: RegionsSrcmV2Api.GetSrcmV2RegionFunc is nil but RegionsSrcmV2Api.GetSrcmV2Region was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.GetSrcmV2Region = append(m.calls.GetSrcmV2Region, call)

	return m.GetSrcmV2RegionFunc(ctx, id)
}

// GetSrcmV2RegionCalled returns true if GetSrcmV2Region was called at least once.
func (m *RegionsSrcmV2Api) GetSrcmV2RegionCalled() bool {
	m.lockGetSrcmV2Region.Lock()
	defer m.lockGetSrcmV2Region.Unlock()

	return len(m.calls.GetSrcmV2Region) > 0
}

// GetSrcmV2RegionCalls returns the calls made to GetSrcmV2Region.
func (m *RegionsSrcmV2Api) GetSrcmV2RegionCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockGetSrcmV2Region.Lock()
	defer m.lockGetSrcmV2Region.Unlock()

	return m.calls.GetSrcmV2Region
}

// GetSrcmV2RegionExecute mocks base method by wrapping the associated func.
func (m *RegionsSrcmV2Api) GetSrcmV2RegionExecute(r github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiGetSrcmV2RegionRequest) (github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.SrcmV2Region, *net_http.Response, error) {
	m.lockGetSrcmV2RegionExecute.Lock()
	defer m.lockGetSrcmV2RegionExecute.Unlock()

	if m.GetSrcmV2RegionExecuteFunc == nil {
		panic("mocker: RegionsSrcmV2Api.GetSrcmV2RegionExecuteFunc is nil but RegionsSrcmV2Api.GetSrcmV2RegionExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiGetSrcmV2RegionRequest
	}{
		R: r,
	}

	m.calls.GetSrcmV2RegionExecute = append(m.calls.GetSrcmV2RegionExecute, call)

	return m.GetSrcmV2RegionExecuteFunc(r)
}

// GetSrcmV2RegionExecuteCalled returns true if GetSrcmV2RegionExecute was called at least once.
func (m *RegionsSrcmV2Api) GetSrcmV2RegionExecuteCalled() bool {
	m.lockGetSrcmV2RegionExecute.Lock()
	defer m.lockGetSrcmV2RegionExecute.Unlock()

	return len(m.calls.GetSrcmV2RegionExecute) > 0
}

// GetSrcmV2RegionExecuteCalls returns the calls made to GetSrcmV2RegionExecute.
func (m *RegionsSrcmV2Api) GetSrcmV2RegionExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiGetSrcmV2RegionRequest
} {
	m.lockGetSrcmV2RegionExecute.Lock()
	defer m.lockGetSrcmV2RegionExecute.Unlock()

	return m.calls.GetSrcmV2RegionExecute
}

// ListSrcmV2Regions mocks base method by wrapping the associated func.
func (m *RegionsSrcmV2Api) ListSrcmV2Regions(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiListSrcmV2RegionsRequest {
	m.lockListSrcmV2Regions.Lock()
	defer m.lockListSrcmV2Regions.Unlock()

	if m.ListSrcmV2RegionsFunc == nil {
		panic("mocker: RegionsSrcmV2Api.ListSrcmV2RegionsFunc is nil but RegionsSrcmV2Api.ListSrcmV2Regions was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.ListSrcmV2Regions = append(m.calls.ListSrcmV2Regions, call)

	return m.ListSrcmV2RegionsFunc(ctx)
}

// ListSrcmV2RegionsCalled returns true if ListSrcmV2Regions was called at least once.
func (m *RegionsSrcmV2Api) ListSrcmV2RegionsCalled() bool {
	m.lockListSrcmV2Regions.Lock()
	defer m.lockListSrcmV2Regions.Unlock()

	return len(m.calls.ListSrcmV2Regions) > 0
}

// ListSrcmV2RegionsCalls returns the calls made to ListSrcmV2Regions.
func (m *RegionsSrcmV2Api) ListSrcmV2RegionsCalls() []struct {
	Ctx context.Context
} {
	m.lockListSrcmV2Regions.Lock()
	defer m.lockListSrcmV2Regions.Unlock()

	return m.calls.ListSrcmV2Regions
}

// ListSrcmV2RegionsExecute mocks base method by wrapping the associated func.
func (m *RegionsSrcmV2Api) ListSrcmV2RegionsExecute(r github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiListSrcmV2RegionsRequest) (github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.SrcmV2RegionList, *net_http.Response, error) {
	m.lockListSrcmV2RegionsExecute.Lock()
	defer m.lockListSrcmV2RegionsExecute.Unlock()

	if m.ListSrcmV2RegionsExecuteFunc == nil {
		panic("mocker: RegionsSrcmV2Api.ListSrcmV2RegionsExecuteFunc is nil but RegionsSrcmV2Api.ListSrcmV2RegionsExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiListSrcmV2RegionsRequest
	}{
		R: r,
	}

	m.calls.ListSrcmV2RegionsExecute = append(m.calls.ListSrcmV2RegionsExecute, call)

	return m.ListSrcmV2RegionsExecuteFunc(r)
}

// ListSrcmV2RegionsExecuteCalled returns true if ListSrcmV2RegionsExecute was called at least once.
func (m *RegionsSrcmV2Api) ListSrcmV2RegionsExecuteCalled() bool {
	m.lockListSrcmV2RegionsExecute.Lock()
	defer m.lockListSrcmV2RegionsExecute.Unlock()

	return len(m.calls.ListSrcmV2RegionsExecute) > 0
}

// ListSrcmV2RegionsExecuteCalls returns the calls made to ListSrcmV2RegionsExecute.
func (m *RegionsSrcmV2Api) ListSrcmV2RegionsExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_srcm_v2.ApiListSrcmV2RegionsRequest
} {
	m.lockListSrcmV2RegionsExecute.Lock()
	defer m.lockListSrcmV2RegionsExecute.Unlock()

	return m.calls.ListSrcmV2RegionsExecute
}

// Reset resets the calls made to the mocked methods.
func (m *RegionsSrcmV2Api) Reset() {
	m.lockGetSrcmV2Region.Lock()
	m.calls.GetSrcmV2Region = nil
	m.lockGetSrcmV2Region.Unlock()
	m.lockGetSrcmV2RegionExecute.Lock()
	m.calls.GetSrcmV2RegionExecute = nil
	m.lockGetSrcmV2RegionExecute.Unlock()
	m.lockListSrcmV2Regions.Lock()
	m.calls.ListSrcmV2Regions = nil
	m.lockListSrcmV2Regions.Unlock()
	m.lockListSrcmV2RegionsExecute.Lock()
	m.calls.ListSrcmV2RegionsExecute = nil
	m.lockListSrcmV2RegionsExecute.Unlock()
}