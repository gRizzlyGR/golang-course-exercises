package main

import "fmt"

type person struct {
	name string
	age  int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hi, I'm", p.name, "and I'm", p.age, "old.")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	me := person{
		name: "Giuseppe",
		age:  31,
	}

	me.speak()        // working
	(&me).speak()     // working
	saySomething(&me) // working
	//saySomething(me) // not working
}
