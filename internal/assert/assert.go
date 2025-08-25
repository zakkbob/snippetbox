package assert

import "testing"

func Equal[T comparable](t *testing.T, actual T, expected T) {
	t.Helper()

	if expected != actual {
		t.Errorf("got %v; want %v", actual, expected)
	}
}
