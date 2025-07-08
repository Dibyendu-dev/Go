package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop")
	days := []string{"sunday","monday","tuesday","wednesday","thrusday","friday","saturday"}
	fmt.Println(days)

	// for i:=0; i<len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	// for i , day := range days {
	// 	fmt.Printf("index is %v, value is %v \n",i,day)
	// }

	i:=1
	for i<10{
		if i == 5{
			// break
			i++
			continue
		}
		fmt.Println("value is ", i)
		i++
	}

}