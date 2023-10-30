package main

import (
	"html/template"
	"net/http"
)

var responses = make([]*Rsvp, 0, 10)                   //init a new slice of pointers to Rsvp structs
var templates = make(map[string]*template.Template, 3) //init a new map of string keys to template pointers

func main() {
	loadTemplates()
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/list", listHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println(err)
	}
}

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

type formData struct {
	*Rsvp
	Errors []string
}

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}

	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			println("Loaded template ", index, name)
		} else {
			panic(err)
		}
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	templates["welcome"].Execute(w, nil)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	templates["list"].Execute(w, responses)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates["form"].Execute(w, formData{
			Rsvp:   &Rsvp{},
			Errors: []string{},
		})
	} else {
		r.ParseForm()
		rsvp := new(Rsvp)
		rsvp.Name = r.FormValue("name")
		rsvp.Email = r.FormValue("email")
		rsvp.Phone = r.FormValue("phone")
		rsvp.WillAttend = r.FormValue("willAttend") == "true"

		errors := make([]string, 0, 3)
		if rsvp.Name == "" {
			errors = append(errors, "Name is required")
		}
		if rsvp.Email == "" {
			errors = append(errors, "Email is required")
		}
		if rsvp.Phone == "" {
			errors = append(errors, "Phone is required")
		}

		if len(errors) == 0 {
			responses = append(responses, rsvp)
			templates["thanks"].Execute(w, rsvp)
		} else {
			data := formData{rsvp, errors}
			templates["form"].Execute(w, data)
		}
	}
}
