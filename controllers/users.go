package controllers

import (
	"auth/enums"
	"auth/functools"
	"auth/inout"
	"auth/models"
	"auth/repositories"
	"context"
	uuid "github.com/satori/go.uuid"
)

func (controller *Controller) CreateUser(ctx context.Context, body *inout.CreateUserResponseV1_Request) (int, *models.User) {

	var identifiers []uuid.UUID

	if body.Password == "" {
		return enums.PasswordRequired, nil
	}

	if body.Email == "" && body.Phone == "" {
		return enums.MinimumOneFieldRequired, nil
	}

	body.Password = controller.passwordProcessor.EncodePassword(ctx, body.Password)
	if body.Password == "" {
		return enums.IncorrectPassword, nil
	}

	if len(body.Phone) > 0 {
		phone := functools.NormalizePhone(body.Phone)
		if phone == "" {
			return enums.IncorrectPhone, nil
		}

		cachedCode := controller.store.GetPhoneConfirmationCode(ctx, phone)
		if cachedCode == "" {
			return enums.PhoneNotFound, nil
		} else if cachedCode != body.PhoneCode {
			return enums.IncorrectPhoneCode, nil
		}

		body.Phone = phone
		_, oldPhone := controller.store.GetPhone(ctx, phone)
		if oldPhone != nil {
			identifiers = append(identifiers, oldPhone.UserId)
		}
	}

	if len(body.Email) > 0 {
		emailStatus, email := controller.validateEmail(ctx, body.Email, body.EmailCode)
		if emailStatus != enums.Ok {
			return emailStatus, nil
		}

		body.Email = email
		_, oldEmail := controller.store.GetEmail(ctx, email)
		if oldEmail != nil {
			identifiers = append(identifiers, oldEmail.UserId)
		}
	}

	status, user := controller.store.CreateUser(ctx, body)
	identifiers = append(identifiers, user.Id)
	controller.OnUserChanged(identifiers)
	return status, user
}

func (controller *Controller) DeleteUser(ctx context.Context, id uuid.UUID) (int, *models.User) {
	status, deletedUser := controller.store.DeleteUser(ctx, id)
	if status == enums.Ok {
		controller.OnUserChanged([]uuid.UUID{deletedUser.Id})
	}

	return status, deletedUser
}

func (controller *Controller) GetUsers(ctx context.Context, query repositories.GetUsersQuery) []*models.User {
	return controller.store.GetUsers(ctx, query)
}

func (controller *Controller) GetUser(ctx context.Context, id uuid.UUID) *models.User {
	return controller.store.GetUser(ctx, id)
}
