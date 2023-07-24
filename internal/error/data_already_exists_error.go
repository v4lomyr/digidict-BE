package error

import (
	"fmt"
)

type DataAlreadyExistsError struct {
	entity string
	field  string
	values any
}

func (e *DataAlreadyExistsError) Error() string {
	return fmt.Sprintf("%s already exists with %s of %v", e.entity, e.field, e.values)
}

func NewDataAlreadyExistsError(entity string, field string, values any) error {
	return NewClientError(&DataAlreadyExistsError{
		entity: entity,
		field:  field,
		values: values,
	})
}
