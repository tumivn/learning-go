package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed")
		return
	}
	fmt.Fprintln(w, "I'm alive!")
	fmt.Fprintf(w, "environment %s\n", app.config.env)
	fmt.Fprintf(w, "version %s\n", "1.0.0")
}

func (app *application) getPostBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "GET Books")
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "POST Book")
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintln(w, "Method not allowed")
	return
}
