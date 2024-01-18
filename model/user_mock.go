// Code generated by MockGen. DO NOT EDIT.
// Source: ./model/interface.go

// Package model is a generated GoMock package.
package model

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/jiebutech/uc/types"
)

// MockEntity is a mock of Entity interface.
type MockEntity struct {
	ctrl     *gomock.Controller
	recorder *MockEntityMockRecorder
}

// MockEntityMockRecorder is the mock recorder for MockEntity.
type MockEntityMockRecorder struct {
	mock *MockEntity
}

// NewMockEntity creates a new mock instance.
func NewMockEntity(ctrl *gomock.Controller) *MockEntity {
	mock := &MockEntity{ctrl: ctrl}
	mock.recorder = &MockEntityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntity) EXPECT() *MockEntityMockRecorder {
	return m.recorder
}

// TableName mocks base method.
func (m *MockEntity) TableName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TableName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TableName indicates an expected call of TableName.
func (mr *MockEntityMockRecorder) TableName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TableName", reflect.TypeOf((*MockEntity)(nil).TableName))
}

// MockUserEntity is a mock of UserEntity interface.
type MockUserEntity struct {
	ctrl     *gomock.Controller
	recorder *MockUserEntityMockRecorder
}

// MockUserEntityMockRecorder is the mock recorder for MockUserEntity.
type MockUserEntityMockRecorder struct {
	mock *MockUserEntity
}

// NewMockUserEntity creates a new mock instance.
func NewMockUserEntity(ctrl *gomock.Controller) *MockUserEntity {
	mock := &MockUserEntity{ctrl: ctrl}
	mock.recorder = &MockUserEntityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserEntity) EXPECT() *MockUserEntityMockRecorder {
	return m.recorder
}

// AppKey mocks base method.
func (m *MockUserEntity) AppKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// AppKey indicates an expected call of AppKey.
func (mr *MockUserEntityMockRecorder) AppKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppKey", reflect.TypeOf((*MockUserEntity)(nil).AppKey))
}

// GetApp mocks base method.
func (m *MockUserEntity) GetApp() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApp")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetApp indicates an expected call of GetApp.
func (mr *MockUserEntityMockRecorder) GetApp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApp", reflect.TypeOf((*MockUserEntity)(nil).GetApp))
}

// GetAvatar mocks base method.
func (m *MockUserEntity) GetAvatar() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvatar")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAvatar indicates an expected call of GetAvatar.
func (mr *MockUserEntityMockRecorder) GetAvatar() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatar", reflect.TypeOf((*MockUserEntity)(nil).GetAvatar))
}

// GetID mocks base method.
func (m *MockUserEntity) GetID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockUserEntityMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockUserEntity)(nil).GetID))
}

// GetIdentify mocks base method.
func (m *MockUserEntity) GetIdentify() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdentify")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetIdentify indicates an expected call of GetIdentify.
func (mr *MockUserEntityMockRecorder) GetIdentify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdentify", reflect.TypeOf((*MockUserEntity)(nil).GetIdentify))
}

// GetLoginType mocks base method.
func (m *MockUserEntity) GetLoginType() types.LoginType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoginType")
	ret0, _ := ret[0].(types.LoginType)
	return ret0
}

// GetLoginType indicates an expected call of GetLoginType.
func (mr *MockUserEntityMockRecorder) GetLoginType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoginType", reflect.TypeOf((*MockUserEntity)(nil).GetLoginType))
}

// GetNickname mocks base method.
func (m *MockUserEntity) GetNickname() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNickname")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNickname indicates an expected call of GetNickname.
func (mr *MockUserEntityMockRecorder) GetNickname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNickname", reflect.TypeOf((*MockUserEntity)(nil).GetNickname))
}

// GetPassword mocks base method.
func (m *MockUserEntity) GetPassword() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassword")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPassword indicates an expected call of GetPassword.
func (mr *MockUserEntityMockRecorder) GetPassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockUserEntity)(nil).GetPassword))
}

// GetUserName mocks base method.
func (m *MockUserEntity) GetUserName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUserName indicates an expected call of GetUserName.
func (mr *MockUserEntityMockRecorder) GetUserName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserName", reflect.TypeOf((*MockUserEntity)(nil).GetUserName))
}

// IdKey mocks base method.
func (m *MockUserEntity) IdKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// IdKey indicates an expected call of IdKey.
func (mr *MockUserEntityMockRecorder) IdKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdKey", reflect.TypeOf((*MockUserEntity)(nil).IdKey))
}

