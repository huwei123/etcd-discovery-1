// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/etcd-discovery/service (interfaces: ServiceResponse)

// Package servicemock is a generated GoMock package.
package servicemock

import (
	service "github.com/Scalingo/etcd-discovery/service"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockServiceResponse is a mock of ServiceResponse interface
type MockServiceResponse struct {
	ctrl     *gomock.Controller
	recorder *MockServiceResponseMockRecorder
}

// MockServiceResponseMockRecorder is the mock recorder for MockServiceResponse
type MockServiceResponseMockRecorder struct {
	mock *MockServiceResponse
}

// NewMockServiceResponse creates a new mock instance
func NewMockServiceResponse(ctrl *gomock.Controller) *MockServiceResponse {
	mock := &MockServiceResponse{ctrl: ctrl}
	mock.recorder = &MockServiceResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceResponse) EXPECT() *MockServiceResponseMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockServiceResponse) All() (service.Hosts, error) {
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].(service.Hosts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All
func (mr *MockServiceResponseMockRecorder) All() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockServiceResponse)(nil).All))
}

// Err mocks base method
func (m *MockServiceResponse) Err() error {
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockServiceResponseMockRecorder) Err() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockServiceResponse)(nil).Err))
}

// First mocks base method
func (m *MockServiceResponse) First() service.HostResponse {
	ret := m.ctrl.Call(m, "First")
	ret0, _ := ret[0].(service.HostResponse)
	return ret0
}

// First indicates an expected call of First
func (mr *MockServiceResponseMockRecorder) First() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockServiceResponse)(nil).First))
}

// One mocks base method
func (m *MockServiceResponse) One() service.HostResponse {
	ret := m.ctrl.Call(m, "One")
	ret0, _ := ret[0].(service.HostResponse)
	return ret0
}

// One indicates an expected call of One
func (mr *MockServiceResponseMockRecorder) One() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockServiceResponse)(nil).One))
}

// Service mocks base method
func (m *MockServiceResponse) Service() (*service.Service, error) {
	ret := m.ctrl.Call(m, "Service")
	ret0, _ := ret[0].(*service.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Service indicates an expected call of Service
func (mr *MockServiceResponseMockRecorder) Service() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Service", reflect.TypeOf((*MockServiceResponse)(nil).Service))
}

// URL mocks base method
func (m *MockServiceResponse) URL(arg0, arg1 string) (string, error) {
	ret := m.ctrl.Call(m, "URL", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// URL indicates an expected call of URL
func (mr *MockServiceResponseMockRecorder) URL(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockServiceResponse)(nil).URL), arg0, arg1)
}