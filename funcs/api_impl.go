package funcs

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetOtherHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Other Hello called!\n")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "name" query parameter from the request
	name := r.URL.Query().Get("name")

	// Mock user data
	user := GetUser(name)

	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Serialize user to JSON and write to response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}
