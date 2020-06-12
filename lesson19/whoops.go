package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}

	planetsMarkII := planets
	planets["Earth"] = "whoops"

	//both maps are impacted by the change
	fmt.Println(planets, planetsMarkII)

	delete(planets, "Earth")
	fmt.Println(planets, planetsMarkII)
}
