package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func main() {
	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	fmt.Println(timmy)

	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)
}
