package main

import "fmt"

func main() {
	fmt.Println("Welcome to functional world...")
	greet()
	result := add(5,8)
	fmt.Println("add of both num is",result)

	prores := proadd(3,6,8,4,5,8,5,7,58,5,7,5,8,5,5,8,5,7)
	fmt.Println("add of all num is",prores)

}

func add (a int , b int) int {
	return a + b
}

func proadd (value ...int) int{
	total := 0

	for _,val := range value{
		total +=val
	}
	return  total
}

func greet (){
	fmt.Println("hello dibyendu")
}