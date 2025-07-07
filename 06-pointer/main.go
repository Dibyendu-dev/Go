package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer")

	// var ptr *int
	// fmt.Println("value of pointer is", ptr)

	 mynumber := 28

	 var ptr = &mynumber

	 fmt.Println("value of pointer is",ptr)
	 fmt.Println("value of pointer is",*ptr)

	 *ptr = *ptr *2
	 fmt.Println("value of mynumber is ", mynumber)

}