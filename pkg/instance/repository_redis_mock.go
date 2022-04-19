// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository_redis.go

// Package instance is a generated GoMock package.
package instance

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v20180412 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"
)

// MockRedisTcInstanceNodeRepository is a mock of RedisTcInstanceNodeRepository interface.
type MockRedisTcInstanceNodeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRedisTcInstanceNodeRepositoryMockRecorder
}

// MockRedisTcInstanceNodeRepositoryMockRecorder is the mock recorder for MockRedisTcInstanceNodeRepository.
type MockRedisTcInstanceNodeRepositoryMockRecorder struct {
	mock *MockRedisTcInstanceNodeRepository
}

// NewMockRedisTcInstanceNodeRepository creates a new mock instance.
func NewMockRedisTcInstanceNodeRepository(ctrl *gomock.Controller) *MockRedisTcInstanceNodeRepository {
	mock := &MockRedisTcInstanceNodeRepository{ctrl: ctrl}
	mock.recorder = &MockRedisTcInstanceNodeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisTcInstanceNodeRepository) EXPECT() *MockRedisTcInstanceNodeRepositoryMockRecorder {
	return m.recorder
}

// GetNodeInfo mocks base method.
func (m *MockRedisTcInstanceNodeRepository) GetNodeInfo(instanceId string) (*v20180412.DescribeInstanceNodeInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeInfo", instanceId)
	ret0, _ := ret[0].(*v20180412.DescribeInstanceNodeInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeInfo indicates an expected call of GetNodeInfo.
func (mr *MockRedisTcInstanceNodeRepositoryMockRecorder) GetNodeInfo(instanceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeInfo", reflect.TypeOf((*MockRedisTcInstanceNodeRepository)(nil).GetNodeInfo), instanceId)
}
