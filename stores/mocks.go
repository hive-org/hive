// Code generated by MockGen. DO NOT EDIT.
// Source: main.go

// Package stores is a generated GoMock package.
package stores

import (
	models "hive/models"
	repositories "hive/repositories"
	context "context"
	gomock "github.com/golang/mock/gomock"
	go_uuid "github.com/satori/go.uuid"
	reflect "reflect"
	time "time"
)

// MockIStore is a mock of IStore interface
type MockIStore struct {
	ctrl     *gomock.Controller
	recorder *MockIStoreMockRecorder
}

// MockIStoreMockRecorder is the mock recorder for MockIStore
type MockIStoreMockRecorder struct {
	mock *MockIStore
}

// NewMockIStore creates a new mock instance
func NewMockIStore(ctrl *gomock.Controller) *MockIStore {
	mock := &MockIStore{ctrl: ctrl}
	mock.recorder = &MockIStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIStore) EXPECT() *MockIStoreMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockIStore) CreateUser(ctx context.Context, password, email, phone string) (int, *models.User) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, password, email, phone)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.User)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockIStoreMockRecorder) CreateUser(ctx, password, email, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIStore)(nil).CreateUser), ctx, password, email, phone)
}

// GetUser mocks base method
func (m *MockIStore) GetUser(context context.Context, id go_uuid.UUID) *models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", context, id)
	ret0, _ := ret[0].(*models.User)
	return ret0
}

// GetUser indicates an expected call of GetUser
func (mr *MockIStoreMockRecorder) GetUser(context, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIStore)(nil).GetUser), context, id)
}

// GetUsers mocks base method
func (m *MockIStore) GetUsers(context context.Context, query repositories.GetUsersQuery) []*models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", context, query)
	ret0, _ := ret[0].([]*models.User)
	return ret0
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockIStoreMockRecorder) GetUsers(context, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockIStore)(nil).GetUsers), context, query)
}

// DeleteUser mocks base method
func (m *MockIStore) DeleteUser(ctx context.Context, id go_uuid.UUID) (int, *models.User) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.User)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockIStoreMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIStore)(nil).DeleteUser), ctx, id)
}

// GetUsersView mocks base method
func (m *MockIStore) GetUsersView(context context.Context, query repositories.GetUsersViewStoreQuery) ([]*models.UserView, *models.PaginationResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersView", context, query)
	ret0, _ := ret[0].([]*models.UserView)
	ret1, _ := ret[1].(*models.PaginationResponse)
	return ret0, ret1
}

// GetUsersView indicates an expected call of GetUsersView
func (mr *MockIStoreMockRecorder) GetUsersView(context, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersView", reflect.TypeOf((*MockIStore)(nil).GetUsersView), context, query)
}

// GetUserView mocks base method
func (m *MockIStore) GetUserView(context context.Context, id go_uuid.UUID) *models.UserView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserView", context, id)
	ret0, _ := ret[0].(*models.UserView)
	return ret0
}

// GetUserView indicates an expected call of GetUserView
func (mr *MockIStoreMockRecorder) GetUserView(context, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserView", reflect.TypeOf((*MockIStore)(nil).GetUserView), context, id)
}

// CreateOrUpdateUsersView mocks base method
func (m *MockIStore) CreateOrUpdateUsersView(context context.Context, query repositories.CreateOrUpdateUsersViewStoreQuery) []*models.UserView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateUsersView", context, query)
	ret0, _ := ret[0].([]*models.UserView)
	return ret0
}

// CreateOrUpdateUsersView indicates an expected call of CreateOrUpdateUsersView
func (mr *MockIStoreMockRecorder) CreateOrUpdateUsersView(context, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateUsersView", reflect.TypeOf((*MockIStore)(nil).CreateOrUpdateUsersView), context, query)
}

// CreateOrUpdateUsersViewByUsersID mocks base method
func (m *MockIStore) CreateOrUpdateUsersViewByUsersID(context context.Context, id []go_uuid.UUID) []*models.UserView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateUsersViewByUsersID", context, id)
	ret0, _ := ret[0].([]*models.UserView)
	return ret0
}

// CreateOrUpdateUsersViewByUsersID indicates an expected call of CreateOrUpdateUsersViewByUsersID
func (mr *MockIStoreMockRecorder) CreateOrUpdateUsersViewByUsersID(context, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateUsersViewByUsersID", reflect.TypeOf((*MockIStore)(nil).CreateOrUpdateUsersViewByUsersID), context, id)
}

// CreateOrUpdateUsersViewByRolesID mocks base method
func (m *MockIStore) CreateOrUpdateUsersViewByRolesID(context context.Context, id []go_uuid.UUID) []*models.UserView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateUsersViewByRolesID", context, id)
	ret0, _ := ret[0].([]*models.UserView)
	return ret0
}

