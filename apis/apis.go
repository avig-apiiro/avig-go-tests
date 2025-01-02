package apis

import (
	"avig-go-tests/funcs"
	"encoding/json"
	. "net/http"
)

type UserApi struct{}

func (userApi UserApi) UserHandler(w ResponseWriter, r *Request) {
	// Get the "name" query parameter from the request
	name := r.URL.Query().Get("name")

	// Mock user data

	user := funcs.GetUser(name)
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Serialize user to JSON and write to response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		Error(w, "Failed to encode user", StatusInternalServerError)
		return
	}
}

func (userApi UserApi) Serve() {
	HandleFunc("/api2", userApi.UserHandler)
}
