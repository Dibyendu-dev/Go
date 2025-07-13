package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup //pointers
var mut sync.Mutex  //pointer

func main() {
	websitelist := []string{
		"https://google.com",
		"https://youtube.com",
		"https://facebook.com",
		"https://amazon.com",
		"https://wikipedia.org",
		"https://twitter.com",
		"https://instagram.com",
		"https://linkedin.com",
		"https://netflix.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("oops sorry! server is down ... for hitting the %s\n",endpoint)
	} else {
		mut.Lock()
		signals = append(signals,endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for hitting the %s\n", res.StatusCode, endpoint)
	}

}
