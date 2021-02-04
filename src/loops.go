package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")
	countdownTimer()
}

func countdownTimer() {
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}
}