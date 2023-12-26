package main

import "demo"

func main() {
	r, err := demo.TryError()
	if err != nil {
		println(err.Error())
	}
	println(r)

	result := demo.SayHi()
	println(result)
}
