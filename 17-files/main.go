package main

import (
	"fmt"
	"io"
	
	"os"
)

func main() {
	fmt.Println("Welcme to files world..")
	content:= "hey, what's ur name. my name is = ddas"

	file , err := os.Create("./mygofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	 length,err := io.WriteString(file,content)
	checkNilErr(err)
	
	fmt.Println("length is ",length)
	defer file.Close()
	readfile("./mygofile.txt")
}

func readfile(filename string)  {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("text data is :\n",string(databyte))
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}