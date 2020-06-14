package main

import "fmt"

func reclassify(planets *[]string) {
	*planets = (*planets)[0:3]
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	fmt.Println(planets)
	reclassify(&planets)
	fmt.Println(planets)
}
