package main

import (
	"backend/internal/database"
	transportHttp "backend/internal/transport/http"
	"net/http"

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

	_, err := database.NewDatabase()

	if err != nil {
		return err
	}

	handler := transportHttp.NewHandler()
	handler.SetupRotues()

	if err := http.ListenAndServe(":4000", handler.Router); err != nil {
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
