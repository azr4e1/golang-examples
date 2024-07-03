package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strings"
)

func init() {
	fsys := os.DirFS("pages")
	fs.WalkDir(fsys, ".", func(_ string, d fs.DirEntry, err error) error {
		if d.Type().IsRegular() {
			name := strings.Split(d.Name(), ".")[0]
			pages = append(pages, name)
		}
		return nil
	})
	pagesRE = regexp.MustCompile("(" + strings.Join(pages, "|") + ")")
}

var templates = template.Must(template.New("").Funcs(template.FuncMap{"replace": replace}).ParseFiles("tmpl/edit.html", "tmpl/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var pagesRE *regexp.Regexp

var pages = []string{}

type Page struct {
	Title string
	Body  []byte
}

func replace(body []byte) string {
	b := pagesRE.ReplaceAllFunc([]byte(body), func(pageName []byte) []byte {
		return []byte(fmt.Sprintf("<a href=\"/view/%s\">%s</a>", pageName, pageName))
	})

	return string(b)
}

func (p *Page) save() error {
	filename := "pages/" + p.Title + ".txt"
	if !slices.Contains(pages, p.Title) {
		pages = append(pages, p.Title)
		pagesRE = regexp.MustCompile("(" + strings.Join(pages, "|") + ")")
	}
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "pages/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, err := template.ParseFiles(tmpl + ".html")
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, err := getTitle(w, r)
		if err != nil {
			return
		}
		fn(w, r, title)
	}
}

func redirectRootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func main() {
	http.HandleFunc("/", redirectRootHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
