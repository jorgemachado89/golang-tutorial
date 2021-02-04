package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays and Slices")

	// arraysAndSlices()
	sliceOfSlice()
}

func arraysAndSlices() {

	// By doing so an Array is created to hold the value of the Slice
	// since the Slice holds only references to an Array
	myCourses := make([]string, 5, 10)

	fmt.Printf("Capacity: %d and Lenght %d \n",
		cap(myCourses), len(myCourses))

}

func sliceOfSlice() {

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice[4])

	mySlice[1] = 0
	fmt.Println(mySlice)

	sliceOfSlice := mySlice[2:5]
	fmt.Println(sliceOfSlice)

}
