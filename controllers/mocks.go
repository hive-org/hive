// Code generated by MockGen. DO NOT EDIT.
// Source: main.go

// Package controllers is a generated GoMock package.
package controllers

import (
	models "hive/models"
	repositories "hive/repositories"
	context "context"
	gomock "github.com/golang/mock/gomock"
	go_uuid "github.com/satori/go.uuid"
	reflect "reflect"
)

// MockIController is a mock of IController interface
type MockIController struct {
	ctrl     *gomock.Controller
	recorder *MockIControllerMockRecorder
}

// MockIControllerMockRecorder is the mock recorder for MockIController
type MockIControllerMockRecorder struct {
	mock *MockIController
}

// NewMockIController creates a new mock instance
func NewMockIController(ctrl *gomock.Controller) *MockIController {
	mock := &MockIController{ctrl: ctrl}
	mock.recorder = &MockIControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIController) EXPECT() *MockIControllerMockRecorder {
	return m.recorder
}

// GetActualSecret mocks base method
func (m *MockIController) GetActualSecret(ctx context.Context) *models.Secret {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActualSecret", ctx)
	ret0, _ := ret[0].(*models.Secret)
	return ret0
}

// GetActualSecret indicates an expected call of GetActualSecret
func (mr *MockIControllerMockRecorder) GetActualSecret(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActualSecret", reflect.TypeOf((*MockIController)(nil).GetActualSecret), ctx)
}

// GetSecret mocks base method
func (m *MockIController) GetSecret(ctx context.Context, id go_uuid.UUID) *models.Secret {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", ctx, id)
	ret0, _ := ret[0].(*models.Secret)
	return ret0
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockIControllerMockRecorder) GetSecret(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockIController)(nil).GetSecret), ctx, id)
}

// CreateSession mocks base method
func (m *MockIController) CreateSession(ctx context.Context, userID go_uuid.UUID, fingerprint, userAgent string) *models.Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, userID, fingerprint, userAgent)
	ret0, _ := ret[0].(*models.Session)
	return ret0
}

// CreateSession indicates an expected call of CreateSession
func (mr *MockIControllerMockRecorder) CreateSession(ctx, userID, fingerprint, userAgent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockIController)(nil).CreateSession), ctx, userID, fingerprint, userAgent)
}

// UpdateSession mocks base method
func (m *MockIController) UpdateSession(ctx context.Context, id go_uuid.UUID, fingerprint, userAgent string) (int, *models.Session) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSession", ctx, id, fingerprint, userAgent)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Session)
	return ret0, ret1
}

// UpdateSession indicates an expected call of UpdateSession
func (mr *MockIControllerMockRecorder) UpdateSession(ctx, id, fingerprint, userAgent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSession", reflect.TypeOf((*MockIController)(nil).UpdateSession), ctx, id, fingerprint, userAgent)
}

// CreatePassword mocks base method
func (m *MockIController) CreatePassword(ctx context.Context, userId go_uuid.UUID, value string) (int, *models.Password) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePassword", ctx, userId, value)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Password)
	return ret0, ret1
}

// CreatePassword indicates an expected call of CreatePassword
func (mr *MockIControllerMockRecorder) CreatePassword(ctx, userId, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePassword", reflect.TypeOf((*MockIController)(nil).CreatePassword), ctx, userId, value)
}

// CreateEmailConfirmation mocks base method
func (m *MockIController) CreateEmailConfirmation(ctx context.Context, email string) (int, *models.EmailConfirmation) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmailConfirmation", ctx, email)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.EmailConfirmation)
	return ret0, ret1
}

// CreateEmailConfirmation indicates an expected call of CreateEmailConfirmation
func (mr *MockIControllerMockRecorder) CreateEmailConfirmation(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmailConfirmation", reflect.TypeOf((*MockIController)(nil).CreateEmailConfirmation), ctx, email)
}

