package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = make(map[string]*template.Template, 5)

func LoadTemplates() {
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

func Handler() {
	http.HandleFunc("/", welcomeHandler)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Printf("error encountered: %v", err)
	}
}
