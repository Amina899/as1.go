package main

import (
	"aitu/aitunews/pkg/models"
	"aitu/aitunews/pkg/models/mysql"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	news          *mysql.NewsModel
	templateCache map[string]*template.Template
	categories    []*models.Category
	contacts      []*models.Contact
}

func main() {
	dsn := flag.String("dsn", "new:amina@/aitunews?parseTime=true", "MySQL data source name")
	addr := flag.String("addr", ":4040", "HTTP network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		news:          &mysql.NewsModel{DB: db},
		templateCache: templateCache,
		categories: []*models.Category{
			{Name: "For Students"},
			{Name: "For Staff"},
			{Name: "For Applicants"},
			{Name: "For Researchers"},
		},
		contacts: []*models.Contact{
			{Name: "John Doe", Email: "john@example.com", Message: "Hello"},
			{Name: "Jane Doe", Email: "jane@example.com", Message: "Hi"},
			// Add more contacts as needed
		},
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

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
