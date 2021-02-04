package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")

	ifConditional()
	switchStatment()
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