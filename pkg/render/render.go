package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//Renders Template using html template

func RenderTemplateTEst(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	// we are going to need some kind of render variables
	// check to see if the parsed templates
	// the best data strcutre to use it is a map
	// declare two variables for individual templates
	//

	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache

	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	// execute the template

	err = tmpl.Execute(w, nil)
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the tempalte
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache (map)

	tc[t] = tmpl

	return nil

}
