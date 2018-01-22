package database

import (
	"../config"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/HouzuoGuo/tiedot/dberr"
	"github.com/juju/errors"
)

func init() {
	databaseDir := config.GetString("databaseDir")

	appDb, err := db.OpenDB(databaseDir)
	if err != nil {
		panic(errors.Annotate(err, "Could not open database path!"))
	}

	initializeCollections(appDb)
}

func initializeCollections(appDb *db.DB) {
	cols := []string{"plugins", "presets", "samples", "kontakt", "reaktor", "uvi"}

	for _, col := range cols {
		if !appDb.ColExists(col) {
			appDb.Create(col)
		}
	}
}

func Scrub()
