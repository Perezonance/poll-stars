package main

import (
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/alexperez/poll-stars/server"
	"github.com/alexperez/poll-stars/storage"

)

func main() {

	//Initialize the zap logger for structured logging
	l, err := getLogger(true)
	if err != nil {
		log.Fatalf("Logger failed to be initialized: %v", err)
	}

	//Initialize the connection to the db
	db, err := storage.NewSQL()
	if err != nil {
		log.Fatalf("Db Connection failed to be initialized: %v", err)
	}

	s := server.NewServer(db, l)

	http.HandleFunc("/", s.VoteHandler)
	http.ListenAndServe(":8080", nil)
}

func getLogger(prod bool) (*zap.Logger, error) {
	if prod {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}
