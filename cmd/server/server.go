package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/weimpact/webother/config"
	"github.com/weimpact/webother/files"
	"github.com/weimpact/webother/ideas"
)

func server() *mux.Router {
	m := mux.NewRouter()
	db, err := DB(config.Database())
	if err != nil {
		panic(fmt.Errorf("couldn't initialize db: %v", err))
	}
	ideaService := ideas.Service{DB: db}
	fileService := files.Service{DB: db}
	m.HandleFunc("/ideas", ideas.SaveIdeaHandler(ideaService)).Methods(http.MethodPut)
	m.HandleFunc("/ideas", ideas.FetchIdeasHandler(ideaService)).Methods(http.MethodGet)
	m.HandleFunc("/files", files.Upload(fileService)).Methods(http.MethodPut)
	m.HandleFunc("/ping", pong).Methods(http.MethodGet)
	staticServer := http.FileServer(http.Dir(config.StoreLocation()))
	m.PathPrefix("/static").Handler(http.StripPrefix("/static", staticServer))
	return m
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"success": true}`))
}

func DB(cfg config.DB) (*sqlx.DB, error) {
	var err error
	db, err := sqlx.Open(cfg.Driver, cfg.URL())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.MaxConnLifetime())
	return db, nil
}
