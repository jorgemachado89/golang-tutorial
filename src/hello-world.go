package main

import ( 
	"fmt" // part of Standard library
	"os"
	//"runtime" 
	//"reflect"
)

/*var (
	name, course	string
	module			float64
)*/

/*var (
	name	= "Nigel"
	course	= "Docker Deep Dive"
	module	= 3.2
)*/

// main is a reserved method and expected program's entry point except if it is a shared lib.
func main() {
	// var name string = "Nigel"
	// var course string = "Docker Deep Dive"
	// var module float64 = 3.2;

	// a	:= 10.00000
	// b	:= 3

	// fmt.Println("a is of type:\t", reflect.TypeOf(a), "\tb is of type:\t", reflect.TypeOf(b))

	// c := int(a) + b
	// 
	// fmt.Println("a is of type:\t", reflect.TypeOf(a), "\tc is of type:\t", reflect.TypeOf(c))
	
	name	:= os.Getenv("USER")
	// module	:= 3.2
	// pointer := &module
	course	:= "Docker Deep Dive"
	fmt.Println("course : ", course)


	// fmt.Println("Name:\t", name, "\tType: ", reflect.TypeOf(name))
	// fmt.Println("Module:\t", module, "\tType: ", reflect.TypeOf(module))
	// fmt.Println("Module address:\t", pointer, "\t and it's value is: ", *pointer)

	fmt.Println("name : ", name)
	fmt.Println("Changed? : ", changeCourse(&course))
	fmt.Println("Course Changed? : ", course)

	//fmt.Println("Os Envs: ", os.Environ())

	/*for _, env := range os.Environ() {
		fmt.Println(env)
	}*/
}

func changeCourse(someParam *string) string {
	*someParam = "Hello Hello"
	fmt.Println("Changing param to: ", *someParam)

	return *someParam
}