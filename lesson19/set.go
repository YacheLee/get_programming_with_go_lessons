package main

import "fmt"

func main() {
	var termperatures = []float32{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float32]bool)

	for _, t := range termperatures {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	fmt.Println(set)
}
