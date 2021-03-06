package infrastructure

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
		log.Fatalf("Failed to ping DB via %s: %v", config.DB.URL, err)
	}
	// db.SetMaxIdleConns(0)
	// db.SetMaxOpenConns(2)

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB via %s: %v", config.DB.URL, err)
	}
	log.Println("Connected to DB")
	return db
}
