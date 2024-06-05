package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fearsd/course_work2024/backend/src/algos"
)

type RequestData struct {
	Arr       []int `json:"arr"`
	EdgeIndex int   `json:"edgeIndex"`
	LastIndex int   `json:"lastIndex"`
}

type ResponseData struct {
	Arr       []int `json:"arr"`
	EdgeIndex int   `json:"edgeIndex"`
	LastIndex int   `json:"lastIndex"`
	RandIndex int   `json:"randIndex"`
}

func setupResponse(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	fmt.Fprintf(w, "it works")
}

func HandleAlgo(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	arr, rand_index := algos.Algo(data.Arr, data.EdgeIndex, data.LastIndex)
	response := ResponseData{arr, data.EdgeIndex, data.LastIndex, rand_index}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	fmt.Println("Request processed successfully")
}
