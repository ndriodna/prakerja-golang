package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	ID      int
	Title   string
	Content string
}

var data = []Article{
	{1, "Lorem", "Ipsum dolor set amet"},
	{2, "Lorem", "Ipsum dolor"},
}

func articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		res, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/articles", articles)
	fmt.Println("starting server at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
