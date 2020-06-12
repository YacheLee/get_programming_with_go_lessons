package main

import "fmt"

//completely useless
func terraform(planets [5]string) {
	for i, planet := range planets {
		planets[i] = "new " + planet
	}
}

func main() {
	planets := [...]string{
		"Ceres",
		"Pluto",
		"Haumea",
		"Makemake",
		"Eris",
	}

	terraform(planets)

	fmt.Println(planets)
}
