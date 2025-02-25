// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-ebpf-sdk-go/pkg/maps (interfaces: BpfMapAPIs)

// Package mock_maps is a generated GoMock package.
package mock_maps

import (
	reflect "reflect"

	maps "github.com/aws/aws-ebpf-sdk-go/pkg/maps"
	gomock "github.com/golang/mock/gomock"
)

// MockBpfMapAPIs is a mock of BpfMapAPIs interface.
type MockBpfMapAPIs struct {
	ctrl     *gomock.Controller
	recorder *MockBpfMapAPIsMockRecorder
}

// MockBpfMapAPIsMockRecorder is the mock recorder for MockBpfMapAPIs.
type MockBpfMapAPIsMockRecorder struct {
	mock *MockBpfMapAPIs
}

// NewMockBpfMapAPIs creates a new mock instance.
func NewMockBpfMapAPIs(ctrl *gomock.Controller) *MockBpfMapAPIs {
	mock := &MockBpfMapAPIs{ctrl: ctrl}
	mock.recorder = &MockBpfMapAPIsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBpfMapAPIs) EXPECT() *MockBpfMapAPIsMockRecorder {
	return m.recorder
}

// BulkDeleteMapEntry mocks base method.
func (m *MockBpfMapAPIs) BulkDeleteMapEntry(arg0 map[uintptr]uintptr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteMapEntry", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteMapEntry indicates an expected call of BulkDeleteMapEntry.
func (mr *MockBpfMapAPIsMockRecorder) BulkDeleteMapEntry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteMapEntry", reflect.TypeOf((*MockBpfMapAPIs)(nil).BulkDeleteMapEntry), arg0)
}

// BulkRefreshMapEntries mocks base method.
func (m *MockBpfMapAPIs) BulkRefreshMapEntries(arg0 map[string][]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkRefreshMapEntries", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkRefreshMapEntries indicates an expected call of BulkRefreshMapEntries.
func (mr *MockBpfMapAPIsMockRecorder) BulkRefreshMapEntries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkRefreshMapEntries", reflect.TypeOf((*MockBpfMapAPIs)(nil).BulkRefreshMapEntries), arg0)
}

// BulkUpdateMapEntry mocks base method.
func (m *MockBpfMapAPIs) BulkUpdateMapEntry(arg0 map[string][]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkUpdateMapEntry", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkUpdateMapEntry indicates an expected call of BulkUpdateMapEntry.
func (mr *MockBpfMapAPIsMockRecorder) BulkUpdateMapEntry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUpdateMapEntry", reflect.TypeOf((*MockBpfMapAPIs)(nil).BulkUpdateMapEntry), arg0)
}

