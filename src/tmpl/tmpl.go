package tmpl

import (
	"net/http"
	"html/template"
	"github.com/Kpovoc/simple-go-wiki/src/page"
)

const tmplRoot = "resources/tmpl/"

// template.Must is a convenience wrapper that panics when passed a non-nil error value, and otherwise returns the
// *Template unaltered. A panic is appropriate here; if the templates can't be loaded the only sensible thing to do is
// exit the program.
var templates = template.Must(template.ParseFiles(tmplRoot + "edit.html", tmplRoot + "view.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}