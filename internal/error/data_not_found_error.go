package error

import (
	"fmt"
	"net/http"
)

type DataNotFoundError struct {
	entity string
	id     interface{}
}

func (e *DataNotFoundError) Error() string {
	return fmt.Sprintf("%s not found for, id of %v", e.entity, e.id)
}

func NewDataNotFoundError(entity string, id interface{}) error {
	return NewClientError(&DataNotFoundError{
		entity: entity,
		id:     id,
	}, http.StatusNotFound)
}
