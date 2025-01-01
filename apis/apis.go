package apis

import (
	"avig-go-tests/funcs"
	"encoding/json"
	"net/http"
)

type UserApi struct{}

func (userApi UserApi) UserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "name" query parameter from the request
	name := r.URL.Query().Get("name")

	// Mock user data

	user := funcs.GetUser(name)
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Serialize user to JSON and write to response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}

func (userApi UserApi) Serve() {
	http.HandleFunc("/api2", userApi.UserHandler)
}
