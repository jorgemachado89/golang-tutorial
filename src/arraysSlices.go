package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays and Slices")

	arraysAndSlices()
}

func arraysAndSlices() {

	// By doing so an Array is created to hold the value of the Slice
	// since the Slice holds only references to an Array
	myCourses := make([]string, 5, 10)

	fmt.Printf("Capacity: %d and Lenght %d \n",
		cap(myCourses), len(myCourses))

}
