package demo

import "fmt"

func TryError() (string, error) {
	return "Hello World!", fmt.Errorf("Error: %s", "This is an error")
}
