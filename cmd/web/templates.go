package main

import (
	"io/fs"
	"path/filepath"
	"text/template"

	"github.com/levisthors/snippetbox/internal/models"
	"github.com/levisthors/snippetbox/ui"
)

type templateData struct {
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.gohtml",
			"html/partials/*.gohtml",
			page,
		}

		ts, err := template.New(name).ParseFS(ui.Files,
			patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
