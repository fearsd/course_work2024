package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server is running on http://:3333")
	http.ListenAndServe(":3333", nil)
}
