package main

import (
	"basic-web-scaffold/internal/config"
	"flag"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

func setupApp() (*string, error) {
	// read flags
	domain := flag.String("domain", "localhost", "domain name (e.g. example.com)")
	identifier := flag.String("identifier", "scaffold", "unique identifier")
	insecurePort := flag.String("port", ":8080", "port to listen on")
	inProduction := flag.Bool("production", false, "application is in production")

	flag.Parse()

	// TODO rough in DB connect

	// session
	log.Printf("Initializing session manager....")
	session = scs.New()
	//session.Store = postgresstore.New(db.SQL)
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	//session.Cookie.Name = fmt.Sprintf("gbsession_id_%s", *identifier)
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = *inProduction

	// scaffold state config
	a := config.AppConfig{
		Session:      session,
		InProduction: *inProduction,
		Domain:       *domain,
		Version:      appVersion,
		Identifier:   *identifier,
	}

	app = a

	// add preferences to map
	preferenceMap = make(map[string]string)
	preferenceMap["version"] = appVersion

	return insecurePort, nil
}
