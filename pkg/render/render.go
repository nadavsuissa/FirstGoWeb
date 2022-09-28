package render

import (
	"FirstGoWeb/pkg/config"
	"FirstGoWeb/pkg/models"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddDefaultData(templatedata *models.TemplateData) *models.TemplateData {
	return templatedata
}
func RenderTemplate(w http.ResponseWriter, gohtml string, templatedata *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[gohtml]
	if !ok {
		log.Fatal("Error")
	}
	buf := new(bytes.Buffer)

	templatedata = AddDefaultData(templatedata)
	_ = t.Execute(buf, templatedata)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error", err)
	}

}

// CreateTemplateCache Creates a Template Cache as a Map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}

		}
		myCache[name] = ts
	}
	return myCache, nil
}
