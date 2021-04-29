package types

import (
	"fmt"
)

// NewStatus returns Status that contains code.
func NewStatus(code int) Status {
	return Status{code: code}
}

// Status represents a status code as an error.
type Status struct {
	code int
}

// Status must be an error.
var _ error = Status{0}

// Status msut be a StatusCoder.
var _ StatusCoder = Status{0}

// Error returns the string of "status {{exit.StatusCode()}}".
func (exit Status) Error() string {
	return fmt.Sprintf("status %d", exit.code)
}

// StatusCode returns the status code that exit contains.
func (exit Status) StatusCode() int {
	return exit.code
}
