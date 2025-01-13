package main


import (
    "github.com/gorilla/mux"
    "net/http"
)

type
apiEvents struct {
    auth any
}

func (a apiEvents) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    //TODO implement me
    panic("implement me")
}

func main() {
    r := mux.NewRouter()
    var s, s2 = r.PathPrefix("/products").Subrouter(), r.PathPrefix("/products2").Subrouter()
    s.HandleFunc("/", ProductsHandler) // "/products/"
    .Methods("GET", "POST")
    s2.HandleFunc("/{key}", ProductHandler) // "/products/{key}/"
    .Methods("GET")
    .Queries("filter", "{filter}", "by", "{by}")
    r.Handle("/ids/{id}/events", &apiEvents{auth: auth}).Methods("GET")

    r.HandleFunc("/configuration", uc.fetchConfiguration()).Methods("GET")
    r.HandleFunc("/configuration", uc.updateConfiguration()).Methods("POST")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {

}
func ProductHandler(w http.ResponseWriter, r *http.Request) {

}