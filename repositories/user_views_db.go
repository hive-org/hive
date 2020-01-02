package repositories

import (
	"auth/inout"
	"auth/models"
	"auth/modelsFunctools"
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/jackc/pgx/v4"
)

type GetUsersViewQuery struct {
	GetUsersQuery
	Roles []models.RoleID
}

type CreateOrUpdateUsersViewQuery = GetUsersViewQuery

type getUsersViewRawQuery struct {
	getUsersRawQuery
	Roles string
}

type createOrUpdateUsersViewRawQuery = getUsersViewRawQuery

// SQL

func getUsersViewSQL() string {
	return `
		SELECT id, created, roles, phones, emails
		FROM users_view u
		WHERE (array_length($1::integer[], 1) IS NULL OR id = ANY ($1::bigint[])) AND 
		      (array_length($2::integer[], 1) IS NULL OR ($2::bigint[]) && role_id)
		LIMIT $3;
		`
}

func updateUsersViewSQL() string {
	return `
		INSERT
		INTO users_view(id, created, roles, phones, emails, role_id)
		SELECT nuv.id, nuv.created, nuv.roles, nuv.phones, nuv.emails, nuv.role_id
		FROM users_view as cuv
				 FULL OUTER JOIN (SELECT u.id,
										 u.created,
										 array_remove(array_agg(DISTINCT r.title), NULL)::text[]          as roles,
										 array_remove(array_agg(p.value), NULL)::text[]                   as phones,
										 array_remove(array_agg(DISTINCT e.value), NULL)::text[]          as emails,
										 array_remove(array_agg(DISTINCT r.id), NULL)                     as role_id
								  FROM users u
										   LEFT JOIN emails e on u.id = e.user_id
										   LEFT JOIN phones p on u.id = p.user_id
										   LEFT JOIN user_roles ur on u.id = ur.user_id
										   LEFT JOIN roles r on ur.role_id = r.id
								  WHERE (array_length($1::bigint[], 1) IS NULL OR u.id = ANY ($1::bigint[]))
									AND (array_length($2::bigint[], 1) IS NULL OR r.id = ANY ($2::bigint[]))
								  GROUP BY u.id, ur.created, r.created, p.created, e.created) as nuv
								 ON nuv.id = cuv.id AND
									nuv.created = cuv.created AND
									nuv.phones = cuv.phones AND
									nuv.roles = cuv.roles AND
									nuv.emails = cuv.emails AND
									nuv.role_id = cuv.role_id
		WHERE cuv.id IS NULL
		ORDER BY id
		ON CONFLICT (id) DO UPDATE SET created=excluded.created,
									   phones=excluded.phones,
									   roles=excluded.roles,
									   emails=excluded.emails,
									   role_id=excluded.role_id
		RETURNING id, created, roles, phones, emails;
    `
}

func scanUserView(row pgx.Row) *inout.GetUserViewResponseV1 {
	userView := &inout.GetUserViewResponseV1{}

	err := row.Scan(&userView.Id, &userView.Created, &userView.Roles, &userView.Phones, &userView.Emails)
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}

	return userView
}

func scanUsersView(rows pgx.Rows, limit int) []*inout.GetUserViewResponseV1 {

	users := make([]*inout.GetUserViewResponseV1, limit)

	var i int

	for rows.Next() {

		user := scanUserView(rows)

		if len(users) <= i {
			users = append(users, user)
		} else {
			users[i] = user
		}

		i++
	}

	rows.Close()
	return users[:i]
}

func convertGetUsersViewQueryToRaw(query GetUsersViewQuery) getUsersViewRawQuery {

	userRawQuery := convertGetUsersQueryToRaw(query.GetUsersQuery)

	return getUsersViewRawQuery{
		getUsersRawQuery: userRawQuery,
		Roles:            modelsFunctools.RoleIDListToPGArray(query.Roles),
	}
}

func convertCreateOrUpdateUsersViewQueryToRaw(query CreateOrUpdateUsersViewQuery) createOrUpdateUsersViewRawQuery {

	userRawQuery := convertGetUsersQueryToRaw(query.GetUsersQuery)

	return createOrUpdateUsersViewRawQuery{
		getUsersRawQuery: userRawQuery,
		Roles:            modelsFunctools.RoleIDListToPGArray(query.Roles),
	}
}

func GetUsersView(db DB, context context.Context, query GetUsersViewQuery) []*inout.GetUserViewResponseV1 {
	sql := getUsersViewSQL()
	rawQuery := convertGetUsersViewQueryToRaw(query)

	rows, err := db.Query(context, sql, rawQuery.Id, rawQuery.Roles, rawQuery.Limit)
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}

	return scanUsersView(rows, rawQuery.Limit)
}

func GetUserView(db DB, context context.Context, id models.UserID) *inout.GetUserViewResponseV1 {
	sql := getUsersViewSQL()
	row := db.QueryRow(context, sql, modelsFunctools.UserIDListToPGArray([]models.UserID{id}), "{}", 1)
	userView := scanUserView(row)
	return userView
}

func CreateOrUpdateUsersView(db DB, context context.Context, query CreateOrUpdateUsersViewQuery) []*inout.GetUserViewResponseV1 {
	sql := updateUsersViewSQL()
	rawQuery := convertCreateOrUpdateUsersViewQueryToRaw(query)
	rows, err := db.Query(context, sql, rawQuery.Id, rawQuery.Roles)
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}

	return scanUsersView(rows, len(query.Id))
}
