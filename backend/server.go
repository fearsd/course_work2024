package main

import (
	"fmt"
	"net/http"

	"github.com/fearsd/course_work2024/backend/src/handlers"
)

func main() {
	http.HandleFunc("/algo", handlers.HandleAlgo)
	http.HandleFunc("/", handlers.HandleIndex)
	fmt.Println("Server is running on http://:3333")
	http.ListenAndServe(":3333", nil)
}
