package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"second-app/cmd/internal/data"
	"strconv"
	"time"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     "1.0.0",
	}

	app.writeJson(w, data)
}

func (app *application) writeJson(w http.ResponseWriter, data any) {
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Error sending JSON", http.StatusInternalServerError)
		return
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (app *application) getPostBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		app.getBooks(w, r)
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

func (app *application) getPutDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getBook(w, r)
	case http.MethodPut:
		fmt.Fprintln(w, "PUT Book")
	case http.MethodDelete:
		fmt.Fprintln(w, "DELETE Book")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (app *application) getBookId(w http.ResponseWriter, r *http.Request) (int, bool) {
	idStr := r.URL.Path[len("/v1/books/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return 0, true
	}
	return id, false
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id, done := app.getBookId(w, r)
	if done {
		return
	}
	book := data.Book{
		ID:        id,
		Title:     "The Hitchhiker's Guide to the Galaxy",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Price:     9.99,
		Pages:     354,
		Rating:    5.0,
	}

	app.writeJson(w, book)
}

func (app *application) getBooks(w http.ResponseWriter, r *http.Request) {
	books := []data.Book{
		{
			ID:        1,
			Title:     "The Hitchhiker's Guide to the Galaxy",
			Author:    "Douglas Adams",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Price:     9.99,
			Pages:     354,
			Rating:    5.0,
			Version:   1,
		},
		{
			ID:        2,
			Title:     "Cloud Native Go",
			Author:    "M.-L. Reimer",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Price:     9.99,
			Pages:     264,
			Rating:    4.5,
			Version:   1,
		},
	}

	app.writeJson(w, books)
}
