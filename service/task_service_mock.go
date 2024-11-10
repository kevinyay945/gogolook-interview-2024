// Code generated by MockGen. DO NOT EDIT.
// Source: gogolook/service (interfaces: TaskServiceInterface)
//
// Generated by this command:
//
//	mockgen -destination=task_service_mock.go -package=service -self_package=gogolook/service . TaskServiceInterface
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	domain "gogolook/domain"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockTaskServiceInterface is a mock of TaskServiceInterface interface.
type MockTaskServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTaskServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockTaskServiceInterfaceMockRecorder is the mock recorder for MockTaskServiceInterface.
type MockTaskServiceInterfaceMockRecorder struct {
	mock *MockTaskServiceInterface
}

// NewMockTaskServiceInterface creates a new mock instance.
func NewMockTaskServiceInterface(ctrl *gomock.Controller) *MockTaskServiceInterface {
	mock := &MockTaskServiceInterface{ctrl: ctrl}
	mock.recorder = &MockTaskServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskServiceInterface) EXPECT() *MockTaskServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockTaskServiceInterface) CreateTask(ctx context.Context, id uuid.UUID, name string, status domain.TaskStatus) (domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", ctx, id, name, status)
	ret0, _ := ret[0].(domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskServiceInterfaceMockRecorder) CreateTask(ctx, id, name, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskServiceInterface)(nil).CreateTask), ctx, id, name, status)
}

// DeleteTaskByID mocks base method.
func (m *MockTaskServiceInterface) DeleteTaskByID(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTaskByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTaskByID indicates an expected call of DeleteTaskByID.
func (mr *MockTaskServiceInterfaceMockRecorder) DeleteTaskByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTaskByID", reflect.TypeOf((*MockTaskServiceInterface)(nil).DeleteTaskByID), ctx, id)
}

// FindAllTasks mocks base method.
func (m *MockTaskServiceInterface) FindAllTasks(ctx context.Context) ([]domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllTasks", ctx)
	ret0, _ := ret[0].([]domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllTasks indicates an expected call of FindAllTasks.
func (mr *MockTaskServiceInterfaceMockRecorder) FindAllTasks(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllTasks", reflect.TypeOf((*MockTaskServiceInterface)(nil).FindAllTasks), ctx)
}

// UpdateTaskByID mocks base method.
func (m *MockTaskServiceInterface) UpdateTaskByID(ctx context.Context, id string, task domain.Task) (domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTaskByID", ctx, id, task)
	ret0, _ := ret[0].(domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTaskByID indicates an expected call of UpdateTaskByID.
func (mr *MockTaskServiceInterfaceMockRecorder) UpdateTaskByID(ctx, id, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTaskByID", reflect.TypeOf((*MockTaskServiceInterface)(nil).UpdateTaskByID), ctx, id, task)
}