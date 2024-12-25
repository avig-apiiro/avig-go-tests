package apis

import (
	"avig-go-tests/funcs"
	"encoding/json"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "name" query parameter from the request
	name := r.URL.Query().Get("name")

	// Mock user data
	user := funcs.User{}
	funcs.GetUser(name, &user)
	user.Cvv = "111"
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Serialize user to JSON and write to response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}

func Serve() {
	http.HandleFunc("/api2", UserHandler)
	user := funcs.User{}
	funcs.GetUser("test", &user)
}
