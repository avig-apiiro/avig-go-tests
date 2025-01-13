package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	"github.com/zenazn/goji/graceful"
)

func zenaznmain() {
	app := web.New()

	app.Get("/", indexHandler)
	app.Get("/hello", indexHandler)

	// Admin router
	admin := web.New()
	app.Handle("/admin*", admin)

	admin.Use(middleware.SubRouter)
	admin.Get("/", adminIndexHandler)
	admin.Get("/secret", adminSecretHandler)
	admin.Post("/", newAdminThingHandler)

	// Start it up
	graceful.AddSignal(syscall.SIGINT, syscall.SIGTERM)
	err := graceful.ListenAndServe(":3333", app)
	if err != nil {
		log.Fatal(err)
	}
	graceful.Wait()
}
func indexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func adminIndexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func adminSecretHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func newAdminThingHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}
