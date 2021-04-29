package types

// NewError returns *Error that contains code and err.
// err must not be nil.
func NewError(code int, err error) *Error {
	return &Error{code: code, err: err}
}

// Error represents a status code as an error containing another error.
type Error struct {
	code int
	err  error
}

// *Error must be an error.
var _ error = (*Error)(nil)

// *Error must be a StatusCoder.
var _ StatusCoder = (*Error)(nil)

// Error returns the same result of exit.Unwrap().Error().
func (exit *Error) Error() string {
	return exit.err.Error()
}

// Unwrap returns the error that exit contains.
func (exit *Error) Unwrap() error {
	return exit.err
}

// StatusCode returns the status code that exit contains.
func (exit *Error) StatusCode() int {
	return exit.code
}
