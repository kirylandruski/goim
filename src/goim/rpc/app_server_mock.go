// Code generated by MockGen. DO NOT EDIT.
// Source: app_server.go

// Package rpc is a generated GoMock package.
package rpc

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAppServer is a mock of AppServer interface
type MockAppServer struct {
	ctrl     *gomock.Controller
	recorder *MockAppServerMockRecorder
}

// MockAppServerMockRecorder is the mock recorder for MockAppServer
type MockAppServerMockRecorder struct {
	mock *MockAppServer
}

// NewMockAppServer creates a new mock instance
func NewMockAppServer(ctrl *gomock.Controller) *MockAppServer {
	mock := &MockAppServer{ctrl: ctrl}
	mock.recorder = &MockAppServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppServer) EXPECT() *MockAppServerMockRecorder {
	return m.recorder
}

// Authenticate mocks base method
func (m *MockAppServer) Authenticate(username, password *string) int32 {
	ret := m.ctrl.Call(m, "Authenticate", username, password)
	ret0, _ := ret[0].(int32)
	return ret0
}

// Authenticate indicates an expected call of Authenticate
func (mr *MockAppServerMockRecorder) Authenticate(username, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAppServer)(nil).Authenticate), username, password)
}

// GetKeys mocks base method
func (m *MockAppServer) GetKeys() ([]*string, int32) {
	ret := m.ctrl.Call(m, "GetKeys")
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(int32)
	return ret0, ret1
}

// GetKeys indicates an expected call of GetKeys
func (mr *MockAppServerMockRecorder) GetKeys() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeys", reflect.TypeOf((*MockAppServer)(nil).GetKeys))
}

// GetStr mocks base method
func (m *MockAppServer) GetStr(key *string) (*string, int64, int32) {
	ret := m.ctrl.Call(m, "GetStr", key)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(int32)
	return ret0, ret1, ret2
}

// GetStr indicates an expected call of GetStr
func (mr *MockAppServerMockRecorder) GetStr(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStr", reflect.TypeOf((*MockAppServer)(nil).GetStr), key)
}

// SetStr mocks base method
func (m *MockAppServer) SetStr(key, str *string, ttl int64) (int64, int32) {
	ret := m.ctrl.Call(m, "SetStr", key, str, ttl)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int32)
	return ret0, ret1
}

// SetStr indicates an expected call of SetStr
func (mr *MockAppServerMockRecorder) SetStr(key, str, ttl interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStr", reflect.TypeOf((*MockAppServer)(nil).SetStr), key, str, ttl)
}

// GetArr mocks base method
func (m *MockAppServer) GetArr(key *string) ([]*string, int64, int32) {
	ret := m.ctrl.Call(m, "GetArr", key)
	ret0, _ := ret[0].([]*string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(int32)
	return ret0, ret1, ret2
}

// GetArr indicates an expected call of GetArr
func (mr *MockAppServerMockRecorder) GetArr(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArr", reflect.TypeOf((*MockAppServer)(nil).GetArr), key)
}

// SetArr mocks base method
func (m *MockAppServer) SetArr(key *string, arr []*string, ttl int64) (int64, int32) {
	ret := m.ctrl.Call(m, "SetArr", key, arr, ttl)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int32)
	return ret0, ret1
}

// SetArr indicates an expected call of SetArr
func (mr *MockAppServerMockRecorder) SetArr(key, arr, ttl interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetArr", reflect.TypeOf((*MockAppServer)(nil).SetArr), key, arr, ttl)
}

// GetArrItem mocks base method
func (m *MockAppServer) GetArrItem(key *string, index int32) (*string, int64, int32) {
	ret := m.ctrl.Call(m, "GetArrItem", key, index)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(int32)
	return ret0, ret1, ret2
}

// GetArrItem indicates an expected call of GetArrItem
func (mr *MockAppServerMockRecorder) GetArrItem(key, index interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArrItem", reflect.TypeOf((*MockAppServer)(nil).GetArrItem), key, index)
}

// SetArrItem mocks base method
func (m *MockAppServer) SetArrItem(key *string, index int32, str *string, ttl int64) (int64, int32) {
	ret := m.ctrl.Call(m, "SetArrItem", key, index, str, ttl)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int32)
	return ret0, ret1
}

// SetArrItem indicates an expected call of SetArrItem
func (mr *MockAppServerMockRecorder) SetArrItem(key, index, str, ttl interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetArrItem", reflect.TypeOf((*MockAppServer)(nil).SetArrItem), key, index, str, ttl)
}

// GetDict mocks base method
func (m *MockAppServer) GetDict(key *string) (map[string]*string, int64, int32) {
	ret := m.ctrl.Call(m, "GetDict", key)
	ret0, _ := ret[0].(map[string]*string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(int32)
	return ret0, ret1, ret2
}

// GetDict indicates an expected call of GetDict
func (mr *MockAppServerMockRecorder) GetDict(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDict", reflect.TypeOf((*MockAppServer)(nil).GetDict), key)
}

// SetDict mocks base method
func (m *MockAppServer) SetDict(key *string, dict map[string]*string, ttl int64) (int64, int32) {
	ret := m.ctrl.Call(m, "SetDict", key, dict, ttl)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int32)
	return ret0, ret1
}

// SetDict indicates an expected call of SetDict
func (mr *MockAppServerMockRecorder) SetDict(key, dict, ttl interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDict", reflect.TypeOf((*MockAppServer)(nil).SetDict), key, dict, ttl)
}

// GetDictItem mocks base method
func (m *MockAppServer) GetDictItem(key, subkey *string) (*string, int64, int32) {
	ret := m.ctrl.Call(m, "GetDictItem", key, subkey)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(int32)
	return ret0, ret1, ret2
}

// GetDictItem indicates an expected call of GetDictItem
func (mr *MockAppServerMockRecorder) GetDictItem(key, subkey interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDictItem", reflect.TypeOf((*MockAppServer)(nil).GetDictItem), key, subkey)
}

// SetDictItem mocks base method
func (m *MockAppServer) SetDictItem(key, subkey, str *string, ttl int64) (int64, int32) {
	ret := m.ctrl.Call(m, "SetDictItem", key, subkey, str, ttl)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int32)
	return ret0, ret1
}

// SetDictItem indicates an expected call of SetDictItem
func (mr *MockAppServerMockRecorder) SetDictItem(key, subkey, str, ttl interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDictItem", reflect.TypeOf((*MockAppServer)(nil).SetDictItem), key, subkey, str, ttl)
}

// RemoveEntry mocks base method
func (m *MockAppServer) RemoveEntry(key *string) int32 {
	ret := m.ctrl.Call(m, "RemoveEntry", key)
	ret0, _ := ret[0].(int32)
	return ret0
}

// RemoveEntry indicates an expected call of RemoveEntry
func (mr *MockAppServerMockRecorder) RemoveEntry(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEntry", reflect.TypeOf((*MockAppServer)(nil).RemoveEntry), key)
}
