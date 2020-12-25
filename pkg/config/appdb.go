package config

import (
	"log"
	"os"

	controller "github.com/adefemi171/postgres-go/pkg/controller"

	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"
	
)




// create DB connection

// func Connection() {
// 	db := pg.Connect(&pg.Options{
// 		User: "postgres",
// 		Password: "db_password",
// 		Addr: "localhost:5432",
// 		Database: "db_dbname",
// 	})

// 	if db == nil {
// 		log.Printf("Connection failed")
// 		os.Exit(100)
// 	}
// 	log.Printf("Connection Successful")

// 	// return db
// }


func Connection() *pg.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	opts := &pg.Options{
		User:     os.Getenv("User"),
		Password: os.Getenv("Password"),
		Addr:     os.Getenv("Addr"),
		Database: os.Getenv("Database"),
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Oops Connection failed")
		os.Exit(100)
	}
	log.Printf("Yay, Connection Successful")
	controller.CreateUserTable(db)
	controller.InitiateDB(db)

	return db
}
