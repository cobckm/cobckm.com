package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	homeTmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/pages/home.html",
	))

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		homeTmpl.ExecuteTemplate(w, "base.html", nil)
	})

	srv := &http.Server{
		Addr:         ":3000",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
