package main

import "fmt"

type report struct {
	sol         int
	temperature temperature
	location    location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius int

func main() {
	taiwan := location{25.105497, 121.597366}
	t := temperature{-1.0, -78}
	report := report{sol: 15, temperature: t, location: taiwan}
	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %vº C\n", report.temperature.high)

}
