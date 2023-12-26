package demo

import "fmt"

func TryError() (string, error) {
	return "Hello World!", fmt.Errorf("Error: %s", "This is an error")
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Error: %s", "Division by zero")
	}
	return a / b, nil
}
