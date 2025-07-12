package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dibyendu-dev/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB api")
	r := router.Router()
	fmt.Println("server is getting started")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("listining at port 4000 ...")

}
