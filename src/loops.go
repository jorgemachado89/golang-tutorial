package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops\n")
	//countdownTimer()
	forRangeLoops()
}

func countdownTimer() {
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}
}

func forRangeLoops() {
	courses := []string {"Docker Deep Dive", "Docker Clustering", "Docker and Kubernetes"}

	for _, course := range courses {
		fmt.Println(course)
	}

}