package query

import (
	"context"
	"it-test/pkg/logs"
)

type GetUserCountRepository interface {
	GetUserCount(ctx context.Context) (int, error)
}

type GetUserCountHandler struct {
	repo GetUserCountRepository
}

type GetUserCount struct {
}

func NewGetUserCountHandler(repo GetUserCountRepository) *GetUserCountHandler {
	return &GetUserCountHandler{repo: repo}
}

func (h *GetUserCountHandler) Handle(ctx context.Context, cmd *GetUserCount) (count int, err error) {
	defer func() {
		logs.LogCommandExecution("GetUserCountHandler", cmd, err)
	}()

	return h.repo.GetUserCount(ctx)
}
