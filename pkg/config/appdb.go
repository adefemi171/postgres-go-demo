package config

import (
	"log"
	"os"

	controller "github.com/adefemi171/postgres-go/pkg/controller"

	"github.com/go-pg/pg/v9"
	
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = ""
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
	opts := &pg.Options{
		User:     "potus",
		Password: "***********",
		Addr:     ":5432",
		Database: "go_test",
	}
	// db := pg.Connect(&pg.Options{
	// 	User: "postgres",
	// 	Password: "db_password",
	// 	Addr: "localhost:5432",
	// 	Database: "db_dbname",
	// })

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Connection failed")
		os.Exit(100)
	}
	log.Printf("Connection Successful")
	controller.CreateUserTable(db)
	controller.InitiateDB(db)

	return db
}
