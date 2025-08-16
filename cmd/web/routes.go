package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nowa75/bookings/pkg/config"
	"github.com/nowa75/bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Compress(5))
	mux.Use(middleware.Recoverer)
	//mux.Use(middleware.Logger)

	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fs := http.FileServer(http.Dir("./css"))
	mux.Handle("/css/*", http.StripPrefix("/css", fs))
	fsj := http.FileServer(http.Dir("./js"))
	mux.Handle("/js/*", http.StripPrefix("/js", fsj))

	return mux
}
