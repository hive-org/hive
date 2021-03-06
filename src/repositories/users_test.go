package repositories

import (
	"hive/config"
	"hive/enums"
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetRawQueryWithModifiedLimit(t *testing.T) {
	t.Parallel()

	repeatedUserID := uuid.NewV4()

	q := GetUsersQuery{
		Limit: 100,
		Id:    []uuid.UUID{uuid.NewV4(), uuid.NewV4(), uuid.NewV4(), uuid.NewV4(), repeatedUserID, repeatedUserID},
	}

	rw := convertGetUsersQueryToRaw(q)

	require.True(t, rw.Limit == 6)
}

func TestGetRawQueryLimitWithEmptyId(t *testing.T) {
	t.Parallel()

	q := GetUsersQuery{
		Limit: 100,
		Id:    []uuid.UUID{},
	}

	rw := convertGetUsersQueryToRaw(q)

	require.True(t, rw.Limit == 100)
}

func TestGetRawQueryWithLimitLessThenId(t *testing.T) {
	t.Parallel()

	repeatedUserID := uuid.NewV4()

	q := GetUsersQuery{
		Limit: 3,
		Id:    []uuid.UUID{uuid.NewV4(), uuid.NewV4(), uuid.NewV4(), uuid.NewV4(), repeatedUserID, repeatedUserID},
	}

	rw := convertGetUsersQueryToRaw(q)

	require.True(t, rw.Limit == 3)
}

func TestGetRawQueryWithEmptyId(t *testing.T) {
	t.Parallel()

	q := GetUsersQuery{
		Id: []uuid.UUID{},
	}

	rw := convertGetUsersQueryToRaw(q)

	require.True(t, rw.Id == "{}")
}

func TestGetRawQueryWithId(t *testing.T) {
	t.Parallel()

	first, second, third := uuid.NewV4(), uuid.NewV4(), uuid.NewV4()

	q := GetUsersQuery{
		Id: []uuid.UUID{first, second, third},
	}

	rw := convertGetUsersQueryToRaw(q)

	require.True(t, rw.Id == fmt.Sprintf("{%s,%s,%s}", first.String(), second.String(), third.String()))
}

func TestGetUser(t *testing.T) {
	pool := config.InitPool(nil, config.InitEnvironment())
	ctx := context.Background()
	PurgeUsers(pool, ctx)
	user := CreateUser(pool, ctx)
	userFromDB := GetUser(pool, ctx, user.Id)
	require.NotNil(t, userFromDB)
	require.Equal(t, user, userFromDB)
}


func TestGetUserWithoutUser(t *testing.T) {
	pool := config.InitPool(nil, config.InitEnvironment())
	ctx := context.Background()
	PurgeUsers(pool, ctx)
	userFromDB := GetUser(pool, ctx, uuid.NewV4())
	require.Nil(t, userFromDB)
}

func TestDeleteUser(t *testing.T) {
	pool := config.InitPool(nil, config.InitEnvironment())
	ctx := context.Background()
	PurgeUsers(pool, ctx)
	user := CreateUser(pool, ctx)
	status, deletedUser := DeleteUser(pool, ctx, user.Id)
	require.Equal(t, enums.Ok, status)
	require.Equal(t, user, deletedUser)
}

func TestDeleteUserWithoutUser(t *testing.T) {
	pool := config.InitPool(nil, config.InitEnvironment())
	ctx := context.Background()
	PurgeUsers(pool, ctx)
	status, deletedUser := DeleteUser(pool, ctx, uuid.NewV4())
	require.Equal(t, enums.UserNotFound, status)
	require.Nil(t, deletedUser)
}


func BenchmarkCreateUser(b *testing.B) {
	pool := config.InitPool(nil, config.InitEnvironment())
	ctx := context.Background()
	PurgeUsers(pool, ctx)
	tx, _ := pool.Begin(ctx)

	for i := 1; i < 1_000; i++ {
		CreateUser(tx, ctx)
	}

	_ = tx.Commit(ctx)
}
