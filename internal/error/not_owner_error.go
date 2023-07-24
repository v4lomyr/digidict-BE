package error

import (
	"fmt"
	"net/http"
)

type NotOwnerError struct {
	owned string
	owner string
	id    interface{}
}

func (e *NotOwnerError) Error() string {
	return fmt.Sprintf("%s not owned by %s, id of %v", e.owned, e.owner, e.id)
}

func NewNotOwnerError(owned string, owner string, id interface{}) error {
	return NewClientError(&NotOwnerError{
		owned: owned,
		owner: owner,
		id:    id,
	}, http.StatusNotFound)
}