// IdentifyKey mocks base method.
func (m *MockUserEntity) IdentifyKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdentifyKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// IdentifyKey indicates an expected call of IdentifyKey.
func (mr *MockUserEntityMockRecorder) IdentifyKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdentifyKey", reflect.TypeOf((*MockUserEntity)(nil).IdentifyKey))
}

// LoginTypeKey mocks base method.
func (m *MockUserEntity) LoginTypeKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginTypeKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// LoginTypeKey indicates an expected call of LoginTypeKey.
func (mr *MockUserEntityMockRecorder) LoginTypeKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginTypeKey", reflect.TypeOf((*MockUserEntity)(nil).LoginTypeKey))
}

// PasswordKey mocks base method.
func (m *MockUserEntity) PasswordKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// PasswordKey indicates an expected call of PasswordKey.
func (mr *MockUserEntityMockRecorder) PasswordKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordKey", reflect.TypeOf((*MockUserEntity)(nil).PasswordKey))
}

// SetApp mocks base method.
func (m *MockUserEntity) SetApp(string2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetApp", string2)
}

// SetApp indicates an expected call of SetApp.
func (mr *MockUserEntityMockRecorder) SetApp(string2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetApp", reflect.TypeOf((*MockUserEntity)(nil).SetApp), string2)
}

// SetAvatar mocks base method.
func (m *MockUserEntity) SetAvatar(avatar string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAvatar", avatar)
}

// SetAvatar indicates an expected call of SetAvatar.
func (mr *MockUserEntityMockRecorder) SetAvatar(avatar interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAvatar", reflect.TypeOf((*MockUserEntity)(nil).SetAvatar), avatar)
}

// SetId mocks base method.
func (m *MockUserEntity) SetId(id int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetId", id)
}

// SetId indicates an expected call of SetId.
func (mr *MockUserEntityMockRecorder) SetId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetId", reflect.TypeOf((*MockUserEntity)(nil).SetId), id)
}

// SetIdentify mocks base method.
func (m *MockUserEntity) SetIdentify(identify string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIdentify", identify)
}

// SetIdentify indicates an expected call of SetIdentify.
func (mr *MockUserEntityMockRecorder) SetIdentify(identify interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIdentify", reflect.TypeOf((*MockUserEntity)(nil).SetIdentify), identify)
}

// SetLoginType mocks base method.
func (m *MockUserEntity) SetLoginType(loginType types.LoginType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLoginType", loginType)
}

// SetLoginType indicates an expected call of SetLoginType.
func (mr *MockUserEntityMockRecorder) SetLoginType(loginType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLoginType", reflect.TypeOf((*MockUserEntity)(nil).SetLoginType), loginType)
}

// SetNickname mocks base method.
func (m *MockUserEntity) SetNickname(nickname string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNickname", nickname)
}

// SetNickname indicates an expected call of SetNickname.
func (mr *MockUserEntityMockRecorder) SetNickname(nickname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNickname", reflect.TypeOf((*MockUserEntity)(nil).SetNickname), nickname)
}

// SetPassword mocks base method.
func (m *MockUserEntity) SetPassword(password string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPassword", password)
}

// SetPassword indicates an expected call of SetPassword.
func (mr *MockUserEntityMockRecorder) SetPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPassword", reflect.TypeOf((*MockUserEntity)(nil).SetPassword), password)
}

// SetUsername mocks base method.
func (m *MockUserEntity) SetUsername(username string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUsername", username)
}

// SetUsername indicates an expected call of SetUsername.
func (mr *MockUserEntityMockRecorder) SetUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUsername", reflect.TypeOf((*MockUserEntity)(nil).SetUsername), username)
}

// TableName mocks base method.
func (m *MockUserEntity) TableName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TableName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TableName indicates an expected call of TableName.
func (mr *MockUserEntityMockRecorder) TableName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TableName", reflect.TypeOf((*MockUserEntity)(nil).TableName))
}

// ToMap mocks base method.
func (m *MockUserEntity) ToMap() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToMap")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// ToMap indicates an expected call of ToMap.
func (mr *MockUserEntityMockRecorder) ToMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToMap", reflect.TypeOf((*MockUserEntity)(nil).ToMap))
}

// UsernameKey mocks base method.
func (m *MockUserEntity) UsernameKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsernameKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// UsernameKey indicates an expected call of UsernameKey.
func (mr *MockUserEntityMockRecorder) UsernameKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsernameKey", reflect.TypeOf((*MockUserEntity)(nil).UsernameKey))
}

