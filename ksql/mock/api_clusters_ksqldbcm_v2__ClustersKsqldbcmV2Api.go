// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: ../v2/api_clusters_ksqldbcm_v2.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2 "github.com/confluentinc/ccloud-sdk-go-v2-internal/ksql/v2"
)

// ClustersKsqldbcmV2Api is a mock of ClustersKsqldbcmV2Api interface
type ClustersKsqldbcmV2Api struct {
	lockCreateKsqldbcmV2Cluster sync.Mutex
	CreateKsqldbcmV2ClusterFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiCreateKsqldbcmV2ClusterRequest

	lockCreateKsqldbcmV2ClusterExecute sync.Mutex
	CreateKsqldbcmV2ClusterExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiCreateKsqldbcmV2ClusterRequest) (github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.KsqldbcmV2Cluster, *net_http.Response, error)

	lockDeleteKsqldbcmV2Cluster sync.Mutex
	DeleteKsqldbcmV2ClusterFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiDeleteKsqldbcmV2ClusterRequest

	lockDeleteKsqldbcmV2ClusterExecute sync.Mutex
	DeleteKsqldbcmV2ClusterExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiDeleteKsqldbcmV2ClusterRequest) (*net_http.Response, error)

	lockGetKsqldbcmV2Cluster sync.Mutex
	GetKsqldbcmV2ClusterFunc func(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiGetKsqldbcmV2ClusterRequest

	lockGetKsqldbcmV2ClusterExecute sync.Mutex
	GetKsqldbcmV2ClusterExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiGetKsqldbcmV2ClusterRequest) (github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.KsqldbcmV2Cluster, *net_http.Response, error)

	lockListKsqldbcmV2Clusters sync.Mutex
	ListKsqldbcmV2ClustersFunc func(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiListKsqldbcmV2ClustersRequest

	lockListKsqldbcmV2ClustersExecute sync.Mutex
	ListKsqldbcmV2ClustersExecuteFunc func(r github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiListKsqldbcmV2ClustersRequest) (github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.KsqldbcmV2ClusterList, *net_http.Response, error)

	calls struct {
		CreateKsqldbcmV2Cluster []struct {
			Ctx context.Context
		}
		CreateKsqldbcmV2ClusterExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiCreateKsqldbcmV2ClusterRequest
		}
		DeleteKsqldbcmV2Cluster []struct {
			Ctx context.Context
			Id  string
		}
		DeleteKsqldbcmV2ClusterExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiDeleteKsqldbcmV2ClusterRequest
		}
		GetKsqldbcmV2Cluster []struct {
			Ctx context.Context
			Id  string
		}
		GetKsqldbcmV2ClusterExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiGetKsqldbcmV2ClusterRequest
		}
		ListKsqldbcmV2Clusters []struct {
			Ctx context.Context
		}
		ListKsqldbcmV2ClustersExecute []struct {
			R github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiListKsqldbcmV2ClustersRequest
		}
	}
}

