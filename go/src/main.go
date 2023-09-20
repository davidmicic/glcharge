package main

import (
	"fmt"
	c "glcharge/go/src/container"
	"glcharge/go/src/handlers"
	"glcharge/go/src/storage"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// @title           GLCharge
// @version         1.0
// @description     GLCharge API documentation.

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	godotenv.Load()
	conn_str := os.Getenv("CONN_STR")

	fmt.Printf("conn str: " + conn_str)

	dalPg := new(storage.DalDB)
	dalPg.InitDB(conn_str)
	container := c.GetContainer()
	container.SetStorage(dalPg)

	r := handlers.Makehttphandlers()
	e := r.Run()
	log.Fatal(e)
}
