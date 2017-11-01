// Code generated by MockGen. DO NOT EDIT.
// Source: conn.go

package skydb

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockConn is a mock of Conn interface
type MockConn struct {
	ctrl     *gomock.Controller
	recorder *MockConnMockRecorder
}

// MockConnMockRecorder is the mock recorder for MockConn
type MockConnMockRecorder struct {
	mock *MockConn
}

// NewMockConn creates a new mock instance
func NewMockConn(ctrl *gomock.Controller) *MockConn {
	mock := &MockConn{ctrl: ctrl}
	mock.recorder = &MockConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockConn) EXPECT() *MockConnMockRecorder {
	return _m.recorder
}

// CreateAuth mocks base method
func (_m *MockConn) CreateAuth(authinfo *AuthInfo) error {
	ret := _m.ctrl.Call(_m, "CreateAuth", authinfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAuth indicates an expected call of CreateAuth
func (_mr *MockConnMockRecorder) CreateAuth(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateAuth", reflect.TypeOf((*MockConn)(nil).CreateAuth), arg0)
}

// GetAuth mocks base method
func (_m *MockConn) GetAuth(id string, authinfo *AuthInfo) error {
	ret := _m.ctrl.Call(_m, "GetAuth", id, authinfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAuth indicates an expected call of GetAuth
func (_mr *MockConnMockRecorder) GetAuth(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAuth", reflect.TypeOf((*MockConn)(nil).GetAuth), arg0, arg1)
}

// GetAuthByPrincipalID mocks base method
func (_m *MockConn) GetAuthByPrincipalID(principalID string, authinfo *AuthInfo) error {
	ret := _m.ctrl.Call(_m, "GetAuthByPrincipalID", principalID, authinfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAuthByPrincipalID indicates an expected call of GetAuthByPrincipalID
func (_mr *MockConnMockRecorder) GetAuthByPrincipalID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAuthByPrincipalID", reflect.TypeOf((*MockConn)(nil).GetAuthByPrincipalID), arg0, arg1)
}

// UpdateAuth mocks base method
func (_m *MockConn) UpdateAuth(authinfo *AuthInfo) error {
	ret := _m.ctrl.Call(_m, "UpdateAuth", authinfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuth indicates an expected call of UpdateAuth
func (_mr *MockConnMockRecorder) UpdateAuth(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateAuth", reflect.TypeOf((*MockConn)(nil).UpdateAuth), arg0)
}

// DeleteAuth mocks base method
func (_m *MockConn) DeleteAuth(id string) error {
	ret := _m.ctrl.Call(_m, "DeleteAuth", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuth indicates an expected call of DeleteAuth
func (_mr *MockConnMockRecorder) DeleteAuth(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteAuth", reflect.TypeOf((*MockConn)(nil).DeleteAuth), arg0)
}

// GetAdminRoles mocks base method
func (_m *MockConn) GetAdminRoles() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetAdminRoles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdminRoles indicates an expected call of GetAdminRoles
func (_mr *MockConnMockRecorder) GetAdminRoles() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAdminRoles", reflect.TypeOf((*MockConn)(nil).GetAdminRoles))
}

// SetAdminRoles mocks base method
func (_m *MockConn) SetAdminRoles(roles []string) error {
	ret := _m.ctrl.Call(_m, "SetAdminRoles", roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAdminRoles indicates an expected call of SetAdminRoles
func (_mr *MockConnMockRecorder) SetAdminRoles(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetAdminRoles", reflect.TypeOf((*MockConn)(nil).SetAdminRoles), arg0)
}

// GetDefaultRoles mocks base method
func (_m *MockConn) GetDefaultRoles() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetDefaultRoles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultRoles indicates an expected call of GetDefaultRoles
func (_mr *MockConnMockRecorder) GetDefaultRoles() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetDefaultRoles", reflect.TypeOf((*MockConn)(nil).GetDefaultRoles))
}

// SetDefaultRoles mocks base method
func (_m *MockConn) SetDefaultRoles(roles []string) error {
	ret := _m.ctrl.Call(_m, "SetDefaultRoles", roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDefaultRoles indicates an expected call of SetDefaultRoles
func (_mr *MockConnMockRecorder) SetDefaultRoles(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetDefaultRoles", reflect.TypeOf((*MockConn)(nil).SetDefaultRoles), arg0)
}

// AssignRoles mocks base method
func (_m *MockConn) AssignRoles(userIDs []string, roles []string) error {
	ret := _m.ctrl.Call(_m, "AssignRoles", userIDs, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignRoles indicates an expected call of AssignRoles
func (_mr *MockConnMockRecorder) AssignRoles(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "AssignRoles", reflect.TypeOf((*MockConn)(nil).AssignRoles), arg0, arg1)
}

// RevokeRoles mocks base method
func (_m *MockConn) RevokeRoles(userIDs []string, roles []string) error {
	ret := _m.ctrl.Call(_m, "RevokeRoles", userIDs, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeRoles indicates an expected call of RevokeRoles
func (_mr *MockConnMockRecorder) RevokeRoles(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RevokeRoles", reflect.TypeOf((*MockConn)(nil).RevokeRoles), arg0, arg1)
}

// GetRoles mocks base method
func (_m *MockConn) GetRoles(userIDs []string) (map[string][]string, error) {
	ret := _m.ctrl.Call(_m, "GetRoles", userIDs)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoles indicates an expected call of GetRoles
func (_mr *MockConnMockRecorder) GetRoles(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetRoles", reflect.TypeOf((*MockConn)(nil).GetRoles), arg0)
}

// SetRecordAccess mocks base method
func (_m *MockConn) SetRecordAccess(recordType string, acl RecordACL) error {
	ret := _m.ctrl.Call(_m, "SetRecordAccess", recordType, acl)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRecordAccess indicates an expected call of SetRecordAccess
func (_mr *MockConnMockRecorder) SetRecordAccess(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRecordAccess", reflect.TypeOf((*MockConn)(nil).SetRecordAccess), arg0, arg1)
}

// SetRecordDefaultAccess mocks base method
func (_m *MockConn) SetRecordDefaultAccess(recordType string, acl RecordACL) error {
	ret := _m.ctrl.Call(_m, "SetRecordDefaultAccess", recordType, acl)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRecordDefaultAccess indicates an expected call of SetRecordDefaultAccess
func (_mr *MockConnMockRecorder) SetRecordDefaultAccess(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRecordDefaultAccess", reflect.TypeOf((*MockConn)(nil).SetRecordDefaultAccess), arg0, arg1)
}

// GetRecordAccess mocks base method
func (_m *MockConn) GetRecordAccess(recordType string) (RecordACL, error) {
	ret := _m.ctrl.Call(_m, "GetRecordAccess", recordType)
	ret0, _ := ret[0].(RecordACL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordAccess indicates an expected call of GetRecordAccess
func (_mr *MockConnMockRecorder) GetRecordAccess(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetRecordAccess", reflect.TypeOf((*MockConn)(nil).GetRecordAccess), arg0)
}

// GetRecordDefaultAccess mocks base method
func (_m *MockConn) GetRecordDefaultAccess(recordType string) (RecordACL, error) {
	ret := _m.ctrl.Call(_m, "GetRecordDefaultAccess", recordType)
	ret0, _ := ret[0].(RecordACL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordDefaultAccess indicates an expected call of GetRecordDefaultAccess
func (_mr *MockConnMockRecorder) GetRecordDefaultAccess(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetRecordDefaultAccess", reflect.TypeOf((*MockConn)(nil).GetRecordDefaultAccess), arg0)
}

// SetRecordFieldAccess mocks base method
func (_m *MockConn) SetRecordFieldAccess(acl FieldACL) error {
	ret := _m.ctrl.Call(_m, "SetRecordFieldAccess", acl)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRecordFieldAccess indicates an expected call of SetRecordFieldAccess
func (_mr *MockConnMockRecorder) SetRecordFieldAccess(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRecordFieldAccess", reflect.TypeOf((*MockConn)(nil).SetRecordFieldAccess), arg0)
}

// GetRecordFieldAccess mocks base method
func (_m *MockConn) GetRecordFieldAccess() (FieldACL, error) {
	ret := _m.ctrl.Call(_m, "GetRecordFieldAccess")
	ret0, _ := ret[0].(FieldACL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordFieldAccess indicates an expected call of GetRecordFieldAccess
func (_mr *MockConnMockRecorder) GetRecordFieldAccess() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetRecordFieldAccess", reflect.TypeOf((*MockConn)(nil).GetRecordFieldAccess))
}

// GetAsset mocks base method
func (_m *MockConn) GetAsset(name string, asset *Asset) error {
	ret := _m.ctrl.Call(_m, "GetAsset", name, asset)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAsset indicates an expected call of GetAsset
func (_mr *MockConnMockRecorder) GetAsset(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAsset", reflect.TypeOf((*MockConn)(nil).GetAsset), arg0, arg1)
}

// GetAssets mocks base method
func (_m *MockConn) GetAssets(names []string) ([]Asset, error) {
	ret := _m.ctrl.Call(_m, "GetAssets", names)
	ret0, _ := ret[0].([]Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssets indicates an expected call of GetAssets
func (_mr *MockConnMockRecorder) GetAssets(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAssets", reflect.TypeOf((*MockConn)(nil).GetAssets), arg0)
}

// SaveAsset mocks base method
func (_m *MockConn) SaveAsset(asset *Asset) error {
	ret := _m.ctrl.Call(_m, "SaveAsset", asset)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAsset indicates an expected call of SaveAsset
func (_mr *MockConnMockRecorder) SaveAsset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SaveAsset", reflect.TypeOf((*MockConn)(nil).SaveAsset), arg0)
}

// QueryRelation mocks base method
func (_m *MockConn) QueryRelation(user string, name string, direction string, config QueryConfig) []AuthInfo {
	ret := _m.ctrl.Call(_m, "QueryRelation", user, name, direction, config)
	ret0, _ := ret[0].([]AuthInfo)
	return ret0
}

// QueryRelation indicates an expected call of QueryRelation
func (_mr *MockConnMockRecorder) QueryRelation(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "QueryRelation", reflect.TypeOf((*MockConn)(nil).QueryRelation), arg0, arg1, arg2, arg3)
}

// QueryRelationCount mocks base method
func (_m *MockConn) QueryRelationCount(user string, name string, direction string) (uint64, error) {
	ret := _m.ctrl.Call(_m, "QueryRelationCount", user, name, direction)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRelationCount indicates an expected call of QueryRelationCount
func (_mr *MockConnMockRecorder) QueryRelationCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "QueryRelationCount", reflect.TypeOf((*MockConn)(nil).QueryRelationCount), arg0, arg1, arg2)
}

// AddRelation mocks base method
func (_m *MockConn) AddRelation(user string, name string, targetUser string) error {
	ret := _m.ctrl.Call(_m, "AddRelation", user, name, targetUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRelation indicates an expected call of AddRelation
func (_mr *MockConnMockRecorder) AddRelation(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "AddRelation", reflect.TypeOf((*MockConn)(nil).AddRelation), arg0, arg1, arg2)
}

// RemoveRelation mocks base method
func (_m *MockConn) RemoveRelation(user string, name string, targetUser string) error {
	ret := _m.ctrl.Call(_m, "RemoveRelation", user, name, targetUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRelation indicates an expected call of RemoveRelation
func (_mr *MockConnMockRecorder) RemoveRelation(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RemoveRelation", reflect.TypeOf((*MockConn)(nil).RemoveRelation), arg0, arg1, arg2)
}

// GetDevice mocks base method
func (_m *MockConn) GetDevice(id string, device *Device) error {
	ret := _m.ctrl.Call(_m, "GetDevice", id, device)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetDevice indicates an expected call of GetDevice
func (_mr *MockConnMockRecorder) GetDevice(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetDevice", reflect.TypeOf((*MockConn)(nil).GetDevice), arg0, arg1)
}

// QueryDevicesByUser mocks base method
func (_m *MockConn) QueryDevicesByUser(user string) ([]Device, error) {
	ret := _m.ctrl.Call(_m, "QueryDevicesByUser", user)
	ret0, _ := ret[0].([]Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryDevicesByUser indicates an expected call of QueryDevicesByUser
func (_mr *MockConnMockRecorder) QueryDevicesByUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "QueryDevicesByUser", reflect.TypeOf((*MockConn)(nil).QueryDevicesByUser), arg0)
}

// QueryDevicesByUserAndTopic mocks base method
func (_m *MockConn) QueryDevicesByUserAndTopic(user string, topic string) ([]Device, error) {
	ret := _m.ctrl.Call(_m, "QueryDevicesByUserAndTopic", user, topic)
	ret0, _ := ret[0].([]Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryDevicesByUserAndTopic indicates an expected call of QueryDevicesByUserAndTopic
func (_mr *MockConnMockRecorder) QueryDevicesByUserAndTopic(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "QueryDevicesByUserAndTopic", reflect.TypeOf((*MockConn)(nil).QueryDevicesByUserAndTopic), arg0, arg1)
}

// SaveDevice mocks base method
func (_m *MockConn) SaveDevice(device *Device) error {
	ret := _m.ctrl.Call(_m, "SaveDevice", device)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveDevice indicates an expected call of SaveDevice
func (_mr *MockConnMockRecorder) SaveDevice(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SaveDevice", reflect.TypeOf((*MockConn)(nil).SaveDevice), arg0)
}

// DeleteDevice mocks base method
func (_m *MockConn) DeleteDevice(id string) error {
	ret := _m.ctrl.Call(_m, "DeleteDevice", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDevice indicates an expected call of DeleteDevice
func (_mr *MockConnMockRecorder) DeleteDevice(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteDevice", reflect.TypeOf((*MockConn)(nil).DeleteDevice), arg0)
}

// DeleteDevicesByToken mocks base method
func (_m *MockConn) DeleteDevicesByToken(token string, t time.Time) error {
	ret := _m.ctrl.Call(_m, "DeleteDevicesByToken", token, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDevicesByToken indicates an expected call of DeleteDevicesByToken
func (_mr *MockConnMockRecorder) DeleteDevicesByToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteDevicesByToken", reflect.TypeOf((*MockConn)(nil).DeleteDevicesByToken), arg0, arg1)
}

// DeleteEmptyDevicesByTime mocks base method
func (_m *MockConn) DeleteEmptyDevicesByTime(t time.Time) error {
	ret := _m.ctrl.Call(_m, "DeleteEmptyDevicesByTime", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmptyDevicesByTime indicates an expected call of DeleteEmptyDevicesByTime
func (_mr *MockConnMockRecorder) DeleteEmptyDevicesByTime(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteEmptyDevicesByTime", reflect.TypeOf((*MockConn)(nil).DeleteEmptyDevicesByTime), arg0)
}

// PublicDB mocks base method
func (_m *MockConn) PublicDB() Database {
	ret := _m.ctrl.Call(_m, "PublicDB")
	ret0, _ := ret[0].(Database)
	return ret0
}

// PublicDB indicates an expected call of PublicDB
func (_mr *MockConnMockRecorder) PublicDB() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "PublicDB", reflect.TypeOf((*MockConn)(nil).PublicDB))
}

// PrivateDB mocks base method
func (_m *MockConn) PrivateDB(userKey string) Database {
	ret := _m.ctrl.Call(_m, "PrivateDB", userKey)
	ret0, _ := ret[0].(Database)
	return ret0
}

// PrivateDB indicates an expected call of PrivateDB
func (_mr *MockConnMockRecorder) PrivateDB(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "PrivateDB", reflect.TypeOf((*MockConn)(nil).PrivateDB), arg0)
}

// UnionDB mocks base method
func (_m *MockConn) UnionDB() Database {
	ret := _m.ctrl.Call(_m, "UnionDB")
	ret0, _ := ret[0].(Database)
	return ret0
}

// UnionDB indicates an expected call of UnionDB
func (_mr *MockConnMockRecorder) UnionDB() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UnionDB", reflect.TypeOf((*MockConn)(nil).UnionDB))
}

// Subscribe mocks base method
func (_m *MockConn) Subscribe(recordEventChan chan RecordEvent) error {
	ret := _m.ctrl.Call(_m, "Subscribe", recordEventChan)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (_mr *MockConnMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Subscribe", reflect.TypeOf((*MockConn)(nil).Subscribe), arg0)
}

// EnsureAuthRecordKeysExist mocks base method
func (_m *MockConn) EnsureAuthRecordKeysExist(authRecordKeys [][]string) error {
	ret := _m.ctrl.Call(_m, "EnsureAuthRecordKeysExist", authRecordKeys)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureAuthRecordKeysExist indicates an expected call of EnsureAuthRecordKeysExist
func (_mr *MockConnMockRecorder) EnsureAuthRecordKeysExist(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "EnsureAuthRecordKeysExist", reflect.TypeOf((*MockConn)(nil).EnsureAuthRecordKeysExist), arg0)
}

// EnsureAuthRecordKeysIndexesMatch mocks base method
func (_m *MockConn) EnsureAuthRecordKeysIndexesMatch(authRecordKeys [][]string) error {
	ret := _m.ctrl.Call(_m, "EnsureAuthRecordKeysIndexesMatch", authRecordKeys)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureAuthRecordKeysIndexesMatch indicates an expected call of EnsureAuthRecordKeysIndexesMatch
func (_mr *MockConnMockRecorder) EnsureAuthRecordKeysIndexesMatch(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "EnsureAuthRecordKeysIndexesMatch", reflect.TypeOf((*MockConn)(nil).EnsureAuthRecordKeysIndexesMatch), arg0)
}

// CreateOAuthInfo mocks base method
func (_m *MockConn) CreateOAuthInfo(oauthinfo *OAuthInfo) error {
	ret := _m.ctrl.Call(_m, "CreateOAuthInfo", oauthinfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOAuthInfo indicates an expected call of CreateOAuthInfo
func (_mr *MockConnMockRecorder) CreateOAuthInfo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateOAuthInfo", reflect.TypeOf((*MockConn)(nil).CreateOAuthInfo), arg0)
}

// GetOAuthInfo mocks base method
func (_m *MockConn) GetOAuthInfo(provider string, principalID string, oauthinfo *OAuthInfo) error {
	ret := _m.ctrl.Call(_m, "GetOAuthInfo", provider, principalID, oauthinfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetOAuthInfo indicates an expected call of GetOAuthInfo
func (_mr *MockConnMockRecorder) GetOAuthInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetOAuthInfo", reflect.TypeOf((*MockConn)(nil).GetOAuthInfo), arg0, arg1, arg2)
}

// GetOAuthInfoByProvicerAndUserID mocks base method
func (_m *MockConn) GetOAuthInfoByProvicerAndUserID(provider string, userID string, oauthinfo *OAuthInfo) error {
	ret := _m.ctrl.Call(_m, "GetOAuthInfoByProvicerAndUserID", provider, userID, oauthinfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetOAuthInfoByProvicerAndUserID indicates an expected call of GetOAuthInfoByProvicerAndUserID
func (_mr *MockConnMockRecorder) GetOAuthInfoByProvicerAndUserID(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetOAuthInfoByProvicerAndUserID", reflect.TypeOf((*MockConn)(nil).GetOAuthInfoByProvicerAndUserID), arg0, arg1, arg2)
}

// UpdateOAuthInfo mocks base method
func (_m *MockConn) UpdateOAuthInfo(oauthinfo *OAuthInfo) error {
	ret := _m.ctrl.Call(_m, "UpdateOAuthInfo", oauthinfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOAuthInfo indicates an expected call of UpdateOAuthInfo
func (_mr *MockConnMockRecorder) UpdateOAuthInfo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateOAuthInfo", reflect.TypeOf((*MockConn)(nil).UpdateOAuthInfo), arg0)
}

// DeleteOAuth mocks base method
func (_m *MockConn) DeleteOAuth(provider string, principalID string) error {
	ret := _m.ctrl.Call(_m, "DeleteOAuth", provider, principalID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOAuth indicates an expected call of DeleteOAuth
func (_mr *MockConnMockRecorder) DeleteOAuth(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteOAuth", reflect.TypeOf((*MockConn)(nil).DeleteOAuth), arg0, arg1)
}

// Close mocks base method
func (_m *MockConn) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockConnMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockConn)(nil).Close))
}
