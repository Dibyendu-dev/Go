package main

import "fmt"

const MyName string = "dibyndu"  //public variable

func main() {
	var username string = "ddas"
	fmt.Println(username)
	fmt.Printf("Variable  is of type :%T  \n", username )

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable  is of type :%T  \n", isLoggedIn )

	var num	 int = 265
	fmt.Println(num)
	fmt.Printf("Variable  is of type :%T  \n", num )

	var smallFloat float32 = 265.654
	fmt.Println(smallFloat)
	fmt.Printf("Variable  is of type :%T  \n", smallFloat )

	// implicit type 
	var website = "http://ddasdev.vercel.dev"
	fmt.Println(website);

	// no var style

	numberofUser :=  300
	fmt.Println(numberofUser)

	fmt.Println(MyName)
}