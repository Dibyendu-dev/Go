package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to Golang")
	fmt.Println("ple rate our app betwwen 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating")

	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating",input)

	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("added to your rating", numRating)
	}
}