package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sufian22/goupon/api"
	"github.com/sufian22/goupon/api/handlers"
	"github.com/sufian22/goupon/config"
	"github.com/sufian22/goupon/db"
	"github.com/sufian22/goupon/db/postgres"
)

func main() {
	configFile := flag.String("config", "./config/config.json", "Configuration file")
	flag.Parse()

	var jsonConfig config.Config
	err := config.ReadConfigFile(*configFile, &jsonConfig)
	if err != nil {
		log.Fatal(err)
	}

	orm, err := configureDatabase(jsonConfig.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	handlers.DB = orm

	srv, err := api.NewServer(jsonConfig.Port)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server listening on port %s", jsonConfig.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func configureDatabase(dbConfig config.DBConfig) (db.ORM, error) {
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
