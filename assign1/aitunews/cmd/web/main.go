package main

import (
	"aitu/aitunews/pkg/models/mysql"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	thenews  *mysql.NewsModel
}

func main() {
	addr := flag.String("addr", ":4040", "HTTP network address")
	dsn := flag.String("dsn", "new:amina@/aitunews?parseTime=true", "MySQL data source name")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t ", log.Ldate|log.Ltime|log.Lshortfile)
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		thenews:  &mysql.NewsModel{DB: db},
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
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
