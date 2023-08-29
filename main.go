package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/sachinchaudhary003/golangAuth/Router"
)

func main() {
	fmt.Println("Hello")
	r := router.Router()
	fmt.Println("Router setup ready ")
	log.Fatal(http.ListenAndServe(":8080", r))
}
