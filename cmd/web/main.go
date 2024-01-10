package main

import (
	"log"
	"net/http"
)





func main() {
	mux := http.NewServeMux() // this ensures MuxServer is scoped to my application

	fileServer := http.FileServer(http.Dir("./ui/static"))

// In Go’s servemux, longer URL patterns always take precedence
// thats the reason even though "/" is catch all route
// defined first, /snippet/view works fine without reordering

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView) 
	mux.HandleFunc("/snippet/create", createSnippet)
	
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Print("starting server on 4000")
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}
