package main

import (
	"fmt"
	"html/template"
	"path/filepath"
	"time"

	"github.com/zakkbob/snippetbox/internal/models"
)

// A holding structure for data we want to pass to templates, since only one piece of dynamic data can be passed
type templateData struct {
	CurrentYear     int
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var funcs = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, fmt.Errorf("finding pages to cache: %w", err)
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles("./ui/html/pages/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts = ts.Funcs(funcs)

		ts, err = ts.ParseGlob("./ui/html/partials/*.tpml.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
