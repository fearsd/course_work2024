package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fearsd/course_work2024/backend/src/algos"
)

type RequestData struct {
	Arr       []int
	EdgeIndex int
	LastIndex int
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it works")
}

func HandleAlgo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	arr := algos.Algo(data.Arr, data.EdgeIndex, data.LastIndex)
	response := fmt.Sprintf("arr: %v; edgeIndex: %d; lastIndex: %d", arr, data.EdgeIndex, data.LastIndex)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": response})
}
