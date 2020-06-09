package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//randomly pick
	bank := 0.0
	for bank < 20.0 {
		switch k := rand.Intn(3) + 1; k {
		case 1:
			bank += 0.05
			break
		case 2:
			bank += 0.1
			break
		case 3:
			bank += 0.25
			break
		}
	}

	fmt.Printf("%5.2f\n", bank)
}
