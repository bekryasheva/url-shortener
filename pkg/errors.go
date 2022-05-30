package pkg

import (
	"errors"
)

var ErrNotFound = errors.New("url is not found")

var ErrInvalidInput = errors.New("input validation failed")
