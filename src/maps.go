package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	leagueTitles := make(map[string]int)
	leagueTitles["Vitoria"] = 1
	leagueTitles["Porto"] = 30

	fmt.Println(leagueTitles)

	premierLeagueTitles := map[string]int{
		"ManUnited": 20,
		"Arsenal":   5,
	}

	fmt.Println(premierLeagueTitles)
}
