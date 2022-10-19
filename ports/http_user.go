package ports

import (
	"it-test/app/query"
	"it-test/domain"
	"it-test/pkg/server/httperr"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (h HTTPServer) PostUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h HTTPServer) GetUserList(w http.ResponseWriter, r *http.Request, params GetUserListParams) {
	//TODO implement me
	panic("implement me")
}

func (h HTTPServer) UpdateUserDetails(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (h HTTPServer) Count(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := h.app.Queries.GetUserCount.Handle(ctx, &query.GetUserCount{})
	if err != nil {
		httperr.InternalError(domain.ErrorInternalServerErrorLabel, getUserCount, uuid.NewString(), err, w, r)
		return
	}
	response := Count{Count: &result}

	render.Respond(w, r, response)
}
