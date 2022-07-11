package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rsvp struct {
	Name       string
	Email      string
	Phone      string
	WillAttend bool
}

type formData struct {
	*Rsvp
	Errors []string
}

var responses = make([]*Rsvp, 0, 10)

var templates = make(map[string]*template.Template, 5)

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}

	for index, name := range templateNames {
		t, err := template.ParseFiles("templates/layout.html", "templates/"+name+".html")

		if err != nil {
			fmt.Println(fmt.Errorf("error encountered when loading templates: %v", err))
			panic(err)
		}

		templates[name] = t
		fmt.Println(fmt.Sprintf("Loading template %d - %s", index, name))
	}
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["form"].Execute(writer, formData{
			Rsvp:   &Rsvp{},
			Errors: []string{},
		})

		return
	}

	if request.Method == http.MethodPost {
		request.ParseForm()

		response := Rsvp{
			Name:       request.Form["name"][0],
			Email:      request.Form["email"][0],
			Phone:      request.Form["phone"][0],
			WillAttend: request.Form["willattend"][0] == "true",
		}

		errors := []string{}

		if response.Name == "" {
			errors = append(errors, "PLease enter your name")
		}
		if response.Email == "" {
			errors = append(errors, "PLease enter your email")
		}
		if response.Phone == "" {
			errors = append(errors, "PLease enter your phone")
		}
		if len(errors) > 0 {
			templates["form"].Execute(writer, formData{
				Rsvp:   &response,
				Errors: errors,
			})

			return
		}

		responses = append(responses, &response)

		if response.WillAttend {
			templates["thanks"].Execute(writer, response.Name)
		} else {
			templates["sorry"].Execute(writer, response.Name)
		}
	}
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}

func handler() {
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/list", listHandler)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Printf("error encountered: %v", err)
	}
}

func main() {
	loadTemplates()
	handler()
}
