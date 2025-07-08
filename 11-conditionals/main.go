package main

import "fmt"

func main() {
	fmt.Println("welcome to conditionals")

	age:=28
	var result string
	if age>18 {
		result = "you are an adult"
	}else{
		result="you are a kid"
	}
	fmt.Println(result)
}
