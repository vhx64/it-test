package query

import (
	"context"
	"errors"
	"fmt"
	"it-test/pkg/logs"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserRepository interface {
	UpdateUser(ctx context.Context, params UpdateDbUser) (GetUser, error)
}

type UpdateUserHandler struct {
	repo UpdateUserRepository
}

type UpdateUser struct {
	FirstName     string
	LastName      string
	Mobile        string
	Password      string
	PasswordCheck string
	UserName      string
}

type UpdateDbUser struct {
	FirstName string
	Id        openapi_types.UUID
	LastName  string
	Mobile    string
	Password  string
	UserName  string
}

func NewUpdateUserHandler(repo UpdateUserRepository) *UpdateUserHandler {
	return &UpdateUserHandler{repo: repo}
}

func (h *UpdateUserHandler) Handle(ctx context.Context, id string, cmd *UpdateUser) (u GetUser, err error) {
	defer func() {
		logs.LogCommandExecution("UpdateUserHandler", cmd, err)
	}()

	if cmd.Password != cmd.PasswordCheck {
		err = errors.New("password check is different from the password")
		return
	}

	var pw []byte
	if pw, err = bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost); err != nil {
		err = fmt.Errorf("error encrypting pw: %w", err)
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		err = errors.New("invalid id")
		return
	}
	return h.repo.UpdateUser(ctx, UpdateDbUser{
		Id:        uid,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Mobile:    cmd.Mobile,
		Password:  string(pw),
		UserName:  cmd.UserName,
	})
}
