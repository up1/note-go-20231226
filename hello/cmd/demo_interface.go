package main

type Employee interface {
	work() string
}

type Manager struct {
	name string
}

func (m Manager) work() string {
	return "managing"
}

type Programmer struct {
	name string
}

func (p Programmer) work() string {
	return "programming"
}

func doWork(e Employee) {
	println(e.work())
}

func main() {
	e1 := Manager{"Somkiat"}
	e2 := Programmer{"Somkiat 2"}
	doWork(e1)
	doWork(e2)

}
