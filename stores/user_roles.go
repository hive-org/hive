package stores

import (
	"auth/models"
	"auth/repositories"
	"context"
	uuid "github.com/satori/go.uuid"
)

func (store *DatabaseStore) CreateUserRole(ctx context.Context, userId uuid.UUID, roleId uuid.UUID) (int, *models.UserRole) {
	return repositories.CreateUserRole(store.Db, ctx, userId, roleId)
}

func (store *DatabaseStore) GetUserRoles(ctx context.Context, query repositories.GetUserRoleQuery) ([]*models.UserRole, *models.PaginationResponse) {
	return repositories.GetUserRoles(store.Db, ctx, query)
}

func (store *DatabaseStore) DeleteUserRole(ctx context.Context, id uuid.UUID) (int, *models.UserRole) {
	return repositories.DeleteUserRole(store.Db, ctx, id)
}
