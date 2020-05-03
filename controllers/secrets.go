package controllers

import (
	"auth/models"
	"context"
	uuid "github.com/satori/go.uuid"
)

func (controller *Controller) GetSecret(ctx context.Context, id uuid.UUID) *models.Secret {
	secret := controller.store.GetSecret(ctx, id)
	return secret
}

func (controller *Controller) GetActualSecret(ctx context.Context) *models.Secret {
	secret := controller.store.GetActualSecret(ctx)
	if secret != nil {
		return secret
	}

	secret = controller.store.CreateSecret(ctx)
	controller.OnSecretCreatedV1(secret)
	return secret
}
