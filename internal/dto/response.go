package dto

type ErrorResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type ValidationErrorResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}
