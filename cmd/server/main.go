package main

import (
	transportHttp "backend/internal/transport/http"
	log "github.com/sirupsen/logrus"
	"net/http"
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
