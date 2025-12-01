package templates

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var tmpl *template.Template

func LoadTemplates() error {
	tmplPath := filepath.Join("web", "templates", "*.gohtml")
	var err error
	tmpl, err = template.ParseGlob(tmplPath)
	return err
}

// Build renders the build template
func Build(w http.ResponseWriter, data interface{}) error {
	return tmpl.ExecuteTemplate(w, "build.gohtml", data)
}

// Home renders the home template
func Home(w http.ResponseWriter, data interface{}) error {
	return tmpl.ExecuteTemplate(w, "home.gohtml", data)
}
