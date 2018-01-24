// Code generated by MockGen. DO NOT EDIT.
// Source: vendor/k8s.io/helm/pkg/helm/interface.go

package helmctlr

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	helm "k8s.io/helm/pkg/helm"
	chart "k8s.io/helm/pkg/proto/hapi/chart"
	services "k8s.io/helm/pkg/proto/hapi/services"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return _m.recorder
}

// ListReleases mocks base method
func (_m *MockInterface) ListReleases(opts ...helm.ReleaseListOption) (*services.ListReleasesResponse, error) {
	_s := []interface{}{}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ListReleases", _s...)
	ret0, _ := ret[0].(*services.ListReleasesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReleases indicates an expected call of ListReleases
func (_mr *MockInterfaceMockRecorder) ListReleases(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ListReleases", reflect.TypeOf((*MockInterface)(nil).ListReleases), arg0...)
}

// InstallRelease mocks base method
func (_m *MockInterface) InstallRelease(chStr string, namespace string, opts ...helm.InstallOption) (*services.InstallReleaseResponse, error) {
	_s := []interface{}{chStr, namespace}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "InstallRelease", _s...)
	ret0, _ := ret[0].(*services.InstallReleaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallRelease indicates an expected call of InstallRelease
func (_mr *MockInterfaceMockRecorder) InstallRelease(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InstallRelease", reflect.TypeOf((*MockInterface)(nil).InstallRelease), _s...)
}

// InstallReleaseFromChart mocks base method
func (_m *MockInterface) InstallReleaseFromChart(chart *chart.Chart, namespace string, opts ...helm.InstallOption) (*services.InstallReleaseResponse, error) {
	_s := []interface{}{chart, namespace}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "InstallReleaseFromChart", _s...)
	ret0, _ := ret[0].(*services.InstallReleaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallReleaseFromChart indicates an expected call of InstallReleaseFromChart
func (_mr *MockInterfaceMockRecorder) InstallReleaseFromChart(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InstallReleaseFromChart", reflect.TypeOf((*MockInterface)(nil).InstallReleaseFromChart), _s...)
}

// DeleteRelease mocks base method
func (_m *MockInterface) DeleteRelease(rlsName string, opts ...helm.DeleteOption) (*services.UninstallReleaseResponse, error) {
	_s := []interface{}{rlsName}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DeleteRelease", _s...)
	ret0, _ := ret[0].(*services.UninstallReleaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRelease indicates an expected call of DeleteRelease
func (_mr *MockInterfaceMockRecorder) DeleteRelease(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteRelease", reflect.TypeOf((*MockInterface)(nil).DeleteRelease), _s...)
}

// ReleaseStatus mocks base method
func (_m *MockInterface) ReleaseStatus(rlsName string, opts ...helm.StatusOption) (*services.GetReleaseStatusResponse, error) {
	_s := []interface{}{rlsName}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ReleaseStatus", _s...)
	ret0, _ := ret[0].(*services.GetReleaseStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseStatus indicates an expected call of ReleaseStatus
func (_mr *MockInterfaceMockRecorder) ReleaseStatus(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReleaseStatus", reflect.TypeOf((*MockInterface)(nil).ReleaseStatus), _s...)
}

// UpdateRelease mocks base method
func (_m *MockInterface) UpdateRelease(rlsName string, chStr string, opts ...helm.UpdateOption) (*services.UpdateReleaseResponse, error) {
	_s := []interface{}{rlsName, chStr}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "UpdateRelease", _s...)
	ret0, _ := ret[0].(*services.UpdateReleaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRelease indicates an expected call of UpdateRelease
func (_mr *MockInterfaceMockRecorder) UpdateRelease(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateRelease", reflect.TypeOf((*MockInterface)(nil).UpdateRelease), _s...)
}

// UpdateReleaseFromChart mocks base method
func (_m *MockInterface) UpdateReleaseFromChart(rlsName string, chart *chart.Chart, opts ...helm.UpdateOption) (*services.UpdateReleaseResponse, error) {
	_s := []interface{}{rlsName, chart}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "UpdateReleaseFromChart", _s...)
	ret0, _ := ret[0].(*services.UpdateReleaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReleaseFromChart indicates an expected call of UpdateReleaseFromChart
func (_mr *MockInterfaceMockRecorder) UpdateReleaseFromChart(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateReleaseFromChart", reflect.TypeOf((*MockInterface)(nil).UpdateReleaseFromChart), _s...)
}

// RollbackRelease mocks base method
func (_m *MockInterface) RollbackRelease(rlsName string, opts ...helm.RollbackOption) (*services.RollbackReleaseResponse, error) {
	_s := []interface{}{rlsName}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "RollbackRelease", _s...)
	ret0, _ := ret[0].(*services.RollbackReleaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RollbackRelease indicates an expected call of RollbackRelease
func (_mr *MockInterfaceMockRecorder) RollbackRelease(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RollbackRelease", reflect.TypeOf((*MockInterface)(nil).RollbackRelease), _s...)
}

// ReleaseContent mocks base method
func (_m *MockInterface) ReleaseContent(rlsName string, opts ...helm.ContentOption) (*services.GetReleaseContentResponse, error) {
	_s := []interface{}{rlsName}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ReleaseContent", _s...)
	ret0, _ := ret[0].(*services.GetReleaseContentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseContent indicates an expected call of ReleaseContent
func (_mr *MockInterfaceMockRecorder) ReleaseContent(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReleaseContent", reflect.TypeOf((*MockInterface)(nil).ReleaseContent), _s...)
}

// ReleaseHistory mocks base method
func (_m *MockInterface) ReleaseHistory(rlsName string, opts ...helm.HistoryOption) (*services.GetHistoryResponse, error) {
	_s := []interface{}{rlsName}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ReleaseHistory", _s...)
	ret0, _ := ret[0].(*services.GetHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseHistory indicates an expected call of ReleaseHistory
func (_mr *MockInterfaceMockRecorder) ReleaseHistory(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReleaseHistory", reflect.TypeOf((*MockInterface)(nil).ReleaseHistory), _s...)
}

// GetVersion mocks base method
func (_m *MockInterface) GetVersion(opts ...helm.VersionOption) (*services.GetVersionResponse, error) {
	_s := []interface{}{}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetVersion", _s...)
	ret0, _ := ret[0].(*services.GetVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion
func (_mr *MockInterfaceMockRecorder) GetVersion(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetVersion", reflect.TypeOf((*MockInterface)(nil).GetVersion), arg0...)
}

// RunReleaseTest mocks base method
func (_m *MockInterface) RunReleaseTest(rlsName string, opts ...helm.ReleaseTestOption) (<-chan *services.TestReleaseResponse, <-chan error) {
	_s := []interface{}{rlsName}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "RunReleaseTest", _s...)
	ret0, _ := ret[0].(<-chan *services.TestReleaseResponse)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// RunReleaseTest indicates an expected call of RunReleaseTest
func (_mr *MockInterfaceMockRecorder) RunReleaseTest(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RunReleaseTest", reflect.TypeOf((*MockInterface)(nil).RunReleaseTest), _s...)
}
