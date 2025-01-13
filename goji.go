package main

import (
	"fmt"
	"goji.io"
	"goji.io/pat"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func gojimain() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	mux.Handle(pat.Post("/users"), usersHandler)

	users := SubMux()
	mux.Handle(pat.New("/users/*"), users)
	albums := SubMux()
	mux.Handle(pat.New("/albums/*"), albums)

	users.Handle(pat.Get("/:name"), renderProfile) // GET /users/carl
	albums.Handle(pat.Post("/"), newAlbum)         // POST /albums/

	http.ListenAndServe("localhost:8000", mux)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

}

func renderProfile(w http.ResponseWriter, r *http.Request) {

}
func newAlbum(w http.ResponseWriter, r *http.Request) {

}
