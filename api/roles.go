package api

import (
	"auth/app"
	"auth/controllers"
	"auth/functools"
	"auth/infrastructure"
	"auth/inout"
	"auth/middlewares"
	"auth/repositories"
	"github.com/getsentry/sentry-go"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func GetRolesV1Query(r *http.Request) repositories.GetRolesQuery {
	limitQuery := r.URL.Query().Get("limit")
	if limitQuery == "" {
		limitQuery = string(DefaultLimit)
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		limit = DefaultLimit
	}

	identifiersQuery := r.URL.Query().Get("id")
	var identifiers []int64

	if identifiersQuery != "" {
		identifiersQueryStr := strings.SplitN(identifiersQuery, ",", -1)
		identifiersQueryInt := make([]int64, len(identifiersQueryStr))
		for i, q := range identifiersQueryStr {
			idQueryInt, _ := strconv.Atoi(q)
			identifiersQueryInt[i] = int64(idQueryInt)
		}

		identifiers = identifiersQueryInt
	}

	return repositories.GetRolesQuery{Limit: limit, Identifiers: identifiers}
}

func getRoleV1(r *functools.Request, app *app.App, id int64) (int, *inout.GetRoleResponseV1) {
	role := app.Store.GetRole(r.Context(), id)

	if role == nil {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, &inout.GetRoleResponseV1{
		Id:      role.Id,
		Created: role.Created,
		Title:   role.Title,
	}
}

func getRolesV1(r *functools.Request, app *app.App) (int, *inout.ListRoleResponseV1) {

	query := GetRolesV1Query(r.Request)
	roles := app.Store.GetRoles(r.Context(), query)
	rolesData := make([]*inout.GetRoleResponseV1, len(roles))

	for i, role := range roles {
		rolesData[i] = &inout.GetRoleResponseV1{
			Id:      role.Id,
			Created: role.Created,
			Title:   role.Title,
		}
	}

	return http.StatusOK, &inout.ListRoleResponseV1{Data: rolesData}
}

func createRoleV1(r *functools.Request, app infrastructure.AppInterface) (int, *inout.GetRoleResponseV1) {

	body := inout.CreateRoleRequestV1{}

	err := r.ParseBody(&body)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	role := controllers.CreateRole(app.GetStore(), app.GetESB(), r.Context(), body.Title)

	if role == nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusCreated, &inout.GetRoleResponseV1{
		Id:      role.Id,
		Created: role.Created,
		Title:   role.Title,
	}
}

func RoleAPIV1(app *app.App) middlewares.ResponseControllerHandler {
	return func(request *functools.Request) (i int, message proto.Message) {

		vars := mux.Vars(request.Request)
		id, err := strconv.ParseInt(vars["id"], 10, 64)

		if err != nil {
			// Сознательно отправляем отчет об ошибке, т.к. в vars["id"] не должны попасть не числовые значения.
			// Если такое произошло - что то пошло не так
			sentry.CaptureException(err)
			return http.StatusBadRequest, nil
		}

		return getRoleV1(request, app, id)
	}
}

func RolesAPIV1(app *app.App) middlewares.ResponseControllerHandler {
	return func(r *functools.Request) (int, proto.Message) {
		if r.Method == "GET" {
			return getRolesV1(r, app)
		} else {
			return createRoleV1(r, app)
		}
	}
}