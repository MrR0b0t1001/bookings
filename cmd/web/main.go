package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/MrR0b0t1001/bookings/pkg/config"
	"github.com/MrR0b0t1001/bookings/pkg/handlers"
	"github.com/MrR0b0t1001/bookings/pkg/render"
)

const portNum = ":8080"

var (
	app            config.AppConfig
	sessionManager *scs.SessionManager
)

func main() {
	// First time we create our templateCache

	app.InProduction = false // Still in developer mode
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Application could not start")
	}
	// We pass the templates we found into our app.TemplateCache to be used elsewhere
	app.TemplateCache = tc
	app.UseCache = false

	// Initializing session
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.SessionManager = sessionManager

	// repo is a pointer to Repository which contains information about our app
	// such as the TemplateCache var
	repo := handlers.NewRepo(&app)

	// We use the repo value to create our Repo value in the handlers.go file
	// Which gives  access to Handlers of our app config
	handlers.NewHandlers(repo)

	// we pass our app variable to be used by the render package
	// this enables the render package to have access to the app config
	render.NewTemplate(&app)

	serve := &http.Server{
		Addr:    portNum,
		Handler: routing(&app),
	}
	fmt.Println(fmt.Sprintf("Starting application on %s", portNum))
	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
