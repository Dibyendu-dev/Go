package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in go")

	mych := make(chan int ,1)
	wg := &sync.WaitGroup{}
	// mych <- 5
	// fmt.Println(<-mych)
	wg.Add(2)
	go func(ch chan int , wg *sync.WaitGroup){
		fmt.Println(<-mych)
		wg.Done()
	}(mych,wg)

	go func(ch chan int , wg *sync.WaitGroup){
		mych <- 5
		mych <- 6
		close(mych)
		wg.Done()
	}(mych,wg)
	wg.Wait()
}
