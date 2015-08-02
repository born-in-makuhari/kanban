package main

import (
	"fmt"
	"github.com/hypebeast/gojistatic"
	"github.com/yosssi/ace"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
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

// Route configures routing. it is also used in test-code
func Route(m *web.Mux) {

	m.Get("/hello/:name", hello)
	m.Get("/", root)
}

func main() {
	// setup static
	goji.Use(gojistatic.Static("public",
		gojistatic.StaticOptions{SkipLogging: true}))

	// app routing
	Route(goji.DefaultMux)

	// start server
	goji.Serve()
}
