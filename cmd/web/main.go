package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)


type application struct{
	infoLogger *log.Logger
	errorLogger *log.Logger
}


func main() {

	addr := flag.String("addr", "localhost:4000", "HTTP network address")
	flag.Parse()

	app := new(application)
	app.errorLogger = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)
	app.infoLogger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)


// Initialize a new http.Server struct. We set the Addr and Handler
// fields so
// that the server uses the same network address and routes as before,
// and set
// the ErrorLog field so that the server now uses the custom errorLog
// logger in
// the event of any problems.

	server := &http.Server{ // To make the mux server to use custom error logging
		Addr: *addr,
		Handler: app.routes(),
		ErrorLog: app.errorLogger, 
	}

	app.infoLogger.Print("starting server on "+*addr)
	err := server.ListenAndServe()
	app.errorLogger.Fatal(err)
}
