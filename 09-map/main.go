package main

import "fmt"

func main() {
	fmt.Println("Welcome to map")

	languages := make(map[string]string)

	languages["js"] = "javaScript"
	languages["py"] = "python"
	languages["go"] = "golang"
	fmt.Println("prog. language are: ",languages)
	fmt.Println("Go language at: ",languages["go"])

	// delete(languages,"py")
	// fmt.Println("prog. language are: ",languages)

	for key,value :=range languages {
		fmt.Printf("for key %v :: value is %v \n" ,key,value)
	}

}