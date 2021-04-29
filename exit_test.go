package exit_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/takumakei/go-exit"
)

func TestStatus(t *testing.T) {
	var result int
	exit.HandleExit = func(code int) { result = code }
	out := new(strings.Builder)
	code := 42
	exit.Exit(out, exit.Status(code))
	if result != code {
		t.Errorf("result %d != %d", result, code)
	}
	if out.String() != "" {
		t.Errorf("out is not empty: %q", out.String())
	}
}

func TestError(t *testing.T) {
	var result int
	exit.HandleExit = func(code int) { result = code }
	out := new(strings.Builder)
	code := 42
	err := errors.New("Deep Thought said 42")
	exit.Exit(out, exit.Error(code, err))
	if result != code {
		t.Errorf("result %d != %d", result, code)
	}
	want := fmt.Sprintf("error: %v\n", err)
	if diff := cmp.Diff(want, out.String()); diff != "" {
		t.Errorf("-want +got\n%s", diff)
	}
}