// CreateBPFMap mocks base method.
func (m *MockBpfMapAPIs) CreateBPFMap(arg0 maps.CreateEBPFMapInput) (maps.BpfMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBPFMap", arg0)
	ret0, _ := ret[0].(maps.BpfMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBPFMap indicates an expected call of CreateBPFMap.
func (mr *MockBpfMapAPIsMockRecorder) CreateBPFMap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBPFMap", reflect.TypeOf((*MockBpfMapAPIs)(nil).CreateBPFMap), arg0)
}

// CreateMapEntry mocks base method.
func (m *MockBpfMapAPIs) CreateMapEntry(arg0, arg1 uintptr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMapEntry", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMapEntry indicates an expected call of CreateMapEntry.
func (mr *MockBpfMapAPIsMockRecorder) CreateMapEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMapEntry", reflect.TypeOf((*MockBpfMapAPIs)(nil).CreateMapEntry), arg0, arg1)
}

// CreateUpdateMapEntry mocks base method.
func (m *MockBpfMapAPIs) CreateUpdateMapEntry(arg0, arg1 uintptr, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUpdateMapEntry", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUpdateMapEntry indicates an expected call of CreateUpdateMapEntry.
func (mr *MockBpfMapAPIsMockRecorder) CreateUpdateMapEntry(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUpdateMapEntry", reflect.TypeOf((*MockBpfMapAPIs)(nil).CreateUpdateMapEntry), arg0, arg1, arg2)
}

// DeleteMapEntry mocks base method.
func (m *MockBpfMapAPIs) DeleteMapEntry(arg0 uintptr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMapEntry", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMapEntry indicates an expected call of DeleteMapEntry.
func (mr *MockBpfMapAPIsMockRecorder) DeleteMapEntry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMapEntry", reflect.TypeOf((*MockBpfMapAPIs)(nil).DeleteMapEntry), arg0)
}

// GetBPFmapInfo mocks base method.
func (m *MockBpfMapAPIs) GetBPFmapInfo(arg0 uint32) (maps.BpfMapInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBPFmapInfo", arg0)
	ret0, _ := ret[0].(maps.BpfMapInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBPFmapInfo indicates an expected call of GetBPFmapInfo.
func (mr *MockBpfMapAPIsMockRecorder) GetBPFmapInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBPFmapInfo", reflect.TypeOf((*MockBpfMapAPIs)(nil).GetBPFmapInfo), arg0)
}

// GetFirstMapEntry mocks base method.
func (m *MockBpfMapAPIs) GetFirstMapEntry(arg0 uintptr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirstMapEntry", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetFirstMapEntry indicates an expected call of GetFirstMapEntry.
func (mr *MockBpfMapAPIsMockRecorder) GetFirstMapEntry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstMapEntry", reflect.TypeOf((*MockBpfMapAPIs)(nil).GetFirstMapEntry), arg0)
}

// GetMapEntry mocks base method.
func (m *MockBpfMapAPIs) GetMapEntry(arg0, arg1 uintptr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMapEntry", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetMapEntry indicates an expected call of GetMapEntry.
func (mr *MockBpfMapAPIsMockRecorder) GetMapEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMapEntry", reflect.TypeOf((*MockBpfMapAPIs)(nil).GetMapEntry), arg0, arg1)
}

// GetMapFromPinPath mocks base method.
func (m *MockBpfMapAPIs) GetMapFromPinPath(arg0 string) (maps.BpfMapInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMapFromPinPath", arg0)
	ret0, _ := ret[0].(maps.BpfMapInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMapFromPinPath indicates an expected call of GetMapFromPinPath.
func (mr *MockBpfMapAPIsMockRecorder) GetMapFromPinPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMapFromPinPath", reflect.TypeOf((*MockBpfMapAPIs)(nil).GetMapFromPinPath), arg0)
}

// GetNextMapEntry mocks base method.
func (m *MockBpfMapAPIs) GetNextMapEntry(arg0, arg1 uintptr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextMapEntry", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetNextMapEntry indicates an expected call of GetNextMapEntry.
func (mr *MockBpfMapAPIsMockRecorder) GetNextMapEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextMapEntry", reflect.TypeOf((*MockBpfMapAPIs)(nil).GetNextMapEntry), arg0, arg1)
}

// PinMap mocks base method.
func (m *MockBpfMapAPIs) PinMap(arg0 string, arg1 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PinMap", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PinMap indicates an expected call of PinMap.
func (mr *MockBpfMapAPIsMockRecorder) PinMap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PinMap", reflect.TypeOf((*MockBpfMapAPIs)(nil).PinMap), arg0, arg1)
}

// UnPinMap mocks base method.
func (m *MockBpfMapAPIs) UnPinMap(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnPinMap", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnPinMap indicates an expected call of UnPinMap.
func (mr *MockBpfMapAPIsMockRecorder) UnPinMap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnPinMap", reflect.TypeOf((*MockBpfMapAPIs)(nil).UnPinMap), arg0)
}

// UpdateMapEntry mocks base method.
func (m *MockBpfMapAPIs) UpdateMapEntry(arg0, arg1 uintptr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMapEntry", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMapEntry indicates an expected call of UpdateMapEntry.
func (mr *MockBpfMapAPIsMockRecorder) UpdateMapEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMapEntry", reflect.TypeOf((*MockBpfMapAPIs)(nil).UpdateMapEntry), arg0, arg1)
}
