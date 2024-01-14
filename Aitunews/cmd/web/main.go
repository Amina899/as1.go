package main

import (
	"Aitunews.aitu/pkg/models/mysql"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	news           *mysql.NewsModel
	templateCache  map[string]*template.Template
	NewsByCategory map[string][]mysql.News
	News           []mysql.News
	Category       string
}

func main() {
	// Parse command-line flags for configuration
	dsn := flag.String("dsn", "web:pass@/Aitunews?parseTime=true", "MySQL data source name")
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	// Initialize loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Open database connection
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Initialize template cache
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Create application instance
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		news:          &mysql.NewsModel{DB: db},
		templateCache: templateCache,
	}

	// Create HTTP server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// Start the server
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
	mux := app.routes()

	// Start the server
	http.ListenAndServe(":8080", mux)

}

// newTemplateCache creates and returns a template cache.
func newTemplateCache(dir string) (map[string]*template.Template, error) {
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

// openDB opens a database connection and returns a *sql.DB.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
