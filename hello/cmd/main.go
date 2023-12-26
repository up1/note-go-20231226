package main

import "demo"

func main() {
	r, err := demo.TryError()
	if err != nil {
		println(err.Error())
	}
	println(r, 1/0)

	result := demo.SayHi()
	println(result)
}
