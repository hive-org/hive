package api

import (
	"auth/app"
	"auth/config"
	"auth/controllers"
	"auth/enums"
	"auth/functools"
	"auth/infrastructure"
	"auth/inout"
	"auth/middlewares"
	"github.com/getsentry/sentry-go"
	"github.com/golang/protobuf/proto"
	"net/http"
	"time"
)

func createSessionV1(r *functools.Request, app infrastructure.AppInterface) (int, proto.Message) {

	body := inout.CreateSessionRequestV1{}
	err := r.ParseBody(&body)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	tokens := body.GetTokens()

	if tokens != nil {
		cookie, err := r.Cookie(enums.RefreshToken)
		if err != nil {
			sentry.CaptureException(err)
		} else {
			tokens.RefreshToken = cookie.Value
			body.Data = &inout.CreateSessionRequestV1_Tokens_{Tokens: tokens}
		}
	}

	status, session := controllers.CreateSession(app.GetStore(), app.GetLoginController(), r.Context(), body)

	switch status {
	case enums.Ok:
		env := config.GetEnvironment()
		http.SetCookie(r.Response, &http.Cookie{
			Name:     enums.RefreshToken,
			Value:    session.RefreshToken,
			Domain:   r.Referer(),
			Expires:  time.Now().Add(time.Hour * 24 * time.Duration(env.RefreshTokenLifetime)),
			Secure:   true,
			HttpOnly: true,
		})
		return http.StatusCreated, &inout.CreateSessionResponseV1{
			RefreshToken: session.RefreshToken,
			AccessToken:  session.AccessToken,
		}
	case
		enums.SessionNotFound,
		enums.UserNotFound,
		enums.IncorrectToken,
		enums.InvalidToken,
		enums.SecretNotFound:
		return http.StatusUnauthorized, nil
	case enums.IncorrectPassword:
		return http.StatusBadRequest, &inout.CreateSessionBadRequestResponseV1{
			Password: []string{"Некорректный пароль"},
		}
	case enums.PasswordNotFound:
		return http.StatusBadRequest, &inout.CreateSessionBadRequestResponseV1{
			Password: []string{"Не установлен пароль"},
		}
	case enums.EmailNotFound:
		return http.StatusBadRequest, &inout.CreateSessionBadRequestResponseV1{
			Email: []string{"Email не найден"},
		}
	case enums.IncorrectEmail:
		return http.StatusBadRequest, &inout.CreateSessionBadRequestResponseV1{
			Email: []string{"Некорректный email"},
		}
	case enums.IncorrectEmailCode:
		return http.StatusBadRequest, &inout.CreateSessionBadRequestResponseV1{
			EmailCode: []string{"Некорректный код подтверждения"},
		}
	case enums.EmailConfirmationCodeNotFound:
		return http.StatusBadRequest, &inout.CreateSessionBadRequestResponseV1{
			EmailCode: []string{"Не найден код подтверждения для данного email"},
		}
	case enums.PhoneNotFound:
		return http.StatusBadRequest, &inout.CreateSessionBadRequestResponseV1{
			Phone: []string{"Телефон не найден"},
		}
	case enums.IncorrectPhone:
		return http.StatusBadRequest, &inout.CreateSessionBadRequestResponseV1{
			Phone: []string{"Некорректный телефон"},
		}
	case enums.IncorrectPhoneCode:
		return http.StatusBadRequest, &inout.CreateSessionBadRequestResponseV1{
			PhoneCode: []string{"Некорректный код подтверждения"},
		}
	case enums.PhoneConfirmationCodeNotFound:
		return http.StatusBadRequest, &inout.CreateSessionBadRequestResponseV1{
			PhoneCode: []string{"Не найден код подтверждения для данного телефона"},
		}
	default:
		return http.StatusInternalServerError, nil
	}
}

func SessionsAPIV1(app *app.App) middlewares.ResponseControllerHandler {
	return func(r *functools.Request) (int, proto.Message) {
		return createSessionV1(r, app)
	}
}