// MockOauthUserEntity is a mock of OauthUserEntity interface.
type MockOauthUserEntity struct {
	ctrl     *gomock.Controller
	recorder *MockOauthUserEntityMockRecorder
}

// MockOauthUserEntityMockRecorder is the mock recorder for MockOauthUserEntity.
type MockOauthUserEntityMockRecorder struct {
	mock *MockOauthUserEntity
}

// NewMockOauthUserEntity creates a new mock instance.
func NewMockOauthUserEntity(ctrl *gomock.Controller) *MockOauthUserEntity {
	mock := &MockOauthUserEntity{ctrl: ctrl}
	mock.recorder = &MockOauthUserEntityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOauthUserEntity) EXPECT() *MockOauthUserEntityMockRecorder {
	return m.recorder
}

// AppKey mocks base method.
func (m *MockOauthUserEntity) AppKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// AppKey indicates an expected call of AppKey.
func (mr *MockOauthUserEntityMockRecorder) AppKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppKey", reflect.TypeOf((*MockOauthUserEntity)(nil).AppKey))
}

// GetApp mocks base method.
func (m *MockOauthUserEntity) GetApp() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApp")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetApp indicates an expected call of GetApp.
func (mr *MockOauthUserEntityMockRecorder) GetApp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApp", reflect.TypeOf((*MockOauthUserEntity)(nil).GetApp))
}

// GetBindUserId mocks base method.
func (m *MockOauthUserEntity) GetBindUserId() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBindUserId")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetBindUserId indicates an expected call of GetBindUserId.
func (mr *MockOauthUserEntityMockRecorder) GetBindUserId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBindUserId", reflect.TypeOf((*MockOauthUserEntity)(nil).GetBindUserId))
}

// GetOpenid mocks base method.
func (m *MockOauthUserEntity) GetOpenid() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenid")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOpenid indicates an expected call of GetOpenid.
func (mr *MockOauthUserEntityMockRecorder) GetOpenid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenid", reflect.TypeOf((*MockOauthUserEntity)(nil).GetOpenid))
}

// LoginTypeKey mocks base method.
func (m *MockOauthUserEntity) LoginTypeKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginTypeKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// LoginTypeKey indicates an expected call of LoginTypeKey.
func (mr *MockOauthUserEntityMockRecorder) LoginTypeKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginTypeKey", reflect.TypeOf((*MockOauthUserEntity)(nil).LoginTypeKey))
}

// OpenidKey mocks base method.
func (m *MockOauthUserEntity) OpenidKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenidKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// OpenidKey indicates an expected call of OpenidKey.
func (mr *MockOauthUserEntityMockRecorder) OpenidKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenidKey", reflect.TypeOf((*MockOauthUserEntity)(nil).OpenidKey))
}

// SetApp mocks base method.
func (m *MockOauthUserEntity) SetApp(string2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetApp", string2)
}

// SetApp indicates an expected call of SetApp.
func (mr *MockOauthUserEntityMockRecorder) SetApp(string2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetApp", reflect.TypeOf((*MockOauthUserEntity)(nil).SetApp), string2)
}

// SetBindUserId mocks base method.
func (m *MockOauthUserEntity) SetBindUserId(uuserid int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBindUserId", uuserid)
}

// SetBindUserId indicates an expected call of SetBindUserId.
func (mr *MockOauthUserEntityMockRecorder) SetBindUserId(uuserid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBindUserId", reflect.TypeOf((*MockOauthUserEntity)(nil).SetBindUserId), uuserid)
}

// SetLoginType mocks base method.
func (m *MockOauthUserEntity) SetLoginType(loginType types.LoginType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLoginType", loginType)
}

// SetLoginType indicates an expected call of SetLoginType.
func (mr *MockOauthUserEntityMockRecorder) SetLoginType(loginType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLoginType", reflect.TypeOf((*MockOauthUserEntity)(nil).SetLoginType), loginType)
}

// SetOpenid mocks base method.
func (m *MockOauthUserEntity) SetOpenid(openid string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOpenid", openid)
}

// SetOpenid indicates an expected call of SetOpenid.
func (mr *MockOauthUserEntityMockRecorder) SetOpenid(openid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOpenid", reflect.TypeOf((*MockOauthUserEntity)(nil).SetOpenid), openid)
}

