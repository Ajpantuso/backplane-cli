// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/backplane-cli/pkg/utils (interfaces: OCMInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

// MockOCMInterface is a mock of OCMInterface interface.
type MockOCMInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOCMInterfaceMockRecorder
}

// MockOCMInterfaceMockRecorder is the mock recorder for MockOCMInterface.
type MockOCMInterfaceMockRecorder struct {
	mock *MockOCMInterface
}

// NewMockOCMInterface creates a new mock instance.
func NewMockOCMInterface(ctrl *gomock.Controller) *MockOCMInterface {
	mock := &MockOCMInterface{ctrl: ctrl}
	mock.recorder = &MockOCMInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOCMInterface) EXPECT() *MockOCMInterfaceMockRecorder {
	return m.recorder
}

// GetBackplaneURL mocks base method.
func (m *MockOCMInterface) GetBackplaneURL() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackplaneURL")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackplaneURL indicates an expected call of GetBackplaneURL.
func (mr *MockOCMInterfaceMockRecorder) GetBackplaneURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackplaneURL", reflect.TypeOf((*MockOCMInterface)(nil).GetBackplaneURL))
}

// GetClusterInfoByID mocks base method.
func (m *MockOCMInterface) GetClusterInfoByID(arg0 string) (*v1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterInfoByID", arg0)
	ret0, _ := ret[0].(*v1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterInfoByID indicates an expected call of GetClusterInfoByID.
func (mr *MockOCMInterfaceMockRecorder) GetClusterInfoByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterInfoByID", reflect.TypeOf((*MockOCMInterface)(nil).GetClusterInfoByID), arg0)
}

// GetManagingCluster mocks base method.
func (m *MockOCMInterface) GetManagingCluster(arg0 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagingCluster", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetManagingCluster indicates an expected call of GetManagingCluster.
func (mr *MockOCMInterfaceMockRecorder) GetManagingCluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagingCluster", reflect.TypeOf((*MockOCMInterface)(nil).GetManagingCluster), arg0)
}

// GetOCMAccessToken mocks base method.
func (m *MockOCMInterface) GetOCMAccessToken() (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOCMAccessToken")
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOCMAccessToken indicates an expected call of GetOCMAccessToken.
func (mr *MockOCMInterfaceMockRecorder) GetOCMAccessToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOCMAccessToken", reflect.TypeOf((*MockOCMInterface)(nil).GetOCMAccessToken))
}

// GetTargetCluster mocks base method.
func (m *MockOCMInterface) GetTargetCluster(arg0 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargetCluster", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTargetCluster indicates an expected call of GetTargetCluster.
func (mr *MockOCMInterfaceMockRecorder) GetTargetCluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargetCluster", reflect.TypeOf((*MockOCMInterface)(nil).GetTargetCluster), arg0)
}

// IsClusterHibernating mocks base method.
func (m *MockOCMInterface) IsClusterHibernating(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClusterHibernating", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsClusterHibernating indicates an expected call of IsClusterHibernating.
func (mr *MockOCMInterfaceMockRecorder) IsClusterHibernating(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClusterHibernating", reflect.TypeOf((*MockOCMInterface)(nil).IsClusterHibernating), arg0)
}

// IsProduction mocks base method.
func (m *MockOCMInterface) IsProduction() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsProduction")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsProduction indicates an expected call of IsProduction.
func (mr *MockOCMInterfaceMockRecorder) IsProduction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsProduction", reflect.TypeOf((*MockOCMInterface)(nil).IsProduction))
}
