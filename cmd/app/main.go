package main

import (
	"log"
	"net/http"
	"todoApp/pkg/database"
	_ "github.com/lib/pq"
)

func main() {
	
	db, err := postgres.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.Close()

	
}