package handler

import (
	"fmt"
	"html/template"
)

var templates = make(map[string]*template.Template, 5)

func LoadTemplates() {
	templateNames := [5]string{"welcome.html", "form.html", "thanks.html", "sorry.html", "list.html"}

	for index, name := range templateNames {
		t, err := template.ParseFiles("templates/layout.html", "templates/"+name)

		if err != nil {
			fmt.Println(fmt.Errorf("error encountered when loading templates: %v", err))
			panic(err)
		}

		templates[name] = t
		fmt.Println(fmt.Sprintf("Loading template %d - %s", index, name))
	}
}
