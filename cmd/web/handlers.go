package main

import(
	"net/http"
	"strconv"
	"fmt"
	"html/template"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	templateFiles := []string {
		"./ui/html/base.tmpl.html",
		"./ui/html/components/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	template, err := template.ParseFiles(templateFiles...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = template.ExecuteTemplate(w,"base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Snippet view"))
	fmt.Fprintf(w, "To view snippet with id: %d - ", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")


		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		// or
		http.Error(w, "Method not allowed", 405)
		
		return
	}

	w.Write([]byte("To create new snippet"))

}