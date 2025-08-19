package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/joho/godotenv"
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

// * ---------- Middleware to check API Key ----------
func apiKeyMiddleware(apiKey string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-Key")
		if key != apiKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
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

// * ---------- Load API Key ----------
func loadAPIKey() string {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("ERROR: API_KEY is not set")
	}
	return apiKey
}

// * ---------- Main ----------
func main() {
	apiKey := loadAPIKey()
	http.Handle("/analyze", apiKeyMiddleware(apiKey, http.HandlerFunc(analyzeHandler)))
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
