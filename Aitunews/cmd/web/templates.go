// In templates.go
package main

import (
	"Aitunews.aitu/pkg/models"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type templateData struct {
	NewsByCategory map[string][]*models.News
	News           []*models.News
	Category       string
}

func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Walk the directory and compile templates
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".page.tmpl") {
			name := strings.TrimSuffix(filepath.Base(path), ".page.tmpl")
			tmpl, err := template.New(name).ParseFiles(path)
			if err != nil {
				return err
			}
			cache[name] = tmpl
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return cache, nil
}
