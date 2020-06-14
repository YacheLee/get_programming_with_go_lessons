package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func main() {
	scott := person{
		name:       "Scott",
		superpower: "programming",
		age:        14,
	}
	birthday(&scott)
	fmt.Printf("%+v\n", scott)
}
