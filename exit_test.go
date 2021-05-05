package exit_test

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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
	exit.Fexit(out, exit.Status(code))
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
	exit.Fexit(out, exit.Error(code, err))
	if result != code {
		t.Errorf("result %d != %d", result, code)
	}
	want := fmt.Sprintf("error: %v\n", err)
	if diff := cmp.Diff(want, out.String()); diff != "" {
		t.Errorf("-want +got\n%s", diff)
	}
}

func TestExitOnError(t *testing.T) {
	var got int
	exit.HandleExit = func(code int) { got = 42 }

	exit.Stderr = ioutil.Discard

	exit.ExitOnError(nil)
	want := 0
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	exit.ExitOnError(io.EOF)
	want = 42
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFexitOnError(t *testing.T) {
	var got int
	exit.HandleExit = func(code int) { got = 42 }

	exit.FexitOnError(ioutil.Discard, nil)
	want := 0
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	exit.FexitOnError(ioutil.Discard, io.EOF)
	want = 42
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
