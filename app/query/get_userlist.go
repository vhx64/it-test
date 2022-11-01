package query

import (
	"context"
	"it-test/pkg/logs"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

type GetUserListRepository interface {
	GetUserList(ctx context.Context, params GetUserList) ([]GetUserListItem, error)
}

type GetUserListHandler struct {
	repo GetUserListRepository
}

type Limit = int
type Order = string
type OrderBy = string
type PageIndex = int

type GetUserList struct {
	EmailFilter *string
	PageIndex   PageIndex
	Limit       Limit
	OrderBy     OrderBy
	Order       Order
}

type GetUserListItem struct {
	Aszf      bool
	Email     string
	FirstName string
	Id        openapi_types.UUID
	LastName  string
	Mobile    string
	UserName  string
}

func NewGetUserListHandler(repo GetUserListRepository) *GetUserListHandler {
	return &GetUserListHandler{repo: repo}
}

func (h *GetUserListHandler) Handle(ctx context.Context, cmd *GetUserList) (u []GetUserListItem, err error) {
	defer func() {
		logs.LogCommandExecution("GetUserListHandler", cmd, err)
	}()

	return h.repo.GetUserList(ctx, *cmd)
}
