package main

import "fmt"

func main() {
	yesNo := "true"

	var launch bool

	switch yesNo {
	case "true", "yes", "1":
		launch = true
		break
	case "false", "no", "0":
		launch = false
		break
	default:
		fmt.Println(launch, "is not valid")
	}

	fmt.Println(launch)
}
