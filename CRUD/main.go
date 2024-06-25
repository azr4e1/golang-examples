package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Memo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var memos []Memo
var idCounter int

func createMemo(w http.ResponseWriter, r *http.Request) {
	var memo Memo
	json.NewDecoder(r.Body).Decode(&memo)
	idCounter++
	memo.ID = idCounter
	memos = append(memos, memo)
	json.NewEncoder(w).Encode(memo)
}

func retrieveMemos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(memos)
}

func retrieveMemo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		panic(err)
	}

	for _, memo := range memos {
		if memo.ID == id {
			json.NewEncoder(w).Encode(memo)
			return
		}
	}

	http.Error(w, "Memo not found", http.StatusNotFound)
}

func updateMemo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		panic(err)
	}

	for i, memo := range memos {
		if memo.ID == id {
			json.NewDecoder(r.Body).Decode(&memo)
			memo.ID = id
			memos[i] = memo
			json.NewEncoder(w).Encode(memo)
			return
		}
	}

	http.Error(w, "Memo not found", http.StatusNotFound)
}

func deleteMemo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		panic(err)
	}

	for i, memo := range memos {
		if memo.ID == id {
			memos = append(memos[:i], memos[i+1:]...)
			json.NewEncoder(w).Encode("The memo was deleted successfully!")
			return
		}
	}

	http.Error(w, "Memo not found", http.StatusNotFound)
}

func initializeRouter() {
	router := http.NewServeMux()

	router.HandleFunc("POST /memos", createMemo)
	router.HandleFunc("GET /memos", retrieveMemos)
	router.HandleFunc("GET /memos/{id}", retrieveMemo)
	router.HandleFunc("PUT /memos/{id}", updateMemo)
	router.HandleFunc("DELETE /memos/{id}", deleteMemo)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	memos = append(memos, Memo{ID: 1, Title: "First memo", Content: "Hello World"})
	idCounter = 1
	initializeRouter()
}
