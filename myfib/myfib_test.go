package myfib

import (
	"testing"
)

func Test1(t *testing.T) {

	if Fib(1) != 1 {
		t.Error("Expected 0, got ", Fib(0))
	}
}
