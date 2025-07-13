package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race conditions")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("1st R")
		mut.Lock()
		score = append(score , 1)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	
	go func(wg *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("2nd R")
		mut.Lock()
		score = append(score , 2)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("3rd R")
		mut.Lock()
		score = append(score , 3)
		mut.Unlock()
		wg.Done()
	}(wg,mut)

	wg.Wait()
	fmt.Println("score is",score)
}
