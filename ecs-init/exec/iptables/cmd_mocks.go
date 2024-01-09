// Copyright 2015-2024 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Source: cmd.go in package iptables
// Code generated by MockGen. DO NOT EDIT.

// Package iptables is a generated GoMock package.
package iptables

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCmd is a mock of Cmd interface.
type MockCmd struct {
	ctrl     *gomock.Controller
	recorder *MockCmdMockRecorder
}

// MockCmdMockRecorder is the mock recorder for MockCmd.
type MockCmdMockRecorder struct {
	mock *MockCmd
}

// NewMockCmd creates a new mock instance.
func NewMockCmd(ctrl *gomock.Controller) *MockCmd {
	mock := &MockCmd{ctrl: ctrl}
	mock.recorder = &MockCmdMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCmd) EXPECT() *MockCmdMockRecorder {
	return m.recorder
}

// CombinedOutput mocks base method.
func (m *MockCmd) CombinedOutput() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CombinedOutput")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CombinedOutput indicates an expected call of CombinedOutput.
func (mr *MockCmdMockRecorder) CombinedOutput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CombinedOutput", reflect.TypeOf((*MockCmd)(nil).CombinedOutput))
}

// Output mocks base method.
func (m *MockCmd) Output() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Output indicates an expected call of Output.
func (mr *MockCmdMockRecorder) Output() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockCmd)(nil).Output))
}
