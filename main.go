package main

import (
	"avig-go-tests/apis"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const ENDPOINT_URI = "/endpoint"
const ENDPOINT_URI_BASE = "/base"

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "homepage called\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello api called    !\n")
}
func usermain() {
	fmt.Println("Hello  world")
	//http.HandleFunc("/", getRoot)
	//http.HandleFunc("/endpoint2", getHello)
	//http.HandleFunc(ENDPOINT_URI, getHello)
	api := apis.UserApi{}
	//http.HandleFunc(ENDPOINT_URI_BASE+ENDPOINT_URI, api.UserHandler)
	api.Serve()

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting  server: %s\n", err)
		os.Exit(1)
	}
}
