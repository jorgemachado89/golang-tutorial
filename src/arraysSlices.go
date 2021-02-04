package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays and Slices")

	// arraysAndSlices()
	// sliceOfSlice()
	appendSlice()
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

func appendSlice() {

	mySlice := make([]int, 1, 4)
	fmt.Printf("Length is: %d Capacity is: %d",
		len(mySlice), cap(mySlice))

	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("\nCapacity is: %d", cap(mySlice))
	}

}
