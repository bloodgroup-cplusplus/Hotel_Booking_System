package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

//Renders Template using html template

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	// get the template (requested template from cache)
	// render the template
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	// we are going to be returning a map of pointer string to template.tempalte
	// we will create a variable

	//myCache := make(map[string] *template.Template)

	myCache := map[string]*template.Template{}
	// we want to populate the cache with every available template
	// all templates are on the root level of template

	// get all of the files named *.page.tmpl from ./templates

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl

	for _, page := range pages {
		// get the actual file name
		// we want just the name of the template

		name := filepath.Base(page) // last element of the path
		// try to parse the file
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// look for any layouts in the directory

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		}
	}

}

/* this is a naive way of doing it all the commented section by keeping the map and adding the name of the template everytime we use it

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
*/
