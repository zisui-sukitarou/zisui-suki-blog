package apperr

type ErrorResponse struct {
	Err string `json:"error"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Err: err.Error(),
	}
}
