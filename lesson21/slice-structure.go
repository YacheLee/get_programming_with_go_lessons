package main

import "fmt"

func main() {
	type location struct {
		name string
		lat  float64
		long float64
	}

	locations := []location{
		{name: "Bradbury Landing", lat: -4.5555, long: 5555},
	}

	fmt.Printf("%+v", locations)
}