// CreateEmail mocks base method
func (m *MockIController) CreateEmail(ctx context.Context, email, code string, userId go_uuid.UUID) (int, *models.Email) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmail", ctx, email, code, userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Email)
	return ret0, ret1
}

// CreateEmail indicates an expected call of CreateEmail
func (mr *MockIControllerMockRecorder) CreateEmail(ctx, email, code, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmail", reflect.TypeOf((*MockIController)(nil).CreateEmail), ctx, email, code, userId)
}

// CreateUser mocks base method
func (m *MockIController) CreateUser(ctx context.Context, password, email, emailCode, phone, phoneCode string) (int, *models.User) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, password, email, emailCode, phone, phoneCode)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.User)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockIControllerMockRecorder) CreateUser(ctx, password, email, emailCode, phone, phoneCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIController)(nil).CreateUser), ctx, password, email, emailCode, phone, phoneCode)
}

// GetUsers mocks base method
func (m *MockIController) GetUsers(ctx context.Context, query repositories.GetUsersQuery) []*models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, query)
	ret0, _ := ret[0].([]*models.User)
	return ret0
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockIControllerMockRecorder) GetUsers(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockIController)(nil).GetUsers), ctx, query)
}

// DeleteUser mocks base method
func (m *MockIController) DeleteUser(ctx context.Context, id go_uuid.UUID) (int, *models.User) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.User)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockIControllerMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIController)(nil).DeleteUser), ctx, id)
}

// GetUser mocks base method
func (m *MockIController) GetUser(ctx context.Context, id go_uuid.UUID) *models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(*models.User)
	return ret0
}

// GetUser indicates an expected call of GetUser
func (mr *MockIControllerMockRecorder) GetUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIController)(nil).GetUser), ctx, id)
}

// CreatePhoneConfirmation mocks base method
func (m *MockIController) CreatePhoneConfirmation(ctx context.Context, phone string) (int, *models.PhoneConfirmation) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePhoneConfirmation", ctx, phone)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.PhoneConfirmation)
	return ret0, ret1
}

// CreatePhoneConfirmation indicates an expected call of CreatePhoneConfirmation
func (mr *MockIControllerMockRecorder) CreatePhoneConfirmation(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePhoneConfirmation", reflect.TypeOf((*MockIController)(nil).CreatePhoneConfirmation), ctx, phone)
}

// CreatePhone mocks base method
func (m *MockIController) CreatePhone(ctx context.Context, phone, code string, userId go_uuid.UUID) (int, *models.Phone) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePhone", ctx, phone, code, userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Phone)
	return ret0, ret1
}

// CreatePhone indicates an expected call of CreatePhone
func (mr *MockIControllerMockRecorder) CreatePhone(ctx, phone, code, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePhone", reflect.TypeOf((*MockIController)(nil).CreatePhone), ctx, phone, code, userId)
}

// GetRole mocks base method
func (m *MockIController) GetRole(ctx context.Context, id go_uuid.UUID) (int, *models.Role) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", ctx, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Role)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole
func (mr *MockIControllerMockRecorder) GetRole(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockIController)(nil).GetRole), ctx, id)
}

// CreateRole mocks base method
func (m *MockIController) CreateRole(ctx context.Context, title string) (int, *models.Role) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", ctx, title)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Role)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole
func (mr *MockIControllerMockRecorder) CreateRole(ctx, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockIController)(nil).CreateRole), ctx, title)
}

// GetRoles mocks base method
func (m *MockIController) GetRoles(ctx context.Context, query repositories.GetRolesQuery) ([]*models.Role, *models.PaginationResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoles", ctx, query)
	ret0, _ := ret[0].([]*models.Role)
	ret1, _ := ret[1].(*models.PaginationResponse)
	return ret0, ret1
}

// GetRoles indicates an expected call of GetRoles
func (mr *MockIControllerMockRecorder) GetRoles(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoles", reflect.TypeOf((*MockIController)(nil).GetRoles), ctx, query)
}

