package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct")
	ddas := User{"Dibyendu","ddas4548@gmail.com",true,28}
	fmt.Println("user is", ddas)
	fmt.Printf("ddas details are: %+v\n",ddas)
	fmt.Printf(" name is :%v \n age is %v ",ddas.Name,ddas.Age)

}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}