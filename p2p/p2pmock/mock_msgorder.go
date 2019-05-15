// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aergoio/aergo/p2p/p2pcommon (interfaces: MsgOrder)

// Package p2pmock is a generated GoMock package.
package p2pmock

import (
	p2pcommon "github.com/aergoio/aergo/p2p/p2pcommon"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMsgOrder is a mock of MsgOrder interface
type MockMsgOrder struct {
	ctrl     *gomock.Controller
	recorder *MockMsgOrderMockRecorder
}

// MockMsgOrderMockRecorder is the mock recorder for MockMsgOrder
type MockMsgOrderMockRecorder struct {
	mock *MockMsgOrder
}

// NewMockMsgOrder creates a new mock instance
func NewMockMsgOrder(ctrl *gomock.Controller) *MockMsgOrder {
	mock := &MockMsgOrder{ctrl: ctrl}
	mock.recorder = &MockMsgOrderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMsgOrder) EXPECT() *MockMsgOrderMockRecorder {
	return m.recorder
}

// GetMsgID mocks base method
func (m *MockMsgOrder) GetMsgID() p2pcommon.MsgID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMsgID")
	ret0, _ := ret[0].(p2pcommon.MsgID)
	return ret0
}

// GetMsgID indicates an expected call of GetMsgID
func (mr *MockMsgOrderMockRecorder) GetMsgID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMsgID", reflect.TypeOf((*MockMsgOrder)(nil).GetMsgID))
}

// GetProtocolID mocks base method
func (m *MockMsgOrder) GetProtocolID() p2pcommon.SubProtocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProtocolID")
	ret0, _ := ret[0].(p2pcommon.SubProtocol)
	return ret0
}

// GetProtocolID indicates an expected call of GetProtocolID
func (mr *MockMsgOrderMockRecorder) GetProtocolID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProtocolID", reflect.TypeOf((*MockMsgOrder)(nil).GetProtocolID))
}

// IsNeedSign mocks base method
func (m *MockMsgOrder) IsNeedSign() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNeedSign")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNeedSign indicates an expected call of IsNeedSign
func (mr *MockMsgOrderMockRecorder) IsNeedSign() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNeedSign", reflect.TypeOf((*MockMsgOrder)(nil).IsNeedSign))
}

// IsRequest mocks base method
func (m *MockMsgOrder) IsRequest() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRequest")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRequest indicates an expected call of IsRequest
func (mr *MockMsgOrderMockRecorder) IsRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRequest", reflect.TypeOf((*MockMsgOrder)(nil).IsRequest))
}

// SendTo mocks base method
func (m *MockMsgOrder) SendTo(arg0 p2pcommon.RemotePeer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTo indicates an expected call of SendTo
func (mr *MockMsgOrderMockRecorder) SendTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTo", reflect.TypeOf((*MockMsgOrder)(nil).SendTo), arg0)
}

// Timestamp mocks base method
func (m *MockMsgOrder) Timestamp() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timestamp")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Timestamp indicates an expected call of Timestamp
func (mr *MockMsgOrderMockRecorder) Timestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timestamp", reflect.TypeOf((*MockMsgOrder)(nil).Timestamp))
}
