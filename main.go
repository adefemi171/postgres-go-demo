package main

import (
	"flag"
	"fmt"
	"log"

	config "github.com/adefemi171/postgres-go/pkg/config"
	route "github.com/adefemi171/postgres-go/pkg/route"

	"github.com/gin-gonic/gin"
)

func main() {

	var host = flag.String("addr", ":8001", "http service address of the application.")
	flag.Parse()

	// Calling DB
	config.Connection()

	r := gin.Default()
	route.UserRoutes(r)

	// http.Handle("/", r)

	log.Println("Starting web server on ", *host)
	fmt.Println("Open browser and redirect to http://localhost:8001")

	// log.Fatal(http.ListenAndServe("localhost:8001", r))
	log.Fatal(r.Run("localhost:8001"))
}
