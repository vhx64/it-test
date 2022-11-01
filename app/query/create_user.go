package query

import (
	"context"
	"errors"
	"fmt"
	"it-test/pkg/logs"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRepository interface {
	CreateUser(ctx context.Context, params CreateDbUser) (GetUser, error)
}

type CreateUserHandler struct {
	repo CreateUserRepository
}

type CreateUser struct {
	Aszf          bool
	Email         string
	FirstName     string
	LastName      string
	Mobile        string
	Password      string
	PasswordCheck string
	UserName      string
}

type CreateDbUser struct {
	Aszf      bool
	Email     string
	FirstName string
	Id        openapi_types.UUID
	LastName  string
	Mobile    string
	Password  string
	UserName  string
}

type GetUser struct {
	Aszf      bool
	Email     string
	FirstName string
	Id        openapi_types.UUID
	LastName  string
	Mobile    string
	UserName  string
}

func NewCreateUserHandler(repo CreateUserRepository) *CreateUserHandler {
	return &CreateUserHandler{repo: repo}
}

func (h *CreateUserHandler) Handle(ctx context.Context, cmd *CreateUser) (u GetUser, err error) {
	defer func() {
		logs.LogCommandExecution("CreateUserHandler", cmd, err)
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

	return h.repo.CreateUser(ctx, CreateDbUser{
		Aszf:      cmd.Aszf,
		Email:     cmd.Email,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Mobile:    cmd.Mobile,
		Password:  string(pw),
		UserName:  cmd.UserName,
	})
}
