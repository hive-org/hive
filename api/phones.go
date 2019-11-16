package api

import (
	"auth/app"
	"auth/controllers"
	"auth/enums"
	"auth/functools"
	"auth/infrastructure"
	"auth/inout"
	"auth/middlewares"
	"github.com/golang/protobuf/proto"
	"net/http"
)

func createPhoneV1(r *functools.Request, app infrastructure.AppInterface) (int, proto.Message) {
	body := inout.CreatePhoneRequestV1{}
	err := r.ParseBody(&body)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	status, phone := controllers.CreatePhone(app.GetStore(), app.GetESB(), r.Context(), body.Phone, body.Code, body.UserId)

	switch status {
	case enums.Ok:
		return http.StatusCreated, &inout.CreatePhoneResponseV1{
			Id:      phone.Id,
			Created: phone.Created,
			UserId:  phone.UserId,
			Phone:   phone.Value,
		}
	case enums.IncorrectPhoneCode:
		return http.StatusBadRequest, &inout.CreatePhoneBadRequestV1{
			Code: []string{"Неверный код"},
		}
	case enums.PhoneNotFound:
		return http.StatusBadRequest, &inout.CreatePhoneBadRequestV1{
			Phone: []string{"Не удалось найти код для данного телефона."},
		}
	case enums.UserNotFound:
		return http.StatusBadRequest, &inout.CreatePhoneBadRequestV1{
			UserId: []string{"Такого пользователя не существует"},
		}
	case enums.IncorrectPhone:
		return http.StatusBadRequest, &inout.CreatePhoneBadRequestV1{
			Phone: []string{"Некорректный номер телефона"},
		}
	default:
		return http.StatusCreated, &inout.CreatePhoneResponseV1{
			Id:      phone.Id,
			Created: phone.Created,
			UserId:  phone.UserId,
			Phone:   phone.Value,
		}
	}
}

func PhonesAPIV1(app *app.App) middlewares.ResponseControllerHandler {
	return func(r *functools.Request) (int, proto.Message) {
		return createPhoneV1(r, app)
	}
}

func createPhoneConfirmationV1(r *functools.Request, app infrastructure.AppInterface) (int, proto.Message) {

	body := inout.CreatePhoneConfirmationRequestV1{}

	err := r.ParseBody(&body)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	status, phoneConfirmation := controllers.CreatePhoneConfirmation(app.GetStore(), app.GetESB(), body.Phone)

	switch status {
	case enums.Ok:
		return http.StatusCreated, &inout.CreatePhoneConfirmationResponseV1{
			Created: phoneConfirmation.Created,
			Expire:  phoneConfirmation.Expire,
			Phone:   phoneConfirmation.Phone,
		}
	case enums.IncorrectPhone:
		return http.StatusBadRequest, &inout.CreatePhoneConfirmationBadRequestV1{
			Phone: []string{"Некорректный номер телефона"},
		}
	default:
		return http.StatusCreated, &inout.CreatePhoneConfirmationResponseV1{
			Created: phoneConfirmation.Created,
			Expire:  phoneConfirmation.Expire,
			Phone:   phoneConfirmation.Phone,
		}
	}
}

func PhoneConfirmationsAPIV1(app *app.App) middlewares.ResponseControllerHandler {
	return func(r *functools.Request) (int, proto.Message) {
		return createPhoneConfirmationV1(r, app)
	}
}