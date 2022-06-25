package main

import (
	"backend/internal/database"
	"backend/internal/item"
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

	// connect and migrate db
	db, err := database.NewDatabase()

	if err != nil {
		return err
	}

	if err := database.MigrateDB(db); err != nil {
		return err
	}

	_ = item.NewService(db)

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