// CreateOrUpdateUsersViewByRolesID indicates an expected call of CreateOrUpdateUsersViewByRolesID
func (mr *MockIStoreMockRecorder) CreateOrUpdateUsersViewByRolesID(context, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateUsersViewByRolesID", reflect.TypeOf((*MockIStore)(nil).CreateOrUpdateUsersViewByRolesID), context, id)
}

// CreateOrUpdateUsersViewByUserID mocks base method
func (m *MockIStore) CreateOrUpdateUsersViewByUserID(context context.Context, id go_uuid.UUID) []*models.UserView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateUsersViewByUserID", context, id)
	ret0, _ := ret[0].([]*models.UserView)
	return ret0
}

// CreateOrUpdateUsersViewByUserID indicates an expected call of CreateOrUpdateUsersViewByUserID
func (mr *MockIStoreMockRecorder) CreateOrUpdateUsersViewByUserID(context, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateUsersViewByUserID", reflect.TypeOf((*MockIStore)(nil).CreateOrUpdateUsersViewByUserID), context, id)
}

// CreateOrUpdateUsersViewByRoleID mocks base method
func (m *MockIStore) CreateOrUpdateUsersViewByRoleID(context context.Context, id go_uuid.UUID) []*models.UserView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateUsersViewByRoleID", context, id)
	ret0, _ := ret[0].([]*models.UserView)
	return ret0
}

// CreateOrUpdateUsersViewByRoleID indicates an expected call of CreateOrUpdateUsersViewByRoleID
func (mr *MockIStoreMockRecorder) CreateOrUpdateUsersViewByRoleID(context, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateUsersViewByRoleID", reflect.TypeOf((*MockIStore)(nil).CreateOrUpdateUsersViewByRoleID), context, id)
}

// CacheUserView mocks base method
func (m *MockIStore) CacheUserView(ctx context.Context, userViews []*models.UserView) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CacheUserView", ctx, userViews)
}

// CacheUserView indicates an expected call of CacheUserView
func (mr *MockIStoreMockRecorder) CacheUserView(ctx, userViews interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheUserView", reflect.TypeOf((*MockIStore)(nil).CacheUserView), ctx, userViews)
}

// CreateEmail mocks base method
func (m *MockIStore) CreateEmail(ctx context.Context, userId go_uuid.UUID, value string) (int, *models.Email) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmail", ctx, userId, value)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Email)
	return ret0, ret1
}

// CreateEmail indicates an expected call of CreateEmail
func (mr *MockIStoreMockRecorder) CreateEmail(ctx, userId, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmail", reflect.TypeOf((*MockIStore)(nil).CreateEmail), ctx, userId, value)
}

// GetEmail mocks base method
func (m *MockIStore) GetEmail(ctx context.Context, email string) (int, *models.Email) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmail", ctx, email)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Email)
	return ret0, ret1
}

// GetEmail indicates an expected call of GetEmail
func (mr *MockIStoreMockRecorder) GetEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmail", reflect.TypeOf((*MockIStore)(nil).GetEmail), ctx, email)
}

// CreateEmailConfirmationCode mocks base method
func (m *MockIStore) CreateEmailConfirmationCode(ctx context.Context, email, code string, duration time.Duration) *models.EmailConfirmation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmailConfirmationCode", ctx, email, code, duration)
	ret0, _ := ret[0].(*models.EmailConfirmation)
	return ret0
}

// CreateEmailConfirmationCode indicates an expected call of CreateEmailConfirmationCode
func (mr *MockIStoreMockRecorder) CreateEmailConfirmationCode(ctx, email, code, duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmailConfirmationCode", reflect.TypeOf((*MockIStore)(nil).CreateEmailConfirmationCode), ctx, email, code, duration)
}

// GetEmailConfirmationCode mocks base method
func (m *MockIStore) GetEmailConfirmationCode(ctx context.Context, email string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmailConfirmationCode", ctx, email)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEmailConfirmationCode indicates an expected call of GetEmailConfirmationCode
func (mr *MockIStoreMockRecorder) GetEmailConfirmationCode(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailConfirmationCode", reflect.TypeOf((*MockIStore)(nil).GetEmailConfirmationCode), ctx, email)
}

// GetRandomCodeForEmailConfirmation mocks base method
func (m *MockIStore) GetRandomCodeForEmailConfirmation() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomCodeForEmailConfirmation")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRandomCodeForEmailConfirmation indicates an expected call of GetRandomCodeForEmailConfirmation
func (mr *MockIStoreMockRecorder) GetRandomCodeForEmailConfirmation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomCodeForEmailConfirmation", reflect.TypeOf((*MockIStore)(nil).GetRandomCodeForEmailConfirmation))
}

// CreatePassword mocks base method
func (m *MockIStore) CreatePassword(ctx context.Context, userId go_uuid.UUID, value string) (int, *models.Password) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePassword", ctx, userId, value)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Password)
	return ret0, ret1
}

