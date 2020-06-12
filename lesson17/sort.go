package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{"  Venus   ", "Earth ", " Mars"}

	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
}
