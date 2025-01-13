package apis

import (
	"avig-go-tests/funcs"
	"encoding/json"
	. "net/http"
)

type UserApi struct{}
type MyUser = UserApi

func (MyUser UserApi) UserHandler(w ResponseWriter, r *Request) {
	// Get the "name" query parameter from the request
	name := r.URL.Query().Get("name")

	//user := funcs.GetUser(name)
	user := funcs.User{}
	user.UserName = name
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		Error(w, "Failed to  encode user", StatusInternalServerError)
		return
	}
}

func (userApi UserApi) Serve() {
	HandleFunc("/api4", userApi.UserHandler)
}