// GetUserRoles mocks base method
func (m *MockIController) GetUserRoles(ctx context.Context, query repositories.GetUserRoleQuery) ([]*models.UserRole, *models.PaginationResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRoles", ctx, query)
	ret0, _ := ret[0].([]*models.UserRole)
	ret1, _ := ret[1].(*models.PaginationResponse)
	return ret0, ret1
}

// GetUserRoles indicates an expected call of GetUserRoles
func (mr *MockIControllerMockRecorder) GetUserRoles(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRoles", reflect.TypeOf((*MockIController)(nil).GetUserRoles), ctx, query)
}

// CreateUserRole mocks base method
func (m *MockIController) CreateUserRole(ctx context.Context, userId, roleID go_uuid.UUID) (int, *models.UserRole) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserRole", ctx, userId, roleID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.UserRole)
	return ret0, ret1
}

// CreateUserRole indicates an expected call of CreateUserRole
func (mr *MockIControllerMockRecorder) CreateUserRole(ctx, userId, roleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserRole", reflect.TypeOf((*MockIController)(nil).CreateUserRole), ctx, userId, roleID)
}

// DeleteUserRole mocks base method
func (m *MockIController) DeleteUserRole(ctx context.Context, id go_uuid.UUID) (int, *models.UserRole) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserRole", ctx, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.UserRole)
	return ret0, ret1
}

// DeleteUserRole indicates an expected call of DeleteUserRole
func (mr *MockIControllerMockRecorder) DeleteUserRole(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserRole", reflect.TypeOf((*MockIController)(nil).DeleteUserRole), ctx, id)
}

// CreateOrUpdateUsersView mocks base method
func (m *MockIController) CreateOrUpdateUsersView(ctx context.Context, id []go_uuid.UUID) []*models.UserView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateUsersView", ctx, id)
	ret0, _ := ret[0].([]*models.UserView)
	return ret0
}

// CreateOrUpdateUsersView indicates an expected call of CreateOrUpdateUsersView
func (mr *MockIControllerMockRecorder) CreateOrUpdateUsersView(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateUsersView", reflect.TypeOf((*MockIController)(nil).CreateOrUpdateUsersView), ctx, id)
}

// CreateOrUpdateUsersViewByRoles mocks base method
func (m *MockIController) CreateOrUpdateUsersViewByRoles(ctx context.Context, rolesIds []go_uuid.UUID) []*models.UserView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateUsersViewByRoles", ctx, rolesIds)
	ret0, _ := ret[0].([]*models.UserView)
	return ret0
}

// CreateOrUpdateUsersViewByRoles indicates an expected call of CreateOrUpdateUsersViewByRoles
func (mr *MockIControllerMockRecorder) CreateOrUpdateUsersViewByRoles(ctx, rolesIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateUsersViewByRoles", reflect.TypeOf((*MockIController)(nil).CreateOrUpdateUsersViewByRoles), ctx, rolesIds)
}

// GetUserView mocks base method
func (m *MockIController) GetUserView(ctx context.Context, id go_uuid.UUID) *models.UserView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserView", ctx, id)
	ret0, _ := ret[0].(*models.UserView)
	return ret0
}

// GetUserView indicates an expected call of GetUserView
func (mr *MockIControllerMockRecorder) GetUserView(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserView", reflect.TypeOf((*MockIController)(nil).GetUserView), ctx, id)
}

// GetUserViews mocks base method
func (m *MockIController) GetUserViews(ctx context.Context, query repositories.GetUsersViewStoreQuery) ([]*models.UserView, *models.PaginationResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserViews", ctx, query)
	ret0, _ := ret[0].([]*models.UserView)
	ret1, _ := ret[1].(*models.PaginationResponse)
	return ret0, ret1
}

// GetUserViews indicates an expected call of GetUserViews
func (mr *MockIControllerMockRecorder) GetUserViews(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserViews", reflect.TypeOf((*MockIController)(nil).GetUserViews), ctx, query)
}

