package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8087", nil)
	if err != nil {
		return 
	}
	
}

func handler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, " welcome to go....ramya")
	if err != nil {
		return 
	}
}