package main

import (
	"fmt"
	"io"

	"net/http"
)
const url ="https://ddasdev.vercel.app/"

func main() {
	fmt.Println("welcome to in the web ")

	response , err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("response type : %T \n", response)

	defer response.Body.Close()
	databytes,err := io.ReadAll(response.Body)
	
	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}