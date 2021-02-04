package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Hello")

	// ifConditional()
	// switchStatment()
	checkForErrorHandling()
}

func ifConditional() {
	if firstRank, secondRank := "10", "11"; firstRank < secondRank {
		fmt.Println("firstRank: ", firstRank)
	}
}

func switchStatment() {
	switch value := "Ola"; value {
	case "Docker Deep Dive":
		fmt.Println(value)
	case "Ole", "Ola":
		fmt.Println("elO")
		fmt.Println("Olare")
		fallthrough
	case "Coisas":
		fmt.Println("pois e")
	default:
		fmt.Println("Default")
	}
}

func checkForErrorHandling() {
	if _, err := os.Open("../README.md"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No errors")
	}
}