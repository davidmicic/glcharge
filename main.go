package main

import (
	"fmt"
	c "glcharge/container"
	"glcharge/handlers"
	"glcharge/storage"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// type server struct{}

// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "hello world"}`))
// }

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