// OnUserChanged mocks base method
func (m *MockIController) OnUserChanged(id []go_uuid.UUID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnUserChanged", id)
}

// OnUserChanged indicates an expected call of OnUserChanged
func (mr *MockIControllerMockRecorder) OnUserChanged(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnUserChanged", reflect.TypeOf((*MockIController)(nil).OnUserChanged), id)
}

// OnEmailCodeConfirmationCreated mocks base method
func (m *MockIController) OnEmailCodeConfirmationCreated(email, code string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEmailCodeConfirmationCreated", email, code)
}

// OnEmailCodeConfirmationCreated indicates an expected call of OnEmailCodeConfirmationCreated
func (mr *MockIControllerMockRecorder) OnEmailCodeConfirmationCreated(email, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEmailCodeConfirmationCreated", reflect.TypeOf((*MockIController)(nil).OnEmailCodeConfirmationCreated), email, code)
}

// OnPhoneCodeConfirmationCreated mocks base method
func (m *MockIController) OnPhoneCodeConfirmationCreated(phone, code string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPhoneCodeConfirmationCreated", phone, code)
}

// OnPhoneCodeConfirmationCreated indicates an expected call of OnPhoneCodeConfirmationCreated
func (mr *MockIControllerMockRecorder) OnPhoneCodeConfirmationCreated(phone, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPhoneCodeConfirmationCreated", reflect.TypeOf((*MockIController)(nil).OnPhoneCodeConfirmationCreated), phone, code)
}

// OnUsersViewChanged mocks base method
func (m *MockIController) OnUsersViewChanged(usersView []*models.UserView) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnUsersViewChanged", usersView)
}

// OnUsersViewChanged indicates an expected call of OnUsersViewChanged
func (mr *MockIControllerMockRecorder) OnUsersViewChanged(usersView interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnUsersViewChanged", reflect.TypeOf((*MockIController)(nil).OnUsersViewChanged), usersView)
}

// OnPasswordChanged mocks base method
func (m *MockIController) OnPasswordChanged(userId go_uuid.UUID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPasswordChanged", userId)
}

// OnPasswordChanged indicates an expected call of OnPasswordChanged
func (mr *MockIControllerMockRecorder) OnPasswordChanged(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPasswordChanged", reflect.TypeOf((*MockIController)(nil).OnPasswordChanged), userId)
}

// OnPhoneChanged mocks base method
func (m *MockIController) OnPhoneChanged(userId []go_uuid.UUID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPhoneChanged", userId)
}

// OnPhoneChanged indicates an expected call of OnPhoneChanged
func (mr *MockIControllerMockRecorder) OnPhoneChanged(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPhoneChanged", reflect.TypeOf((*MockIController)(nil).OnPhoneChanged), userId)
}

// OnEmailChanged mocks base method
func (m *MockIController) OnEmailChanged(userId []go_uuid.UUID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEmailChanged", userId)
}

// OnEmailChanged indicates an expected call of OnEmailChanged
func (mr *MockIControllerMockRecorder) OnEmailChanged(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEmailChanged", reflect.TypeOf((*MockIController)(nil).OnEmailChanged), userId)
}

// OnRoleChanged mocks base method
func (m *MockIController) OnRoleChanged(roleId []go_uuid.UUID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnRoleChanged", roleId)
}

// OnRoleChanged indicates an expected call of OnRoleChanged
func (mr *MockIControllerMockRecorder) OnRoleChanged(roleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnRoleChanged", reflect.TypeOf((*MockIController)(nil).OnRoleChanged), roleId)
}

// OnSecretCreatedV1 mocks base method
func (m *MockIController) OnSecretCreatedV1(secret *models.Secret) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnSecretCreatedV1", secret)
}

// OnSecretCreatedV1 indicates an expected call of OnSecretCreatedV1
func (mr *MockIControllerMockRecorder) OnSecretCreatedV1(secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnSecretCreatedV1", reflect.TypeOf((*MockIController)(nil).OnSecretCreatedV1), secret)
}
