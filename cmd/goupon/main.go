package main

import (
	"flag"
	"log"

	"github.com/sufian22/goupon/api"
	"github.com/sufian22/goupon/api/handlers"
	"github.com/sufian22/goupon/config"
)

func main() {
	configFile := flag.String("config", "./config/config.json", "Configuration file")
	flag.Parse()

	var jsonConfig config.Config
	err := config.ReadConfigFile(*configFile, &jsonConfig)
	if err != nil {
		log.Fatal(err)
	}

	orm, err := config.ConfigureDatabase(jsonConfig.DBConfig)
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
