package main

import "fmt"

func main() {
	dwarfs := [...]string{
		"Ceres",
		"Pluto",
		"Haumea",
		"Makemake",
		"Eris",
	}

	for _, drwarf := range dwarfs {
		fmt.Println(drwarf)
	}
}
