package types

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewError(t *testing.T) {
	t.Run("NewError with err = nil", func(t *testing.T) {
		defer func() {
			v := recover()
			if v == nil {
				t.Fatal("panic wanted")
			}
			err, ok := v.(error)
			if !ok {
				t.Fatal("error wanted")
			}
			want := "runtime error: invalid memory address or nil pointer dereference"
			if diff := cmp.Diff(want, err.Error()); diff != "" {
				t.Errorf("-want +got\n%s", diff)
			}
		}()

		// never calling NewError with err = nil.
		_ = NewError(0, nil).Error()
	})
}
