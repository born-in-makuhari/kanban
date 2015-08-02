package main

import (
	"fmt"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
	"path/filepath"
)

// Ace Template Engine
// https://github.com/yosssi/ace/blob/master/documentation/getting-started.md

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func root(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("index", "", nil)
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{"Title": "Kanban"}
	err = tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func Route(m *web.Mux) {

	m.Get("/hello/:name", hello)
	m.Get("/", root)
}

func main() {
	// setup static
	static := web.New()
	publicPath, err := filepath.Abs("public")
	if err != nil {
		panic(err)
	}
	static.Get("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir(publicPath))))
	goji.Handle("/public/*", static)

	// app routing
	Route(goji.DefaultMux)

	// start server
	goji.Serve()
}