// TableName mocks base method.
func (m *MockOauthUserEntity) TableName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TableName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TableName indicates an expected call of TableName.
func (mr *MockOauthUserEntityMockRecorder) TableName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TableName", reflect.TypeOf((*MockOauthUserEntity)(nil).TableName))
}

// MockUserResource is a mock of UserResource interface.
type MockUserResource struct {
	ctrl     *gomock.Controller
	recorder *MockUserResourceMockRecorder
}

// MockUserResourceMockRecorder is the mock recorder for MockUserResource.
type MockUserResourceMockRecorder struct {
	mock *MockUserResource
}

// NewMockUserResource creates a new mock instance.
func NewMockUserResource(ctrl *gomock.Controller) *MockUserResource {
	mock := &MockUserResource{ctrl: ctrl}
	mock.recorder = &MockUserResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserResource) EXPECT() *MockUserResourceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserResource) CreateUser(dest Entity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserResourceMockRecorder) CreateUser(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserResource)(nil).CreateUser), dest)
}

// GenOauthUser mocks base method.
func (m *MockUserResource) GenOauthUser() OauthUserEntity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenOauthUser")
	ret0, _ := ret[0].(OauthUserEntity)
	return ret0
}

// GenOauthUser indicates an expected call of GenOauthUser.
func (mr *MockUserResourceMockRecorder) GenOauthUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenOauthUser", reflect.TypeOf((*MockUserResource)(nil).GenOauthUser))
}

// GenUser mocks base method.
func (m *MockUserResource) GenUser() UserEntity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenUser")
	ret0, _ := ret[0].(UserEntity)
	return ret0
}

// GenUser indicates an expected call of GenUser.
func (mr *MockUserResourceMockRecorder) GenUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenUser", reflect.TypeOf((*MockUserResource)(nil).GenUser))
}

// GetOauthByOpenid mocks base method.
func (m *MockUserResource) GetOauthByOpenid(dest OauthUserEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOauthByOpenid", dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetOauthByOpenid indicates an expected call of GetOauthByOpenid.
func (mr *MockUserResourceMockRecorder) GetOauthByOpenid(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOauthByOpenid", reflect.TypeOf((*MockUserResource)(nil).GetOauthByOpenid), dest)
}

// GetUserById mocks base method.
func (m *MockUserResource) GetUserById(dest UserEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserResourceMockRecorder) GetUserById(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserResource)(nil).GetUserById), dest)
}

// GetUserByIdentify mocks base method.
func (m *MockUserResource) GetUserByIdentify(dest UserEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByIdentify", dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUserByIdentify indicates an expected call of GetUserByIdentify.
func (mr *MockUserResourceMockRecorder) GetUserByIdentify(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByIdentify", reflect.TypeOf((*MockUserResource)(nil).GetUserByIdentify), dest)
}

// GetUserByUsername mocks base method.
func (m *MockUserResource) GetUserByUsername(dest UserEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockUserResourceMockRecorder) GetUserByUsername(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockUserResource)(nil).GetUserByUsername), dest)
}

// IsUserNotFound mocks base method.
func (m *MockUserResource) IsUserNotFound(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserNotFound", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUserNotFound indicates an expected call of IsUserNotFound.
func (mr *MockUserResourceMockRecorder) IsUserNotFound(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserNotFound", reflect.TypeOf((*MockUserResource)(nil).IsUserNotFound), err)
}

// SaveUser mocks base method.
func (m *MockUserResource) SaveUser(dest Entity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockUserResourceMockRecorder) SaveUser(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockUserResource)(nil).SaveUser), dest)
}

// TransactionCreate mocks base method.
func (m *MockUserResource) TransactionCreate(tablers map[Entity]func()) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionCreate", tablers)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransactionCreate indicates an expected call of TransactionCreate.
func (mr *MockUserResourceMockRecorder) TransactionCreate(tablers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionCreate", reflect.TypeOf((*MockUserResource)(nil).TransactionCreate), tablers)
}

// TransactionSave mocks base method.
func (m *MockUserResource) TransactionSave(tablers map[Entity]func()) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionSave", tablers)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransactionSave indicates an expected call of TransactionSave.
func (mr *MockUserResourceMockRecorder) TransactionSave(tablers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionSave", reflect.TypeOf((*MockUserResource)(nil).TransactionSave), tablers)
}

// UpdatePassword mocks base method.
func (m *MockUserResource) UpdatePassword(dest UserEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserResourceMockRecorder) UpdatePassword(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserResource)(nil).UpdatePassword), dest)
}