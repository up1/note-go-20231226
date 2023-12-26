package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		defer func(x int) {
			println(x)
		}(i)
	}

	defer trySth()
	println("hello world")
	fmt.Println("hello world")
	panic("panic test")
}

func trySth() {
	println("try something")
	r := recover()
	println(r.(string))
}
