package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var db Storage
var router *httprouter.Router

func main() {
	var err error

	db, err = NewStorage(Memory)

	if err != nil {
		log.Fatal(err)
	}

	PopulateData()

	router = httprouter.New()

	router.GET("/people", getPeople)
	router.GET("/people/:id", getPerson)

	fmt.Println("People api listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