// CreateKsqldbcmV2Cluster mocks base method by wrapping the associated func.
func (m *ClustersKsqldbcmV2Api) CreateKsqldbcmV2Cluster(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiCreateKsqldbcmV2ClusterRequest {
	m.lockCreateKsqldbcmV2Cluster.Lock()
	defer m.lockCreateKsqldbcmV2Cluster.Unlock()

	if m.CreateKsqldbcmV2ClusterFunc == nil {
		panic("mocker: ClustersKsqldbcmV2Api.CreateKsqldbcmV2ClusterFunc is nil but ClustersKsqldbcmV2Api.CreateKsqldbcmV2Cluster was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.CreateKsqldbcmV2Cluster = append(m.calls.CreateKsqldbcmV2Cluster, call)

	return m.CreateKsqldbcmV2ClusterFunc(ctx)
}

// CreateKsqldbcmV2ClusterCalled returns true if CreateKsqldbcmV2Cluster was called at least once.
func (m *ClustersKsqldbcmV2Api) CreateKsqldbcmV2ClusterCalled() bool {
	m.lockCreateKsqldbcmV2Cluster.Lock()
	defer m.lockCreateKsqldbcmV2Cluster.Unlock()

	return len(m.calls.CreateKsqldbcmV2Cluster) > 0
}

// CreateKsqldbcmV2ClusterCalls returns the calls made to CreateKsqldbcmV2Cluster.
func (m *ClustersKsqldbcmV2Api) CreateKsqldbcmV2ClusterCalls() []struct {
	Ctx context.Context
} {
	m.lockCreateKsqldbcmV2Cluster.Lock()
	defer m.lockCreateKsqldbcmV2Cluster.Unlock()

	return m.calls.CreateKsqldbcmV2Cluster
}

// CreateKsqldbcmV2ClusterExecute mocks base method by wrapping the associated func.
func (m *ClustersKsqldbcmV2Api) CreateKsqldbcmV2ClusterExecute(r github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiCreateKsqldbcmV2ClusterRequest) (github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.KsqldbcmV2Cluster, *net_http.Response, error) {
	m.lockCreateKsqldbcmV2ClusterExecute.Lock()
	defer m.lockCreateKsqldbcmV2ClusterExecute.Unlock()

	if m.CreateKsqldbcmV2ClusterExecuteFunc == nil {
		panic("mocker: ClustersKsqldbcmV2Api.CreateKsqldbcmV2ClusterExecuteFunc is nil but ClustersKsqldbcmV2Api.CreateKsqldbcmV2ClusterExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiCreateKsqldbcmV2ClusterRequest
	}{
		R: r,
	}

	m.calls.CreateKsqldbcmV2ClusterExecute = append(m.calls.CreateKsqldbcmV2ClusterExecute, call)

	return m.CreateKsqldbcmV2ClusterExecuteFunc(r)
}

// CreateKsqldbcmV2ClusterExecuteCalled returns true if CreateKsqldbcmV2ClusterExecute was called at least once.
func (m *ClustersKsqldbcmV2Api) CreateKsqldbcmV2ClusterExecuteCalled() bool {
	m.lockCreateKsqldbcmV2ClusterExecute.Lock()
	defer m.lockCreateKsqldbcmV2ClusterExecute.Unlock()

	return len(m.calls.CreateKsqldbcmV2ClusterExecute) > 0
}

// CreateKsqldbcmV2ClusterExecuteCalls returns the calls made to CreateKsqldbcmV2ClusterExecute.
func (m *ClustersKsqldbcmV2Api) CreateKsqldbcmV2ClusterExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiCreateKsqldbcmV2ClusterRequest
} {
	m.lockCreateKsqldbcmV2ClusterExecute.Lock()
	defer m.lockCreateKsqldbcmV2ClusterExecute.Unlock()

	return m.calls.CreateKsqldbcmV2ClusterExecute
}

// DeleteKsqldbcmV2Cluster mocks base method by wrapping the associated func.
func (m *ClustersKsqldbcmV2Api) DeleteKsqldbcmV2Cluster(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiDeleteKsqldbcmV2ClusterRequest {
	m.lockDeleteKsqldbcmV2Cluster.Lock()
	defer m.lockDeleteKsqldbcmV2Cluster.Unlock()

	if m.DeleteKsqldbcmV2ClusterFunc == nil {
		panic("mocker: ClustersKsqldbcmV2Api.DeleteKsqldbcmV2ClusterFunc is nil but ClustersKsqldbcmV2Api.DeleteKsqldbcmV2Cluster was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.DeleteKsqldbcmV2Cluster = append(m.calls.DeleteKsqldbcmV2Cluster, call)

	return m.DeleteKsqldbcmV2ClusterFunc(ctx, id)
}

// DeleteKsqldbcmV2ClusterCalled returns true if DeleteKsqldbcmV2Cluster was called at least once.
func (m *ClustersKsqldbcmV2Api) DeleteKsqldbcmV2ClusterCalled() bool {
	m.lockDeleteKsqldbcmV2Cluster.Lock()
	defer m.lockDeleteKsqldbcmV2Cluster.Unlock()

	return len(m.calls.DeleteKsqldbcmV2Cluster) > 0
}

// DeleteKsqldbcmV2ClusterCalls returns the calls made to DeleteKsqldbcmV2Cluster.
func (m *ClustersKsqldbcmV2Api) DeleteKsqldbcmV2ClusterCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockDeleteKsqldbcmV2Cluster.Lock()
	defer m.lockDeleteKsqldbcmV2Cluster.Unlock()

	return m.calls.DeleteKsqldbcmV2Cluster
}

// DeleteKsqldbcmV2ClusterExecute mocks base method by wrapping the associated func.
func (m *ClustersKsqldbcmV2Api) DeleteKsqldbcmV2ClusterExecute(r github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiDeleteKsqldbcmV2ClusterRequest) (*net_http.Response, error) {
	m.lockDeleteKsqldbcmV2ClusterExecute.Lock()
	defer m.lockDeleteKsqldbcmV2ClusterExecute.Unlock()

	if m.DeleteKsqldbcmV2ClusterExecuteFunc == nil {
		panic("mocker: ClustersKsqldbcmV2Api.DeleteKsqldbcmV2ClusterExecuteFunc is nil but ClustersKsqldbcmV2Api.DeleteKsqldbcmV2ClusterExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiDeleteKsqldbcmV2ClusterRequest
	}{
		R: r,
	}

	m.calls.DeleteKsqldbcmV2ClusterExecute = append(m.calls.DeleteKsqldbcmV2ClusterExecute, call)

	return m.DeleteKsqldbcmV2ClusterExecuteFunc(r)
}

// DeleteKsqldbcmV2ClusterExecuteCalled returns true if DeleteKsqldbcmV2ClusterExecute was called at least once.
func (m *ClustersKsqldbcmV2Api) DeleteKsqldbcmV2ClusterExecuteCalled() bool {
	m.lockDeleteKsqldbcmV2ClusterExecute.Lock()
	defer m.lockDeleteKsqldbcmV2ClusterExecute.Unlock()

	return len(m.calls.DeleteKsqldbcmV2ClusterExecute) > 0
}

// DeleteKsqldbcmV2ClusterExecuteCalls returns the calls made to DeleteKsqldbcmV2ClusterExecute.
func (m *ClustersKsqldbcmV2Api) DeleteKsqldbcmV2ClusterExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiDeleteKsqldbcmV2ClusterRequest
} {
	m.lockDeleteKsqldbcmV2ClusterExecute.Lock()
	defer m.lockDeleteKsqldbcmV2ClusterExecute.Unlock()

	return m.calls.DeleteKsqldbcmV2ClusterExecute
}

// GetKsqldbcmV2Cluster mocks base method by wrapping the associated func.
func (m *ClustersKsqldbcmV2Api) GetKsqldbcmV2Cluster(ctx context.Context, id string) github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiGetKsqldbcmV2ClusterRequest {
	m.lockGetKsqldbcmV2Cluster.Lock()
	defer m.lockGetKsqldbcmV2Cluster.Unlock()

	if m.GetKsqldbcmV2ClusterFunc == nil {
		panic("mocker: ClustersKsqldbcmV2Api.GetKsqldbcmV2ClusterFunc is nil but ClustersKsqldbcmV2Api.GetKsqldbcmV2Cluster was called.")
	}

	call := struct {
		Ctx context.Context
		Id  string
	}{
		Ctx: ctx,
		Id:  id,
	}

	m.calls.GetKsqldbcmV2Cluster = append(m.calls.GetKsqldbcmV2Cluster, call)

	return m.GetKsqldbcmV2ClusterFunc(ctx, id)
}

// GetKsqldbcmV2ClusterCalled returns true if GetKsqldbcmV2Cluster was called at least once.
func (m *ClustersKsqldbcmV2Api) GetKsqldbcmV2ClusterCalled() bool {
	m.lockGetKsqldbcmV2Cluster.Lock()
	defer m.lockGetKsqldbcmV2Cluster.Unlock()

	return len(m.calls.GetKsqldbcmV2Cluster) > 0
}

// GetKsqldbcmV2ClusterCalls returns the calls made to GetKsqldbcmV2Cluster.
func (m *ClustersKsqldbcmV2Api) GetKsqldbcmV2ClusterCalls() []struct {
	Ctx context.Context
	Id  string
} {
	m.lockGetKsqldbcmV2Cluster.Lock()
	defer m.lockGetKsqldbcmV2Cluster.Unlock()

	return m.calls.GetKsqldbcmV2Cluster
}

// GetKsqldbcmV2ClusterExecute mocks base method by wrapping the associated func.
func (m *ClustersKsqldbcmV2Api) GetKsqldbcmV2ClusterExecute(r github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiGetKsqldbcmV2ClusterRequest) (github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.KsqldbcmV2Cluster, *net_http.Response, error) {
	m.lockGetKsqldbcmV2ClusterExecute.Lock()
	defer m.lockGetKsqldbcmV2ClusterExecute.Unlock()

	if m.GetKsqldbcmV2ClusterExecuteFunc == nil {
		panic("mocker: ClustersKsqldbcmV2Api.GetKsqldbcmV2ClusterExecuteFunc is nil but ClustersKsqldbcmV2Api.GetKsqldbcmV2ClusterExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiGetKsqldbcmV2ClusterRequest
	}{
		R: r,
	}

	m.calls.GetKsqldbcmV2ClusterExecute = append(m.calls.GetKsqldbcmV2ClusterExecute, call)

	return m.GetKsqldbcmV2ClusterExecuteFunc(r)
}

// GetKsqldbcmV2ClusterExecuteCalled returns true if GetKsqldbcmV2ClusterExecute was called at least once.
func (m *ClustersKsqldbcmV2Api) GetKsqldbcmV2ClusterExecuteCalled() bool {
	m.lockGetKsqldbcmV2ClusterExecute.Lock()
	defer m.lockGetKsqldbcmV2ClusterExecute.Unlock()

	return len(m.calls.GetKsqldbcmV2ClusterExecute) > 0
}

// GetKsqldbcmV2ClusterExecuteCalls returns the calls made to GetKsqldbcmV2ClusterExecute.
func (m *ClustersKsqldbcmV2Api) GetKsqldbcmV2ClusterExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiGetKsqldbcmV2ClusterRequest
} {
	m.lockGetKsqldbcmV2ClusterExecute.Lock()
	defer m.lockGetKsqldbcmV2ClusterExecute.Unlock()

	return m.calls.GetKsqldbcmV2ClusterExecute
}

// ListKsqldbcmV2Clusters mocks base method by wrapping the associated func.
func (m *ClustersKsqldbcmV2Api) ListKsqldbcmV2Clusters(ctx context.Context) github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiListKsqldbcmV2ClustersRequest {
	m.lockListKsqldbcmV2Clusters.Lock()
	defer m.lockListKsqldbcmV2Clusters.Unlock()

	if m.ListKsqldbcmV2ClustersFunc == nil {
		panic("mocker: ClustersKsqldbcmV2Api.ListKsqldbcmV2ClustersFunc is nil but ClustersKsqldbcmV2Api.ListKsqldbcmV2Clusters was called.")
	}

	call := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}

	m.calls.ListKsqldbcmV2Clusters = append(m.calls.ListKsqldbcmV2Clusters, call)

	return m.ListKsqldbcmV2ClustersFunc(ctx)
}

// ListKsqldbcmV2ClustersCalled returns true if ListKsqldbcmV2Clusters was called at least once.
func (m *ClustersKsqldbcmV2Api) ListKsqldbcmV2ClustersCalled() bool {
	m.lockListKsqldbcmV2Clusters.Lock()
	defer m.lockListKsqldbcmV2Clusters.Unlock()

	return len(m.calls.ListKsqldbcmV2Clusters) > 0
}

// ListKsqldbcmV2ClustersCalls returns the calls made to ListKsqldbcmV2Clusters.
func (m *ClustersKsqldbcmV2Api) ListKsqldbcmV2ClustersCalls() []struct {
	Ctx context.Context
} {
	m.lockListKsqldbcmV2Clusters.Lock()
	defer m.lockListKsqldbcmV2Clusters.Unlock()

	return m.calls.ListKsqldbcmV2Clusters
}

// ListKsqldbcmV2ClustersExecute mocks base method by wrapping the associated func.
func (m *ClustersKsqldbcmV2Api) ListKsqldbcmV2ClustersExecute(r github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiListKsqldbcmV2ClustersRequest) (github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.KsqldbcmV2ClusterList, *net_http.Response, error) {
	m.lockListKsqldbcmV2ClustersExecute.Lock()
	defer m.lockListKsqldbcmV2ClustersExecute.Unlock()

	if m.ListKsqldbcmV2ClustersExecuteFunc == nil {
		panic("mocker: ClustersKsqldbcmV2Api.ListKsqldbcmV2ClustersExecuteFunc is nil but ClustersKsqldbcmV2Api.ListKsqldbcmV2ClustersExecute was called.")
	}

	call := struct {
		R github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiListKsqldbcmV2ClustersRequest
	}{
		R: r,
	}

	m.calls.ListKsqldbcmV2ClustersExecute = append(m.calls.ListKsqldbcmV2ClustersExecute, call)

	return m.ListKsqldbcmV2ClustersExecuteFunc(r)
}

// ListKsqldbcmV2ClustersExecuteCalled returns true if ListKsqldbcmV2ClustersExecute was called at least once.
func (m *ClustersKsqldbcmV2Api) ListKsqldbcmV2ClustersExecuteCalled() bool {
	m.lockListKsqldbcmV2ClustersExecute.Lock()
	defer m.lockListKsqldbcmV2ClustersExecute.Unlock()

	return len(m.calls.ListKsqldbcmV2ClustersExecute) > 0
}

// ListKsqldbcmV2ClustersExecuteCalls returns the calls made to ListKsqldbcmV2ClustersExecute.
func (m *ClustersKsqldbcmV2Api) ListKsqldbcmV2ClustersExecuteCalls() []struct {
	R github_com_confluentinc_ccloud_sdk_go_v2_internal_ksql_v2.ApiListKsqldbcmV2ClustersRequest
} {
	m.lockListKsqldbcmV2ClustersExecute.Lock()
	defer m.lockListKsqldbcmV2ClustersExecute.Unlock()

	return m.calls.ListKsqldbcmV2ClustersExecute
}

// Reset resets the calls made to the mocked methods.
func (m *ClustersKsqldbcmV2Api) Reset() {
	m.lockCreateKsqldbcmV2Cluster.Lock()
	m.calls.CreateKsqldbcmV2Cluster = nil
	m.lockCreateKsqldbcmV2Cluster.Unlock()
	m.lockCreateKsqldbcmV2ClusterExecute.Lock()
	m.calls.CreateKsqldbcmV2ClusterExecute = nil
	m.lockCreateKsqldbcmV2ClusterExecute.Unlock()
	m.lockDeleteKsqldbcmV2Cluster.Lock()
	m.calls.DeleteKsqldbcmV2Cluster = nil
	m.lockDeleteKsqldbcmV2Cluster.Unlock()
	m.lockDeleteKsqldbcmV2ClusterExecute.Lock()
	m.calls.DeleteKsqldbcmV2ClusterExecute = nil
	m.lockDeleteKsqldbcmV2ClusterExecute.Unlock()
	m.lockGetKsqldbcmV2Cluster.Lock()
	m.calls.GetKsqldbcmV2Cluster = nil
	m.lockGetKsqldbcmV2Cluster.Unlock()
	m.lockGetKsqldbcmV2ClusterExecute.Lock()
	m.calls.GetKsqldbcmV2ClusterExecute = nil
	m.lockGetKsqldbcmV2ClusterExecute.Unlock()
	m.lockListKsqldbcmV2Clusters.Lock()
	m.calls.ListKsqldbcmV2Clusters = nil
	m.lockListKsqldbcmV2Clusters.Unlock()
	m.lockListKsqldbcmV2ClustersExecute.Lock()
	m.calls.ListKsqldbcmV2ClustersExecute = nil
	m.lockListKsqldbcmV2ClustersExecute.Unlock()
}
