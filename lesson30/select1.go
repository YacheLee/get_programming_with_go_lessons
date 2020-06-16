package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.After(2 * time.Second)
	c := make(chan int)
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}

	fmt.Println("asdad")
}
