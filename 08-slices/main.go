package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"apple","mango","banana"}
	fmt.Printf("type of fruitlist is %T \n ", fruitList)

	fruitList = append(fruitList, "Guava","pineapple")

	fmt.Println("fruiList are", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("fruiList are", fruitList)

	highScores := make([]int,4)
	highScores[0]= 199
	highScores[1]= 699
	highScores[2]= 399
	highScores[3]= 599
	highScores = append(highScores, 999,299,799,899,499)
	fmt.Println("HighScores: ",highScores)
	sort.Ints(highScores)

	fmt.Println("HighScores: ",highScores)

	 var courses = []string{"react","node","mongo","express"}
	 fmt.Println(courses)
	 var index int = 2
	 courses = append(courses[:index],courses[index+1:]...)
	 fmt.Println(courses)

}