package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")
	// why do pointers exist?
	// pointers are used to store the memory address of another variable

	var ptr *int     // declaring a pointer
	fmt.Println(ptr) // <nil> ,default value of a pointer is nil

	myNumber := 23
	// reference means &
	ptr = &myNumber                               // & is used to get the memory address of a variable
	fmt.Println("Value of assigned pointer", ptr) // 0xc0000b4008 - this implies ptr is reference to direct memory location
	fmt.Println("Value of the pointer", *ptr)     // 23 - this implies *ptr is the value at the memory location ptr is pointing to

	// operation

	*ptr = *ptr + 3
	fmt.Println("New value of the pointer", *ptr) // 26
}
