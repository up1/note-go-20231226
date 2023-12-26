package demo_test

import (
	"demo"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := demo.SayHi()
	if result != "Hello World!" {
		t.Errorf("SayHi() failed, expected %v, got %v", "Hello World!", result)
	}
}
