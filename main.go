package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		return 
	}
	
}

func handler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, " <h1>welcome to go....ramya <h2>")
	if err != nil {
		return 
	}
}