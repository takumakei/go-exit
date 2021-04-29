package types

// StatusCoder is the interface of type that has an status code.
type StatusCoder interface {
	// StatusCode returns the status code.
	StatusCode() int
}
