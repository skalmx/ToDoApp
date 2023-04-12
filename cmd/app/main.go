package main

import (
	_ "github.com/lib/pq"
	"log"
	"todoApp/pkg/database"
)

func main() {
	
	db, err := postgres.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("1232323")
	defer db.Close()
}