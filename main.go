package main

import (
	"avig-go-tests/funcs"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const ENDPOINT_URI = "/endpoint2"
const ENDPOINT_URI_BASE = "/base"

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "homepage called\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello api called!\n")
}
func main() {
	fmt.Println("Hello world")
	//http.HandleFunc("/", getRoot)
	//http.HandleFunc("/endpoint", getHello)
	http.HandleFunc("/user", funcs.UserHandler)
	//http.HandleFunc(ENDPOINT_URI, getHello)
	//http.HandleFunc(ENDPOINT_URI_BASE+ENDPOINT_URI, funcs.GetOtherHello)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
