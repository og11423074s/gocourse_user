package user

import (
	"errors"
	"fmt"
)

var ErrFirstNameRequired = errors.New("first name is required")
var ErrLastNameRequired = errors.New("last name is required")

type ErrorNotFound struct {
	UserID string
}

func (e ErrorNotFound) Error() string {
	return fmt.Sprintf("user %s doesn't exist", e.UserID)
}
