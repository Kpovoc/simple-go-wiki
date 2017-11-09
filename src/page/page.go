package page

import "io/ioutil"

const dataRoot = "resources/data/"

// Page is an object used to store and build wiki-pages.
type Page struct {
	Title string
	Body  []byte
}

// New allocates a new wiki page with the given properties.
func New(title string, body []byte) *Page {
	return &Page {
		Title: title,
		Body: body,
	}
}

// LoadPage creates a Page from a .txt file in resources/data/ .
func LoadPage(title string) (*Page, error) {
	body, err := ioutil.ReadFile(genFilename(title))
	return New(title, body), err
}

// Save writes the Page to a file labeled {Page.Title}.txt in resources/data/ .
func (p *Page) Save() error {
	return ioutil.WriteFile(genFilename(p.Title), p.Body, 0600)
}

func genFilename(title string) string {
	return dataRoot + title + ".txt"
}