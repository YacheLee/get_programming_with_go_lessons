package main

import "fmt"

func main() {
	type location struct{ lat, long float64 }
	bradbury := location{-4.585, 137.4417}

	curiosity := bradbury
	curiosity.long = 123
	fmt.Printf("%+v\n", bradbury)
	fmt.Printf("%+v", curiosity)
}
