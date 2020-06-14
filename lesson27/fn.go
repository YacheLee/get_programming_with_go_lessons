package main

import "fmt"

var fn func(a, b int) int

func main() {
	fmt.Println(fn == nil)
}
