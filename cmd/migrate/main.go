package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/io/database"
)

const (
	actionUp   = "up"
	actionDown = "down"
)

func main() {
	args := os.Args[1:]
	if args[0] != actionUp && args[0] != actionDown {
		log.Fatal("Invalid input. Unable to run migrations")
	}

	config := conf.LoadViperConfig()
	db := database.InitDB(config)

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatal("Error on the drive: ", err.Error())
	}

	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../../internal/migrations")
	migrationDir := filepath.Join("file://" + basePath)

	m, err := migrate.NewWithDatabaseInstance(
		migrationDir, "postgres", driver)

	if err != nil {
		log.Fatal("Error on creating the migrator: ", err.Error())
	}

	var migrationFactor = 1
	if args[0] == actionDown {
		migrationFactor = -1
	}

	if len(args) == 2 {
		steps, errConv := strconv.Atoi(args[1])
		if errConv != nil {
			log.Fatal("Error converting the steps: ", err.Error())
		}
		err = m.Steps(steps * migrationFactor)
	} else {
		if args[0] == actionUp {
			err = m.Up()
		}
		if args[0] == actionDown {
			err = m.Down()
		}
	}

	if err != nil {
		if err.Error() != "no change" {
			panic(err.Error())
		}
	}
}
