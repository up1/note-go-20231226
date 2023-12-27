// go test -bench .
package demo_test

import (
	"fmt"
	"testing"
)

func BenchmarkHelloV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprint("Hello, World!")
	}
}

func BenchmarkHelloV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello, World!")
	}
}
