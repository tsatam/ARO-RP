// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/operator/metrics (interfaces: Client)

// Package mock_metrics is a generated GoMock package.
package mock_metrics

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// UpdateDnsConfigurationValid mocks base method.
func (m *MockClient) UpdateDnsConfigurationValid(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDnsConfigurationValid", arg0)
}

// UpdateDnsConfigurationValid indicates an expected call of UpdateDnsConfigurationValid.
func (mr *MockClientMockRecorder) UpdateDnsConfigurationValid(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDnsConfigurationValid", reflect.TypeOf((*MockClient)(nil).UpdateDnsConfigurationValid), arg0)
}

// UpdateIngressCertificateValid mocks base method.
func (m *MockClient) UpdateIngressCertificateValid(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateIngressCertificateValid", arg0)
}

// UpdateIngressCertificateValid indicates an expected call of UpdateIngressCertificateValid.
func (mr *MockClientMockRecorder) UpdateIngressCertificateValid(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIngressCertificateValid", reflect.TypeOf((*MockClient)(nil).UpdateIngressCertificateValid), arg0)
}

// UpdateRequiredEndpointAccessible mocks base method.
func (m *MockClient) UpdateRequiredEndpointAccessible(arg0, arg1 string, arg2 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateRequiredEndpointAccessible", arg0, arg1, arg2)
}

// UpdateRequiredEndpointAccessible indicates an expected call of UpdateRequiredEndpointAccessible.
func (mr *MockClientMockRecorder) UpdateRequiredEndpointAccessible(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequiredEndpointAccessible", reflect.TypeOf((*MockClient)(nil).UpdateRequiredEndpointAccessible), arg0, arg1, arg2)
}

// UpdateServicePrincipalValid mocks base method.
func (m *MockClient) UpdateServicePrincipalValid(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateServicePrincipalValid", arg0)
}

// UpdateServicePrincipalValid indicates an expected call of UpdateServicePrincipalValid.
func (mr *MockClientMockRecorder) UpdateServicePrincipalValid(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServicePrincipalValid", reflect.TypeOf((*MockClient)(nil).UpdateServicePrincipalValid), arg0)
}
