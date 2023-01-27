// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/shared/aws/interfaces (interfaces: SSMAPI)

// Package interfaces is a generated GoMock package.
package interfaces

import (
	context "context"
	reflect "reflect"

	ssm "github.com/aws/aws-sdk-go-v2/service/ssm"
	gomock "github.com/golang/mock/gomock"
)

// MockSSMAPI is a mock of SSMAPI interface.
type MockSSMAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSSMAPIMockRecorder
}

// MockSSMAPIMockRecorder is the mock recorder for MockSSMAPI.
type MockSSMAPIMockRecorder struct {
	mock *MockSSMAPI
}

// NewMockSSMAPI creates a new mock instance.
func NewMockSSMAPI(ctrl *gomock.Controller) *MockSSMAPI {
	mock := &MockSSMAPI{ctrl: ctrl}
	mock.recorder = &MockSSMAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSSMAPI) EXPECT() *MockSSMAPIMockRecorder {
	return m.recorder
}

// GetParameter mocks base method.
func (m *MockSSMAPI) GetParameter(arg0 context.Context, arg1 *ssm.GetParameterInput, arg2 ...func(*ssm.Options)) (*ssm.GetParameterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetParameter", varargs...)
	ret0, _ := ret[0].(*ssm.GetParameterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParameter indicates an expected call of GetParameter.
func (mr *MockSSMAPIMockRecorder) GetParameter(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParameter", reflect.TypeOf((*MockSSMAPI)(nil).GetParameter), varargs...)
}

// PutParameter mocks base method.
func (m *MockSSMAPI) PutParameter(arg0 context.Context, arg1 *ssm.PutParameterInput, arg2 ...func(*ssm.Options)) (*ssm.PutParameterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutParameter", varargs...)
	ret0, _ := ret[0].(*ssm.PutParameterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutParameter indicates an expected call of PutParameter.
func (mr *MockSSMAPIMockRecorder) PutParameter(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutParameter", reflect.TypeOf((*MockSSMAPI)(nil).PutParameter), varargs...)
}
