package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")
	
	var fruitList [4]string

	fruitList[0]="apple"
	fruitList[1]="mango"
	fruitList[3]="pineapple"
	
	fmt.Println("Fruit List is :",fruitList);
	fmt.Println("Fruit List  Length is :",len(fruitList));

	var veg = [3]string{"potato","onion","ginger"}
	fmt.Println("veg list are",veg)

}