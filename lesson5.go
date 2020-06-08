package main

import (
	"fmt"
	"math/rand"
)

const distance = 62100000
const secondsPerDay = 86400

func getCompany() string {
	langs := [3]string{
		"Space Adventures",
		"SpaceX",
		"Virgin Galactic",
	}
	n := rand.Intn(3)
	return langs[n]
}

func getDuration(speed int) int {
	return distance / speed / secondsPerDay
}

func getTrip() string {
	if rand.Intn(2) == 1 {
		return "Round-trip"
	} else {
		return "One-way"
	}
}

func getPrice(speed int) int {
	return speed + 20.0
}

func main() {
	fmt.Printf("Spaceline         Days Trip  type Price\n")
	for i := 0; i < 10; i++ {
		speed := rand.Intn(15) + 16
		fmt.Printf("%-16v %4v %-10v $%4v \n", getCompany(), getDuration(speed), getTrip(), getPrice(speed))
	}
}
