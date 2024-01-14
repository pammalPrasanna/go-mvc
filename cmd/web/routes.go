package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux() // this ensures MuxServer is scoped to my application

	fileServer := http.FileServer(http.Dir("./ui/static"))

// In Go’s servemux, longer URL patterns always take precedence
// thats the reason even though "/" is catch all route
// defined first, /snippet/view works fine without reordering

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView) 
	mux.HandleFunc("/snippet/create", app.createSnippet)
	
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}