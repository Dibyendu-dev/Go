package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to future..")

	presentime := time.Now()
	
	fmt.Println("Today Date:",presentime.Format("01-02-2006 15:04:05 Monday"))

	covidStarDate := time.Date(2020,time.March,25,06,00,0,0,time.UTC)

	
	fmt.Println("Lockdown starts in India:",covidStarDate.Format("01-02-2006 15:04:05 Monday"))

	timepassed := presentime.Sub(covidStarDate)
	days := int(timepassed.Hours()) / 24
	hours := int(timepassed.Hours()) % 24
	minutes := int(timepassed.Minutes()) % 60
	seconds := int(timepassed.Seconds()) % 60
	fmt.Printf("Time passed: %d days, %d hours, %d minutes, %d seconds\n", days, hours, minutes, seconds)
}