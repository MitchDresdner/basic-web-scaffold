package main

import (
	"basic-web-scaffold/internal/config"
	"basic-web-scaffold/internal/handlers"
	"basic-web-scaffold/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"runtime"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager
var preferenceMap map[string]string

const appVersion = "1.0.0"

// main is the main function
func main() {
	// change this to true when in production
	app.InProduction = false

	insecurePort, err := setupApp()
	if err != nil {
		log.Fatal("main::setupApp", err)
	}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	banner(app.Identifier)

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// create http server
	srv := &http.Server{
		Addr:              *insecurePort,
		Handler:           routes(&app),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Printf("HTTP server starting on port %s....", *insecurePort)

	// start the server
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("main::ListenAndServe", err)
	}
}

func banner(name string) {
	// print banner
	log.Printf("******************************************")
	log.Printf("** \033[31m%s\033[0m v%s built in %s", name, appVersion, runtime.Version())
	log.Printf("**----------------------------------------")
	log.Printf("** Running with %d Processors", runtime.NumCPU())
	log.Printf("** Running on %s", runtime.GOOS)
	log.Printf("******************************************")
}
