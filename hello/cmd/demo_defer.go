package main

import "fmt"

func main() {
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
