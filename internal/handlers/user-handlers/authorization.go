package userhandlers

import (
	"context"
	"github.com/sergejpm/product/internal/commands/user-commands"
	"github.com/sergejpm/product/internal/domain/service/authorization"
)

type AuthorizationHandler struct {
	service *authorization.Service
}

func NewAuthorizationHandler(service *authorization.Service) *AuthorizationHandler {
	return &AuthorizationHandler{service: service}
}

func (h AuthorizationHandler) Handle(ctx context.Context, cmd *usercommands.AuthorizationCommand) (string, error) {
	return h.service.GetToken(ctx, cmd.Username, cmd.Password)
}
