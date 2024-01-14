package main

import(
	"net/http"
	"strconv"
	"fmt"
	"html/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	templateFiles := []string {
		"./ui/html/base.tmpl.html",
		"./ui/html/components/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	template, err := template.ParseFiles(templateFiles...)

	if err != nil {
		app.errorLogger.Print(err.Error())
		app.serverError(w, err)
		return
	}

	err = template.ExecuteTemplate(w,"base", nil)
	if err != nil {
		app.errorLogger.Print(err.Error())
		app.serverError(w, err)
		return
	}

}

func (app *application)snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.errorLogger.Print(err.Error())
		app.notFound(w)
		return
	}
	// w.Write([]byte("Snippet view"))
	fmt.Fprintf(w, "To view snippet with id: %d - ", id)
}

func (app *application)createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")


		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		// or
		app.clientError(w, http.StatusMethodNotAllowed)
		
		return
	}

	w.Write([]byte("To create new snippet"))

}