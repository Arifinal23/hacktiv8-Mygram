package helper

type BaseResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

func NewResponse(status int, data interface{}, errors error) *BaseResponse {
	if errors != nil && !isValidationError(errors) {
		return &BaseResponse{
			Status: status,
			Data:   data,
			Errors: map[string]interface{}{"message": errors.Error()},
		}
	}

	return &BaseResponse{
		Status: status,
		Data:   data,
		Errors: errors,
	}
}

// ExampleErrorResponse only for example swaggo docs
type ExampleErrorResponse struct {
	Message  string `json:"message" example:"user bad request"`
	Email    string `json:"email" example:"cannot be blank"`
	Password string `json:"password" example:"cannot be blank"`
}
