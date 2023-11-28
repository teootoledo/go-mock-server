package resources

// ErrorResponse Error model
// @Description Error base info.
type ErrorResponse struct {
	Message string `json:"message" example:"status code message"`
	Details string `json:"details" example:"error details"`
}

// ErrorResponseHTTPBadRequest Error model
// @Description Error response for http BadRequest.
type ErrorResponseHTTPBadRequest struct {
	ErrorResponse ErrorResponse `json:"error-response"`
	StatusCode    int           `json:"status-code" example:"400"`
}

// ErrorResponseHTTPNotFound Error model
// @Description Error response for http NotFound.
type ErrorResponseHTTPNotFound struct {
	ErrorResponse ErrorResponse `json:"error-response"`
	StatusCode    int           `json:"status-code" example:"404"`
}

// ErrorResponseHTTPInternalServerError Error model
// @Description Error response for http InternalServerError.
type ErrorResponseHTTPInternalServerError struct {
	ErrorResponse ErrorResponse `json:"error-response"`
	StatusCode    int           `json:"status-code" example:"500"`
}
