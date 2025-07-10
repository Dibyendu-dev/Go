package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON world ")
	// EncodedJson()
	DecodeJson()
}

func EncodedJson() {
	ddasCourses := []course{
		{"React.js", 5999, "yt.in", "pop121", []string{"webdev", "js", "frontend"}},
		{"Node.js", 5599, "yt.in", "pkp129", []string{"webdev", "js", "backend"}},
		{"Express.js", 2999, "yt.in", "oop121", nil},
	}

	//package this data into json format

	finalJson, err := json.MarshalIndent(ddasCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		
        {
                "coursename": "React.js",
                "price": 5999,
                "website": "yt.in",
                "tags": [
                        "webdev",
                        "js",
                        "frontend"
                ]
        }
	
	`)

	var myCourse course

  	checkValid:= json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n",myCourse)
	}else {
		fmt.Println("JSON is  NOT valid")

	}

	// data into key value  

	var myOCourse map[string]interface {}
		json.Unmarshal(jsonDataFromWeb, &myOCourse)
		fmt.Printf("%#v\n",myOCourse)
		
	for key, v := range myOCourse {
		fmt.Printf("key is %v and value is %v and type is %T \n ", key,v,v)
	}
}
