package main

import (
	"fmt"
	"net/url"
)

 const myurl string = "http://ddasdev.com:3000/learn?coursename=golang&uid=ddas4548"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	result , _ :=url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.Hostname())
	// fmt.Println(result.RawQuery)

	qparms := result.Query()
	fmt.Printf("query params is :%T\n",qparms)
	fmt.Println(qparms["coursename"])

	for _,val := range qparms{
		fmt.Println("params is",val)
	}

	partsofurl := &url.URL{
		Scheme: "https",
		Host: "ddas.dev",
		Path: "create",
		RawPath: "user=ddas",
	}

	anotherURL := partsofurl.String()
	fmt.Println(anotherURL)

}