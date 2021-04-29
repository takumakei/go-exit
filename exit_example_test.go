package exit_test

import (
	"errors"
	"fmt"
	"os"

	"github.com/takumakei/go-exit"
)

func Example() {
	// Printing the status code instead of calling os.Exit just for this example.
	exit.HandleExit = func(code int) {
		fmt.Printf("status %d", code)
	}

	exit.Exit(os.Stdout, run())
	// Output:
	// error: Deep Thought said 42
	// status 42
}

func run() error {
	return exit.Error(42, errors.New("Deep Thought said 42"))
}
