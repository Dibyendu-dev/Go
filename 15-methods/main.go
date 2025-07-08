package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct")
	ddas := User{"Dibyendu","ddas4548@gmail.com",true,28}
	fmt.Println("user is", ddas)
	ddas.GetStatus()
	ddas.NewEmail()
	fmt.Printf("ddas details are: %+v\n",ddas)


}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus (){
	fmt.Println("user active :", u.Status)
}

func (u User) NewEmail (){
	u.Email = "ddasdev@dev"
	fmt.Println("New email is ", u.Email)
}