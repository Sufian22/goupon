package main

import (
	"flag"
	"log"

	"github.com/sufian22/goupon/api"
)

func main() {
	port := flag.String("port", "3000", "The desired port where server will run")
	flag.Parse()

	srv, err := api.NewServer(*port)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server listening on port %s", *port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
