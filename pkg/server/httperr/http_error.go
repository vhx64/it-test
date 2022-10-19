package httperr

import (
	"it-test/pkg/logs"
	"net/http"

	"github.com/go-chi/render"
)

type ErrorMessageBody struct {
	Code          int       `json:"-"`
	CorrelationID string    `json:"correlationId"`
	FunctionCode  string    `json:"functionCode"`
	Messages      []Message `json:"messages"`
}

type Message struct {
	Label        string `json:"label"`
	FromProperty string `json:"fromProperty"`
}

func NewErrorMessageBody(label, fromProperty, functionCode, correlationID string) *ErrorMessageBody {
	var msg = []Message{}
	msg = append(msg, Message{
		Label:        label,
		FromProperty: fromProperty,
	})
	return &ErrorMessageBody{
		Messages:      msg,
		FunctionCode:  functionCode,
		CorrelationID: correlationID,
	}
}

func NotFound(m ErrorMessageBody, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, w, r, http.StatusNotFound, m)
}

func InternalError(label, functionCode, correlationID string, err error, w http.ResponseWriter, r *http.Request) {
	m := NewErrorMessageBody(label, "server", functionCode, correlationID)
	httpRespondWithError(err, w, r, http.StatusInternalServerError, *m)
}

func Unauthorized(label, functionCode, correlationD string, err error, w http.ResponseWriter, r *http.Request) {
	m := NewErrorMessageBody(label, "client", functionCode, correlationD)
	httpRespondWithError(err, w, r, http.StatusUnauthorized, *m)
}

func BadRequest(m ErrorMessageBody, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, w, r, http.StatusBadRequest, m)
}

func UnprocessableEntityError(m ErrorMessageBody, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, w, r, http.StatusUnprocessableEntity, m)
}

func httpRespondWithError(err error, w http.ResponseWriter, r *http.Request, statusCode int, m ErrorMessageBody) {
	logs.GetLogEntry(r).
		WithError(err).
		WithField("errors", m).
		Warn(http.StatusText(statusCode))
	resp := ErrorMessageBody{
		Code:          statusCode,
		CorrelationID: m.CorrelationID,
		FunctionCode:  m.FunctionCode,
		Messages:      m.Messages,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

func (e ErrorMessageBody) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.Code)
	return nil
}
