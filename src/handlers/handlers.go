package handlers

import (
	"net/http"
	"regexp"
	"github.com/Kpovoc/simple-go-wiki/src/page"
	"github.com/Kpovoc/simple-go-wiki/src/tmpl"
)

func Init() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// Function Literal
func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2]) // The title is the second subexpression.
	}
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func viewHandler(w http.ResponseWriter, request *http.Request, title string) {
	p, err := page.LoadPage(title)
	if err != nil {
		http.Redirect(w, request, "/edit/"+title, http.StatusFound)
		return
	}
	tmpl.RenderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, request *http.Request, title string) {
	p, err := page.LoadPage(title)
	if err != nil {
		p = page.New(title, nil)
	}
	tmpl.RenderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, request *http.Request, title string) {
	body := request.FormValue("body")
	p := page.New(title, []byte(body))
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, request, "/view/"+title, http.StatusFound)
}