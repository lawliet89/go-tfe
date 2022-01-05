// Code generated by MockGen. DO NOT EDIT.
// Source: policy_set_version.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockPolicySetVersions is a mock of PolicySetVersions interface.
type MockPolicySetVersions struct {
	ctrl     *gomock.Controller
	recorder *MockPolicySetVersionsMockRecorder
}

// MockPolicySetVersionsMockRecorder is the mock recorder for MockPolicySetVersions.
type MockPolicySetVersionsMockRecorder struct {
	mock *MockPolicySetVersions
}

// NewMockPolicySetVersions creates a new mock instance.
func NewMockPolicySetVersions(ctrl *gomock.Controller) *MockPolicySetVersions {
	mock := &MockPolicySetVersions{ctrl: ctrl}
	mock.recorder = &MockPolicySetVersionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicySetVersions) EXPECT() *MockPolicySetVersionsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPolicySetVersions) Create(ctx context.Context, policySetID string) (*tfe.PolicySetVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, policySetID)
	ret0, _ := ret[0].(*tfe.PolicySetVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPolicySetVersionsMockRecorder) Create(ctx, policySetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPolicySetVersions)(nil).Create), ctx, policySetID)
}

// Read mocks base method.
func (m *MockPolicySetVersions) Read(ctx context.Context, policySetVersionID string) (*tfe.PolicySetVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, policySetVersionID)
	ret0, _ := ret[0].(*tfe.PolicySetVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockPolicySetVersionsMockRecorder) Read(ctx, policySetVersionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockPolicySetVersions)(nil).Read), ctx, policySetVersionID)
}

// Upload mocks base method.
func (m *MockPolicySetVersions) Upload(ctx context.Context, psv tfe.PolicySetVersion, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", ctx, psv, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockPolicySetVersionsMockRecorder) Upload(ctx, psv, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockPolicySetVersions)(nil).Upload), ctx, psv, path)
}