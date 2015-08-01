package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func root(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello goji!")
}

func Route(m *web.Mux) {
	m.Get("/hello/:name", hello)
	m.Get("/", root)
}

func main() {
	Route(goji.DefaultMux)
	goji.Serve()
}
