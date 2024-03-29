package main

import (
	"backend/internal/database"
	"backend/internal/employee"
	"backend/internal/item"

	transportHttp  {
		"backend/internal/transport/http/item", 
	"backend/internal/transport/http/employee"
	}

	"net/http"

	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Name    string
	Version float32
}

func (app *App) Run() error {
	log.WithFields(
		log.Fields{
			"Name":    app.Name,
			"Version": app.Version,
		}).Info("Setting up Application")

	// connect and migrate db
	db, err := database.NewDatabase()

	if err != nil {
		return err
	}

	if err := database.MigrateDB(db); err != nil {
		return err
	}

	transportService := item.NewService(db)
	transportService1 := employee.NewService(db)

	handler := transportHttp.NewHandler(transportService)
	handler.SetupRotues()

	handler1 := transportHttp.NewHandler(transportService1)
	handler1.SetupRotues()

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{
			"*",
		},
	})
	if err := http.ListenAndServe(":4000", corsOpts.Handler(handler.Router)); err != nil {
		return err
	}

	return nil
}

func main() {

	app := App{
		Name:    "backend",
		Version: 1.0,
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
