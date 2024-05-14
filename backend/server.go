package main

import (
	"fmt"
	"net/http"

	"github.com/fearsd/course_work2024/backend/src/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Handler)

	fmt.Println("Server is running on http://:3333")
	http.ListenAndServe(":3333", nil)
}

// TODO: функция алгоритма
// TODO: хендлер, который принимает запрос на генерацию перестановки по заданному массиву
// TODO: сделать интерактивный режим (придумать как объяснить это все по шагам)
