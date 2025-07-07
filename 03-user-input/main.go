package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"

	fmt.Println(welcome)

	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Enter your Age:")

	// comma ok || error 
	input, _ :=reader.ReadString('\n')
	fmt.Println("Thanks for telling your age is :",input)
	fmt.Printf("Type your age is :%T",input)

}