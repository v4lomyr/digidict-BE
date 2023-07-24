package error

import "net/http"

type InvalidParamError struct {
	field string
}

func (e *InvalidParamError) Error() string {
	return "invalid param of field: " + e.field
}

func NewInvalidParamError(field string) error {
	return NewClientError(&InvalidParamError{field: field}, http.StatusUnauthorized)
}