// CreatePassword indicates an expected call of CreatePassword
func (mr *MockIStoreMockRecorder) CreatePassword(ctx, userId, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePassword", reflect.TypeOf((*MockIStore)(nil).CreatePassword), ctx, userId, value)
}

// GetPasswords mocks base method
func (m *MockIStore) GetPasswords(ctx context.Context, userId go_uuid.UUID) []*models.Password {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPasswords", ctx, userId)
	ret0, _ := ret[0].([]*models.Password)
	return ret0
}

// GetPasswords indicates an expected call of GetPasswords
func (mr *MockIStoreMockRecorder) GetPasswords(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPasswords", reflect.TypeOf((*MockIStore)(nil).GetPasswords), ctx, userId)
}

// GetLatestPassword mocks base method
func (m *MockIStore) GetLatestPassword(ctx context.Context, userId go_uuid.UUID) (int, *models.Password) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestPassword", ctx, userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Password)
	return ret0, ret1
}

// GetLatestPassword indicates an expected call of GetLatestPassword
func (mr *MockIStoreMockRecorder) GetLatestPassword(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestPassword", reflect.TypeOf((*MockIStore)(nil).GetLatestPassword), ctx, userId)
}

// CreatePhone mocks base method
func (m *MockIStore) CreatePhone(ctx context.Context, userId go_uuid.UUID, value string) (int, *models.Phone) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePhone", ctx, userId, value)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Phone)
	return ret0, ret1
}

// CreatePhone indicates an expected call of CreatePhone
func (mr *MockIStoreMockRecorder) CreatePhone(ctx, userId, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePhone", reflect.TypeOf((*MockIStore)(nil).CreatePhone), ctx, userId, value)
}

// GetPhone mocks base method
func (m *MockIStore) GetPhone(ctx context.Context, phone string) (int, *models.Phone) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhone", ctx, phone)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Phone)
	return ret0, ret1
}

// GetPhone indicates an expected call of GetPhone
func (mr *MockIStoreMockRecorder) GetPhone(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhone", reflect.TypeOf((*MockIStore)(nil).GetPhone), ctx, phone)
}

// CreatePhoneConfirmationCode mocks base method
func (m *MockIStore) CreatePhoneConfirmationCode(ctx context.Context, phone, code string, duration time.Duration) *models.PhoneConfirmation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePhoneConfirmationCode", ctx, phone, code, duration)
	ret0, _ := ret[0].(*models.PhoneConfirmation)
	return ret0
}

// CreatePhoneConfirmationCode indicates an expected call of CreatePhoneConfirmationCode
func (mr *MockIStoreMockRecorder) CreatePhoneConfirmationCode(ctx, phone, code, duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePhoneConfirmationCode", reflect.TypeOf((*MockIStore)(nil).CreatePhoneConfirmationCode), ctx, phone, code, duration)
}

// GetPhoneConfirmationCode mocks base method
func (m *MockIStore) GetPhoneConfirmationCode(ctx context.Context, phone string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhoneConfirmationCode", ctx, phone)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPhoneConfirmationCode indicates an expected call of GetPhoneConfirmationCode
func (mr *MockIStoreMockRecorder) GetPhoneConfirmationCode(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhoneConfirmationCode", reflect.TypeOf((*MockIStore)(nil).GetPhoneConfirmationCode), ctx, phone)
}

// GetRandomCodeForPhoneConfirmation mocks base method
func (m *MockIStore) GetRandomCodeForPhoneConfirmation() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomCodeForPhoneConfirmation")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRandomCodeForPhoneConfirmation indicates an expected call of GetRandomCodeForPhoneConfirmation
func (mr *MockIStoreMockRecorder) GetRandomCodeForPhoneConfirmation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomCodeForPhoneConfirmation", reflect.TypeOf((*MockIStore)(nil).GetRandomCodeForPhoneConfirmation))
}

// CreateRole mocks base method
func (m *MockIStore) CreateRole(context context.Context, title string) (int, *models.Role) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", context, title)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Role)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole
func (mr *MockIStoreMockRecorder) CreateRole(context, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockIStore)(nil).CreateRole), context, title)
}

// GetRole mocks base method
func (m *MockIStore) GetRole(context context.Context, id go_uuid.UUID) (int, *models.Role) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", context, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Role)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole
func (mr *MockIStoreMockRecorder) GetRole(context, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockIStore)(nil).GetRole), context, id)
}

// GetRoles mocks base method
func (m *MockIStore) GetRoles(context context.Context, query repositories.GetRolesQuery) ([]*models.Role, *models.PaginationResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoles", context, query)
	ret0, _ := ret[0].([]*models.Role)
	ret1, _ := ret[1].(*models.PaginationResponse)
	return ret0, ret1
}

