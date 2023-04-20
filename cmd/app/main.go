package main

import (
	"log"
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