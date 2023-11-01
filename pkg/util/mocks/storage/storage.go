// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/storage (interfaces: Manager)

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	storage "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-06-01/storage"
	storage0 "github.com/Azure/azure-sdk-for-go/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// BlobService mocks base method.
func (m *MockManager) BlobService(arg0 context.Context, arg1, arg2 string, arg3 storage.Permissions, arg4 storage.SignedResourceTypes) (*storage0.BlobStorageClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlobService", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*storage0.BlobStorageClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlobService indicates an expected call of BlobService.
func (mr *MockManagerMockRecorder) BlobService(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlobService", reflect.TypeOf((*MockManager)(nil).BlobService), arg0, arg1, arg2, arg3, arg4)
}

// UpdateAccount mocks base method.
func (m *MockManager) UpdateAccount(arg0 context.Context, arg1, arg2 string, arg3 storage.AccountUpdateParameters) (storage.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(storage.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockManagerMockRecorder) UpdateAccount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockManager)(nil).UpdateAccount), arg0, arg1, arg2, arg3)
}