// GetRoles indicates an expected call of GetRoles
func (mr *MockIStoreMockRecorder) GetRoles(context, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoles", reflect.TypeOf((*MockIStore)(nil).GetRoles), context, query)
}

// GetRoleByTitle mocks base method
func (m *MockIStore) GetRoleByTitle(ctx context.Context, title string) (int, *models.Role) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleByTitle", ctx, title)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Role)
	return ret0, ret1
}

// GetRoleByTitle indicates an expected call of GetRoleByTitle
func (mr *MockIStoreMockRecorder) GetRoleByTitle(ctx, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleByTitle", reflect.TypeOf((*MockIStore)(nil).GetRoleByTitle), ctx, title)
}

// GetAdminRole mocks base method
func (m *MockIStore) GetAdminRole(ctx context.Context) (int, *models.Role) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdminRole", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.Role)
	return ret0, ret1
}

// GetAdminRole indicates an expected call of GetAdminRole
func (mr *MockIStoreMockRecorder) GetAdminRole(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdminRole", reflect.TypeOf((*MockIStore)(nil).GetAdminRole), ctx)
}

// CreateUserRole mocks base method
func (m *MockIStore) CreateUserRole(ctx context.Context, userId, roleId go_uuid.UUID) (int, *models.UserRole) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserRole", ctx, userId, roleId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.UserRole)
	return ret0, ret1
}

// CreateUserRole indicates an expected call of CreateUserRole
func (mr *MockIStoreMockRecorder) CreateUserRole(ctx, userId, roleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserRole", reflect.TypeOf((*MockIStore)(nil).CreateUserRole), ctx, userId, roleId)
}

// GetUserRoles mocks base method
func (m *MockIStore) GetUserRoles(ctx context.Context, query repositories.GetUserRoleQuery) ([]*models.UserRole, *models.PaginationResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRoles", ctx, query)
	ret0, _ := ret[0].([]*models.UserRole)
	ret1, _ := ret[1].(*models.PaginationResponse)
	return ret0, ret1
}

// GetUserRoles indicates an expected call of GetUserRoles
func (mr *MockIStoreMockRecorder) GetUserRoles(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRoles", reflect.TypeOf((*MockIStore)(nil).GetUserRoles), ctx, query)
}

// DeleteUserRole mocks base method
func (m *MockIStore) DeleteUserRole(ctx context.Context, id go_uuid.UUID) (int, *models.UserRole) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserRole", ctx, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*models.UserRole)
	return ret0, ret1
}

// DeleteUserRole indicates an expected call of DeleteUserRole
func (mr *MockIStoreMockRecorder) DeleteUserRole(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserRole", reflect.TypeOf((*MockIStore)(nil).DeleteUserRole), ctx, id)
}

// GetSecret mocks base method
func (m *MockIStore) GetSecret(ctx context.Context, id go_uuid.UUID) *models.Secret {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", ctx, id)
	ret0, _ := ret[0].(*models.Secret)
	return ret0
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockIStoreMockRecorder) GetSecret(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockIStore)(nil).GetSecret), ctx, id)
}

// GetActualSecret mocks base method
func (m *MockIStore) GetActualSecret(ctx context.Context) *models.Secret {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActualSecret", ctx)
	ret0, _ := ret[0].(*models.Secret)
	return ret0
}

// GetActualSecret indicates an expected call of GetActualSecret
func (mr *MockIStoreMockRecorder) GetActualSecret(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActualSecret", reflect.TypeOf((*MockIStore)(nil).GetActualSecret), ctx)
}

// CreateSecret mocks base method
func (m *MockIStore) CreateSecret(ctx context.Context) *models.Secret {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", ctx)
	ret0, _ := ret[0].(*models.Secret)
	return ret0
}

// CreateSecret indicates an expected call of CreateSecret
func (mr *MockIStoreMockRecorder) CreateSecret(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockIStore)(nil).CreateSecret), ctx)
}

// CreateSession mocks base method
func (m *MockIStore) CreateSession(ctx context.Context, userID, secretID go_uuid.UUID, fingerprint, userAgent string) *models.Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, userID, secretID, fingerprint, userAgent)
	ret0, _ := ret[0].(*models.Session)
	return ret0
}

// CreateSession indicates an expected call of CreateSession
func (mr *MockIStoreMockRecorder) CreateSession(ctx, userID, secretID, fingerprint, userAgent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockIStore)(nil).CreateSession), ctx, userID, secretID, fingerprint, userAgent)
}

// DeleteSession mocks base method
func (m *MockIStore) DeleteSession(ctx context.Context, id go_uuid.UUID) *models.Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", ctx, id)
	ret0, _ := ret[0].(*models.Session)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession
func (mr *MockIStoreMockRecorder) DeleteSession(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockIStore)(nil).DeleteSession), ctx, id)
}
