package main

import (
	"fmt"
	"log"
	"todoApp/pkg/database/postgres"
)

func main() {
	fmt.Println("test")
	db, err :=  postgres.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}