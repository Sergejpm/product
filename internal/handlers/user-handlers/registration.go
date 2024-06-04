package userhandlers

import (
	"context"
	usercommands "github.com/sergejpm/product/internal/commands/user-commands"
	"github.com/sergejpm/product/internal/domain/service/registration"
)

type RegistrationHandler struct {
	service *registration.Service
}

func NewRegistrationHandler(service *registration.Service) *RegistrationHandler {
	return &RegistrationHandler{service: service}
}

func (h RegistrationHandler) Handle(ctx context.Context, cmd *usercommands.RegistrationCommand) error {
	return h.service.CreateUser(ctx, cmd.Username, cmd.Password)
}
