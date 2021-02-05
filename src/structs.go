package main

import (
	"fmt"
)

type courseMeta struct {
	Author string
	Level  string
	Rating float64
}

func main() {
	fmt.Println("structs")

	ThisCourse := courseMeta{
		Author: "Someone",
		Level:  "Aiai",
		Rating: 12.3,
	}

	fmt.Println(ThisCourse)
}
