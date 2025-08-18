package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unicode"
)

// * -- Request Struct
type Request struct {
	Sentence string `json:"sentence"`
}

// * -- Response Struct
type Response struct {
	Words      int `json:"words"`
	Vowels     int `json:"vowels"`
	Consonants int `json:"consonants"`
}

// * ---------- Analyze the sentence ----------
func analyze(sentence string) Response {
	words := len(strings.Fields(sentence))
	vowels, consonants := 0, 0

	for _, r := range strings.ToLower(sentence) {
		if unicode.IsLetter(r) {
			if strings.ContainsRune("aeiou", r) {
				vowels++
			} else {
				consonants++
			}
		}
	}

	return Response{Words: words, Vowels: vowels, Consonants: consonants}
}

// * ---------- Handles The Request ----------
func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp := analyze(req.Sentence)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// * -- Main
func main() {
	http.HandleFunc("/analyze", analyzeHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
