package main

import (
	_ "Aitunews.aitu/pkg/models"
	"Aitunews.aitu/pkg/models/mysql"
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

type templateData struct {
	NewsByCategory map[string][]mysql.News
	News           []mysql.News
	Category       string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	newsByCategory := make(map[string][]mysql.News)

	for _, category := range []string{"For Students", "For Staff", "For Applicants", "For Researchers"} {
		news, err := app.news.ByCategory(category)
		if err != nil {
			app.serverError(w, err)
			return
		}

		newsByCategory[category] = news
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		NewsByCategory: newsByCategory,
	})
}

func (app *application) category(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	if category == "" {
		app.serverError(w, http.StatusBadRequest)
		return
	}

	news, err := app.news.ByCategory(category)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "category.page.tmpl", &templateData{
		News:     news,
		Category: category,
	})
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "contact.page.tmpl", nil)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, http.StatusText(http.StatusInternalServerError))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	if err != nil {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		app.errorLog.Output(2, trace)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	defaultData := &templateData{
		NewsByCategory: make(map[string][]mysql.News),
		News:           []mysql.News{},
		Category:       "",
	}

	for category, news := range td.NewsByCategory {
		defaultData.NewsByCategory[category] = news
	}

	return defaultData
}

// Initialize your dependencies and routes
func newApplication() *application {
	return &application{
		news:           mysql.NewNewsModel(nil), // Initialize your dependencies with an appropriate DB connection
		NewsByCategory: make(map[string][]mysql.News),
		News:           []mysql.News{},
		Category:       "",
	}
}
