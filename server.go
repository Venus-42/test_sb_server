package main

import (
	"fmt"
	"log"
	"net/http"
)

func myFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/hello", myFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
