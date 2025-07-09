package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to API world")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code :", response.StatusCode)
	fmt.Println("Content Length is  :", response.ContentLength)

	var responseStr strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseStr.Write(content)
	fmt.Println("ByteCount is ", byteCount)
	fmt.Println(responseStr.String())
}

func PerformPostJsonRequest() {

	const myurl = "https://jsonplaceholder.typicode.com/posts"

	// fake JSON payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "lets learn react.js",
			"price" : 0,
			"platform" :"yt.com"

		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func PerformPostFormRequest() {

	const myurl = "https://jsonplaceholder.typicode.com/postsform" // pls create new, it's not exsist

	// formdata
	data := url.Values{}
	data.Add("firstname","dibyendu")
	data.Add("lastname","das")
	data.Add("email","ddas4548@gmail.com")

	response, err := http.PostForm(myurl,data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
