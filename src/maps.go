package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	/*leagueTitles := make(map[string]int)
	leagueTitles["Vitoria"] = 1
	leagueTitles["Porto"] = 30

	fmt.Println(leagueTitles)

	premierLeagueTitles := map[string]int{
		"ManUnited": 20,
		"Arsenal":   5,
	}

	fmt.Println(premierLeagueTitles)*/

	itereateOrdering()
}

func itereateOrdering() {
	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
	}

	for key, value := range testMap {
		fmt.Printf("key: %v, value: %v \n", key, value)
	}
}
