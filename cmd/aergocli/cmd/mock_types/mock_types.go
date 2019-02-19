// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aergoio/aergo/types (interfaces: AergoRPCServiceClient)

// Package mock_types is a generated GoMock package.
package mock_types

import (
	context "context"
	types "github.com/aergoio/aergo/types"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAergoRPCServiceClient is a mock of AergoRPCServiceClient interface
type MockAergoRPCServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAergoRPCServiceClientMockRecorder
}

// MockAergoRPCServiceClientMockRecorder is the mock recorder for MockAergoRPCServiceClient
type MockAergoRPCServiceClientMockRecorder struct {
	mock *MockAergoRPCServiceClient
}

// NewMockAergoRPCServiceClient creates a new mock instance
func NewMockAergoRPCServiceClient(ctrl *gomock.Controller) *MockAergoRPCServiceClient {
	mock := &MockAergoRPCServiceClient{ctrl: ctrl}
	mock.recorder = &MockAergoRPCServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAergoRPCServiceClient) EXPECT() *MockAergoRPCServiceClientMockRecorder {
	return m.recorder
}

// Blockchain mocks base method
func (m *MockAergoRPCServiceClient) Blockchain(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (*types.BlockchainStatus, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Blockchain", varargs...)
	ret0, _ := ret[0].(*types.BlockchainStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Blockchain indicates an expected call of Blockchain
func (mr *MockAergoRPCServiceClientMockRecorder) Blockchain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Blockchain", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).Blockchain), varargs...)
}

// CommitTX mocks base method
func (m *MockAergoRPCServiceClient) CommitTX(arg0 context.Context, arg1 *types.TxList, arg2 ...grpc.CallOption) (*types.CommitResultList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CommitTX", varargs...)
	ret0, _ := ret[0].(*types.CommitResultList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitTX indicates an expected call of CommitTX
func (mr *MockAergoRPCServiceClientMockRecorder) CommitTX(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitTX", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).CommitTX), varargs...)
}

// CreateAccount mocks base method
func (m *MockAergoRPCServiceClient) CreateAccount(arg0 context.Context, arg1 *types.Personal, arg2 ...grpc.CallOption) (*types.Account, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAccount", varargs...)
	ret0, _ := ret[0].(*types.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockAergoRPCServiceClientMockRecorder) CreateAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).CreateAccount), varargs...)
}

// ExportAccount mocks base method
func (m *MockAergoRPCServiceClient) ExportAccount(arg0 context.Context, arg1 *types.Personal, arg2 ...grpc.CallOption) (*types.SingleBytes, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExportAccount", varargs...)
	ret0, _ := ret[0].(*types.SingleBytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportAccount indicates an expected call of ExportAccount
func (mr *MockAergoRPCServiceClientMockRecorder) ExportAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportAccount", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).ExportAccount), varargs...)
}

// GetABI mocks base method
func (m *MockAergoRPCServiceClient) GetABI(arg0 context.Context, arg1 *types.SingleBytes, arg2 ...grpc.CallOption) (*types.ABI, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetABI", varargs...)
	ret0, _ := ret[0].(*types.ABI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetABI indicates an expected call of GetABI
func (mr *MockAergoRPCServiceClientMockRecorder) GetABI(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetABI", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetABI), varargs...)
}

// GetAccounts mocks base method
func (m *MockAergoRPCServiceClient) GetAccounts(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (*types.AccountList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccounts", varargs...)
	ret0, _ := ret[0].(*types.AccountList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts
func (mr *MockAergoRPCServiceClientMockRecorder) GetAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetAccounts), varargs...)
}

// GetBlock mocks base method
func (m *MockAergoRPCServiceClient) GetBlock(arg0 context.Context, arg1 *types.SingleBytes, arg2 ...grpc.CallOption) (*types.Block, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlock", varargs...)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock
func (mr *MockAergoRPCServiceClientMockRecorder) GetBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetBlock), varargs...)
}

// GetBlockBody mocks base method
func (m *MockAergoRPCServiceClient) GetBlockBody(arg0 context.Context, arg1 *types.BlockBodyParams, arg2 ...grpc.CallOption) (*types.BlockBodyPaged, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlockBody", varargs...)
	ret0, _ := ret[0].(*types.BlockBodyPaged)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockBody indicates an expected call of GetBlockBody
func (mr *MockAergoRPCServiceClientMockRecorder) GetBlockBody(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockBody", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetBlockBody), varargs...)
}

// GetBlockMetadata mocks base method
func (m *MockAergoRPCServiceClient) GetBlockMetadata(arg0 context.Context, arg1 *types.SingleBytes, arg2 ...grpc.CallOption) (*types.BlockMetadata, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlockMetadata", varargs...)
	ret0, _ := ret[0].(*types.BlockMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockMetadata indicates an expected call of GetBlockMetadata
func (mr *MockAergoRPCServiceClientMockRecorder) GetBlockMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockMetadata", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetBlockMetadata), varargs...)
}

// GetBlockTX mocks base method
func (m *MockAergoRPCServiceClient) GetBlockTX(arg0 context.Context, arg1 *types.SingleBytes, arg2 ...grpc.CallOption) (*types.TxInBlock, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlockTX", varargs...)
	ret0, _ := ret[0].(*types.TxInBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockTX indicates an expected call of GetBlockTX
func (mr *MockAergoRPCServiceClientMockRecorder) GetBlockTX(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockTX", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetBlockTX), varargs...)
}

// GetChainInfo mocks base method
func (m *MockAergoRPCServiceClient) GetChainInfo(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (*types.ChainInfo, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChainInfo", varargs...)
	ret0, _ := ret[0].(*types.ChainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainInfo indicates an expected call of GetChainInfo
func (mr *MockAergoRPCServiceClientMockRecorder) GetChainInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainInfo", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetChainInfo), varargs...)
}

// GetNameInfo mocks base method
func (m *MockAergoRPCServiceClient) GetNameInfo(arg0 context.Context, arg1 *types.Name, arg2 ...grpc.CallOption) (*types.NameInfo, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNameInfo", varargs...)
	ret0, _ := ret[0].(*types.NameInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNameInfo indicates an expected call of GetNameInfo
func (mr *MockAergoRPCServiceClientMockRecorder) GetNameInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNameInfo", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetNameInfo), varargs...)
}

// GetPeers mocks base method
func (m *MockAergoRPCServiceClient) GetPeers(arg0 context.Context, arg1 *types.PeersParams, arg2 ...grpc.CallOption) (*types.PeerList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPeers", varargs...)
	ret0, _ := ret[0].(*types.PeerList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeers indicates an expected call of GetPeers
func (mr *MockAergoRPCServiceClientMockRecorder) GetPeers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeers", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetPeers), varargs...)
}

// GetReceipt mocks base method
func (m *MockAergoRPCServiceClient) GetReceipt(arg0 context.Context, arg1 *types.SingleBytes, arg2 ...grpc.CallOption) (*types.Receipt, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReceipt", varargs...)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceipt indicates an expected call of GetReceipt
func (mr *MockAergoRPCServiceClientMockRecorder) GetReceipt(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipt", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetReceipt), varargs...)
}

// GetStaking mocks base method
func (m *MockAergoRPCServiceClient) GetStaking(arg0 context.Context, arg1 *types.SingleBytes, arg2 ...grpc.CallOption) (*types.Staking, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStaking", varargs...)
	ret0, _ := ret[0].(*types.Staking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStaking indicates an expected call of GetStaking
func (mr *MockAergoRPCServiceClientMockRecorder) GetStaking(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStaking", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetStaking), varargs...)
}

// GetState mocks base method
func (m *MockAergoRPCServiceClient) GetState(arg0 context.Context, arg1 *types.SingleBytes, arg2 ...grpc.CallOption) (*types.State, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetState", varargs...)
	ret0, _ := ret[0].(*types.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetState indicates an expected call of GetState
func (mr *MockAergoRPCServiceClientMockRecorder) GetState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetState), varargs...)
}

// GetStateAndProof mocks base method
func (m *MockAergoRPCServiceClient) GetStateAndProof(arg0 context.Context, arg1 *types.AccountAndRoot, arg2 ...grpc.CallOption) (*types.AccountProof, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStateAndProof", varargs...)
	ret0, _ := ret[0].(*types.AccountProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateAndProof indicates an expected call of GetStateAndProof
func (mr *MockAergoRPCServiceClientMockRecorder) GetStateAndProof(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateAndProof", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetStateAndProof), varargs...)
}

// GetTX mocks base method
func (m *MockAergoRPCServiceClient) GetTX(arg0 context.Context, arg1 *types.SingleBytes, arg2 ...grpc.CallOption) (*types.Tx, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTX", varargs...)
	ret0, _ := ret[0].(*types.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTX indicates an expected call of GetTX
func (mr *MockAergoRPCServiceClientMockRecorder) GetTX(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTX", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetTX), varargs...)
}

// GetVotes mocks base method
func (m *MockAergoRPCServiceClient) GetVotes(arg0 context.Context, arg1 *types.SingleBytes, arg2 ...grpc.CallOption) (*types.VoteList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVotes", varargs...)
	ret0, _ := ret[0].(*types.VoteList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVotes indicates an expected call of GetVotes
func (mr *MockAergoRPCServiceClientMockRecorder) GetVotes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotes", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).GetVotes), varargs...)
}

// ImportAccount mocks base method
func (m *MockAergoRPCServiceClient) ImportAccount(arg0 context.Context, arg1 *types.ImportFormat, arg2 ...grpc.CallOption) (*types.Account, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportAccount", varargs...)
	ret0, _ := ret[0].(*types.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportAccount indicates an expected call of ImportAccount
func (mr *MockAergoRPCServiceClientMockRecorder) ImportAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportAccount", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).ImportAccount), varargs...)
}

// ListBlockHeaders mocks base method
func (m *MockAergoRPCServiceClient) ListBlockHeaders(arg0 context.Context, arg1 *types.ListParams, arg2 ...grpc.CallOption) (*types.BlockHeaderList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBlockHeaders", varargs...)
	ret0, _ := ret[0].(*types.BlockHeaderList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBlockHeaders indicates an expected call of ListBlockHeaders
func (mr *MockAergoRPCServiceClientMockRecorder) ListBlockHeaders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBlockHeaders", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).ListBlockHeaders), varargs...)
}

// ListBlockMetadata mocks base method
func (m *MockAergoRPCServiceClient) ListBlockMetadata(arg0 context.Context, arg1 *types.ListParams, arg2 ...grpc.CallOption) (*types.BlockMetadataList, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBlockMetadata", varargs...)
	ret0, _ := ret[0].(*types.BlockMetadataList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBlockMetadata indicates an expected call of ListBlockMetadata
func (mr *MockAergoRPCServiceClientMockRecorder) ListBlockMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBlockMetadata", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).ListBlockMetadata), varargs...)
}

// ListBlockMetadataStream mocks base method
func (m *MockAergoRPCServiceClient) ListBlockMetadataStream(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (types.AergoRPCService_ListBlockMetadataStreamClient, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBlockMetadataStream", varargs...)
	ret0, _ := ret[0].(types.AergoRPCService_ListBlockMetadataStreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBlockMetadataStream indicates an expected call of ListBlockMetadataStream
func (mr *MockAergoRPCServiceClientMockRecorder) ListBlockMetadataStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBlockMetadataStream", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).ListBlockMetadataStream), varargs...)
}

// ListBlockStream mocks base method
func (m *MockAergoRPCServiceClient) ListBlockStream(arg0 context.Context, arg1 *types.Empty, arg2 ...grpc.CallOption) (types.AergoRPCService_ListBlockStreamClient, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBlockStream", varargs...)
	ret0, _ := ret[0].(types.AergoRPCService_ListBlockStreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBlockStream indicates an expected call of ListBlockStream
func (mr *MockAergoRPCServiceClientMockRecorder) ListBlockStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBlockStream", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).ListBlockStream), varargs...)
}

// LockAccount mocks base method
func (m *MockAergoRPCServiceClient) LockAccount(arg0 context.Context, arg1 *types.Personal, arg2 ...grpc.CallOption) (*types.Account, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LockAccount", varargs...)
	ret0, _ := ret[0].(*types.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LockAccount indicates an expected call of LockAccount
func (mr *MockAergoRPCServiceClientMockRecorder) LockAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockAccount", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).LockAccount), varargs...)
}

// Metric mocks base method
func (m *MockAergoRPCServiceClient) Metric(arg0 context.Context, arg1 *types.MetricsRequest, arg2 ...grpc.CallOption) (*types.Metrics, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Metric", varargs...)
	ret0, _ := ret[0].(*types.Metrics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Metric indicates an expected call of Metric
func (mr *MockAergoRPCServiceClientMockRecorder) Metric(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metric", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).Metric), varargs...)
}

// NodeState mocks base method
func (m *MockAergoRPCServiceClient) NodeState(arg0 context.Context, arg1 *types.NodeReq, arg2 ...grpc.CallOption) (*types.SingleBytes, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NodeState", varargs...)
	ret0, _ := ret[0].(*types.SingleBytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeState indicates an expected call of NodeState
func (mr *MockAergoRPCServiceClientMockRecorder) NodeState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeState", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).NodeState), varargs...)
}

// QueryContract mocks base method
func (m *MockAergoRPCServiceClient) QueryContract(arg0 context.Context, arg1 *types.Query, arg2 ...grpc.CallOption) (*types.SingleBytes, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContract", varargs...)
	ret0, _ := ret[0].(*types.SingleBytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContract indicates an expected call of QueryContract
func (mr *MockAergoRPCServiceClientMockRecorder) QueryContract(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContract", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).QueryContract), varargs...)
}

// QueryContractState mocks base method
func (m *MockAergoRPCServiceClient) QueryContractState(arg0 context.Context, arg1 *types.StateQuery, arg2 ...grpc.CallOption) (*types.StateQueryProof, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContractState", varargs...)
	ret0, _ := ret[0].(*types.StateQueryProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContractState indicates an expected call of QueryContractState
func (mr *MockAergoRPCServiceClientMockRecorder) QueryContractState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContractState", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).QueryContractState), varargs...)
}

// SendTX mocks base method
func (m *MockAergoRPCServiceClient) SendTX(arg0 context.Context, arg1 *types.Tx, arg2 ...grpc.CallOption) (*types.CommitResult, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendTX", varargs...)
	ret0, _ := ret[0].(*types.CommitResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTX indicates an expected call of SendTX
func (mr *MockAergoRPCServiceClientMockRecorder) SendTX(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTX", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).SendTX), varargs...)
}

// SignTX mocks base method
func (m *MockAergoRPCServiceClient) SignTX(arg0 context.Context, arg1 *types.Tx, arg2 ...grpc.CallOption) (*types.Tx, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignTX", varargs...)
	ret0, _ := ret[0].(*types.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignTX indicates an expected call of SignTX
func (mr *MockAergoRPCServiceClientMockRecorder) SignTX(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTX", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).SignTX), varargs...)
}

// UnlockAccount mocks base method
func (m *MockAergoRPCServiceClient) UnlockAccount(arg0 context.Context, arg1 *types.Personal, arg2 ...grpc.CallOption) (*types.Account, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnlockAccount", varargs...)
	ret0, _ := ret[0].(*types.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnlockAccount indicates an expected call of UnlockAccount
func (mr *MockAergoRPCServiceClientMockRecorder) UnlockAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockAccount", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).UnlockAccount), varargs...)
}

// VerifyTX mocks base method
func (m *MockAergoRPCServiceClient) VerifyTX(arg0 context.Context, arg1 *types.Tx, arg2 ...grpc.CallOption) (*types.VerifyResult, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VerifyTX", varargs...)
	ret0, _ := ret[0].(*types.VerifyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyTX indicates an expected call of VerifyTX
func (mr *MockAergoRPCServiceClientMockRecorder) VerifyTX(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyTX", reflect.TypeOf((*MockAergoRPCServiceClient)(nil).VerifyTX), varargs...)
}
