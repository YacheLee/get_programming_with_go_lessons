package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	if moon, ok := temperature["the"]; ok {
		fmt.Printf("On average the Earth is %v C.\n", moon)

	} else {
		fmt.Println(ok)
		fmt.Println("Where is the Moon?")
	}
}
