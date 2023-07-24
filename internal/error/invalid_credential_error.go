package error

import "net/http"

type InvalidCredential struct {
}

func (e *InvalidCredential) Error() string {
	return "Invalid Crendential"
}

func NewInvalidCredentialError() error {
	return NewClientError(&InvalidCredential{}, http.StatusBadRequest)
}
