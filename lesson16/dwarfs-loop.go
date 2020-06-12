package main

import "fmt"

func main() {
	dwarfs := [5]string{
		"Ceres",
		"Pluto",
		"Haumea",
		"Makemake",
		"Eris",
	}
	for i := 0; i < len(dwarfs); i++ {
		fmt.Println(dwarfs[i])
	}
}
