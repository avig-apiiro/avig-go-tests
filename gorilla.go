package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type apiEvents struct {
	auth any
}

func (a apiEvents) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a apiEvents) fetchConfiguration() func(http.ResponseWriter, *http.Request) {
	panic("implement me")
}

func (a apiEvents) updateConfiguration() func(http.ResponseWriter, *http.Request) {
	panic("implement me")
}

func gorillamain() {
	r := mux.NewRouter()
	var s, s2 = r.PathPrefix("/products").Subrouter(), r.PathPrefix("/products2").Subrouter()
	s.HandleFunc("/", ProductsHandler).Methods("GET", "POST")                                          // "/products/"
	s2.HandleFunc("/{key}", ProductHandler).Methods("GET").Queries("filter", "{filter}", "by", "{by}") // "/products/{key}/"
	r.Handle("/ids/{id}/events", &apiEvents{auth: "auth"}).Methods("GET")
	uc := &apiEvents{}
	r.HandleFunc("/configuration", uc.fetchConfiguration()).Methods("GET")
	r.HandleFunc("/configuration", uc.updateConfiguration()).Methods("POST")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
