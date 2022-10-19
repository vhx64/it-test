package domain

const (
	// bad request errors
	ErrorRequiredOrderParameter      = "error.required_order_parameter"
	ErrorInvalidOrderParameter       = "error.invalid_order_parameter"
	ErrorInvalidRequestParams        = "error.invalid_request_params"
	ErrorInvalidRequestBodyParameter = "error.invalid_request_body_parameter"
	ErrorBadRequestErrorLabel        = "error.bad_request_error"
	ErrorRequiredUserDetails         = "error.required_user_details"
	// conflict
	// unauthorized errors
	ErrorUserUnauthorizedLabel = "error.user_unauthorized"

	// unprocessable entity errors
	ErrorInvalidRequestBodyLabel = "error.invalid_request_body"

	// not found errors

	// internal errors
	ErrorInternalServerErrorLabel = "error.internal_server_error"

	// http.ok
	UpdateSuccessfully = "Update successfully!"

	// user errors
	ErrorUserNotFound = "error.user_not_found"
)
