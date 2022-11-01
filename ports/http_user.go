package ports

import (
	"encoding/json"
	"it-test/app/query"
	"it-test/domain"
	"it-test/pkg/server/httperr"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (h HTTPServer) PostUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var u query.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		httperr.BadRequest(*httperr.NewErrorMessageBody(domain.ErrorBadRequestErrorLabel, "client", postUser, uuid.NewString()), err, w, r)
		return
	}
	response, err := h.app.Queries.CreateUser.Handle(ctx, &u)
	if err != nil {
		httperr.InternalError(domain.ErrorInternalServerErrorLabel, postUser, uuid.NewString(), err, w, r)
		return
	}
	render.Respond(w, r, response)
}

func (h HTTPServer) GetUserList(w http.ResponseWriter, r *http.Request, params GetUserListParams) {
	ctx := r.Context()
	result, err := h.app.Queries.GetUserList.Handle(ctx, &query.GetUserList{
		EmailFilter: params.EmailFilter,
		PageIndex:   params.PageIndex,
		Limit:       params.Limit,
		OrderBy:     params.OrderBy,
		Order:       params.Order,
	})
	if err != nil {
		httperr.InternalError(domain.ErrorInternalServerErrorLabel, getUserList, uuid.NewString(), err, w, r)
		return
	}
	response := UserList{ResultsLength: len(result), Results: make([]UserListItem, 0, len(result))}
	for _, r := range result {
		response.Results = append(response.Results, UserListItem{
			Aszf:      r.Aszf,
			Email:     r.Email,
			FirstName: r.FirstName,
			Id:        r.Id,
			LastName:  r.LastName,
			Mobile:    r.Mobile,
			UserName:  r.UserName,
		})
	}
	render.Respond(w, r, response)
}

func (h HTTPServer) UpdateUserDetails(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()
	var u query.UpdateUser
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		httperr.BadRequest(*httperr.NewErrorMessageBody(domain.ErrorBadRequestErrorLabel, "client", updateUserDetails, uuid.NewString()), err, w, r)
		return
	}
	response, err := h.app.Queries.UpdateUser.Handle(ctx, id, &u)
	if err != nil {
		httperr.InternalError(domain.ErrorInternalServerErrorLabel, updateUserDetails, uuid.NewString(), err, w, r)
		return
	}
	render.Respond(w, r, response)
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
