package config

import (
	"fmt"

	"github.com/sufian22/goupon/db"
	"github.com/sufian22/goupon/db/postgres"
)

func ConfigureDatabase(dbConfig DBConfig) (db.ORM, error) {
	var db db.ORM

	url := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port,
		dbConfig.User, dbConfig.Password, dbConfig.Name)

	var err error
	dbDriver := dbConfig.Driver
	switch dbDriver {
	case "postgres":
		db, err = postgres.NewDB(url)
	default:
		return db, fmt.Errorf("Not supported SQL driver %s", dbDriver)
	}
	if err != nil {
		return db, fmt.Errorf("Could not connect to database: %v", err)
	}

	return db, nil
}
