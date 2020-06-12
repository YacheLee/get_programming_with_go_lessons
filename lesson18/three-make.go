package main

import "fmt"

func main() {
	dwarfs := make([]string, 0, 10)
	fmt.Println(dwarfs, len(dwarfs), cap(dwarfs))
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris", "Makemake", "Eris", "Makemake", "Eris", "Makemake", "Eris", "Makemake", "Eris", "Makemake", "Eris", "Makemake", "Eris", "Makemake", "Eris")
	fmt.Println(dwarfs, len(dwarfs), cap(dwarfs))
}
