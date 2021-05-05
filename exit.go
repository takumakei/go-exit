package exit

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/takumakei/go-exit/types"
)

// Status returns the error that represents the status code.
func Status(code int) types.Status {
	return types.NewStatus(code)
}

// Error returns the error that wraps the err with the status code.
//
//   returns nil if err == nil and code == 0.
//   returns Status(code) if err == nil.
//   returns types.NewError(code, err), otherwise.
func Error(code int, err error) error {
	if err == nil {
		if code == 0 {
			return nil
		}
		return Status(code)
	}
	return types.NewError(code, err)
}

// HandleExit is os.Exit.
// This should not be changed not in case of testing.
var HandleExit = os.Exit

// Format is the format used by Exit to write an error.
var Format = "error: %v\n"

// Stderr is the writer that Exit writes a message to.
var Stderr io.Writer = os.Stderr

// Exit calls Fexit(Stderr, err).
func Exit(err error) {
	Fexit(Stderr, err)
}

// ExitOnError calls Exit(err) if err != nil, otherwise do nothing.
func ExitOnError(err error) {
	if err != nil {
		Exit(err)
	}
}

// Fexit writes the error message of err to w, then calls os.Fexit.
//
// The error message is not written in case of that
// `err is nil` nor `err is types.Status itself`.
//
// If you do not want to output any error message under any circumstances,
// consider the following.
// `os.Exit(StatusCode(err))`
func Fexit(w io.Writer, err error) {
	if err == nil {
		// no error, no error message.
		HandleExit(0)
	}

	code := 1
	// if and only if the err is types.Status itself, exit with no error message.
	if exit, ok := err.(types.Status); ok {
		code = exit.StatusCode()
	} else {
		fmt.Fprintf(w, Format, err)
		if exit, ok := LookupStatusCoder(err); ok {
			code = exit
		}
	}
	HandleExit(code)
}

// FexitOnError calls Fexit(w, err) if err != nil, otherwise do nothing.
func FexitOnError(w io.Writer, err error) {
	if err != nil {
		Fexit(w, err)
	}
}

// LookupStatusCoder returns the StatusCode() of types.StatusCoder with ok = true
// in case of the target contains a types.StatusCoder.
func LookupStatusCoder(target error) (code int, ok bool) {
	var exit types.StatusCoder
	if errors.As(target, &exit) {
		code, ok = exit.StatusCode(), true
	}
	return
}

// StatusCode returns the status code for err.
//
//   returns 0 if err == nil.
//   returns 1 if err is not a types.StatusCoder.
//   returns StatusCode() of types.StatusCoder, otherwise.
func StatusCode(err error) int {
	if err == nil {
		return 0
	}
	if code, ok := LookupStatusCoder(err); ok {
		return code
	}
	return 1
}
