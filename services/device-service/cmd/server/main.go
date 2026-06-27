package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Device Service 🚀")
}

func main() {

	http.HandleFunc("/", homeHandler)

	fmt.Println("Device Service running on port 8080...")

	http.ListenAndServe(":8080", nil)
}