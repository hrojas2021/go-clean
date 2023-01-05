package db

import (
	"fmt"
	"log"

	"github.com/hugo.rojas/custom-api/conf"
	"github.com/jmoiron/sqlx"
)

const dbType = "postgres"

func InitDB(config *conf.Configuration) *sqlx.DB {
	dbConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DB.HOST,
		config.DB.USER,
		config.DB.PASSWORD,
		config.DB.NAME,
		config.DB.PORT,
		config.DB.SSL,
	)
	db, err := sqlx.Open(dbType, dbConnection)
	if err != nil {
		log.Fatalf("Failed to open DB via %s: %v", config.DB.URL, err)
	}
	err = db.Ping() // Err on docker?
	if err != nil {
		fmt.Println("ERROR")
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(4)

	return db
}
