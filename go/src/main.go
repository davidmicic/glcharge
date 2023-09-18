package main

import (
	"fmt"
	c "glcharge/go/src/container"
	"glcharge/go/src/handlers"
	"glcharge/go/src/storage"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// postgresql://${process.env.POSTGRES_USER}:${process.env.POSTGRES_PASS}@${process.env.POSTGRES_HOST}:${process.env.POSTGRES_PORT}/${process.env.POSTGRES_DB}`;

	conn_str :=
		"postgresql://root:root@localhost:5432/glcharge?sslmode=disable"

	fmt.Printf("conn str: " + conn_str)

	dalPg := new(storage.DalDB)
	dalPg.InitDB(conn_str)

	container := c.GetContainer()
	container.SetStorage(dalPg)

	// s := &server{}
	// http.Handle("/", s)
	h := handlers.Makehttphandlers()
	e := http.ListenAndServe(":8080", h)
	log.Fatal(e)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
