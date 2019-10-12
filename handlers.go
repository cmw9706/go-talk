package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func getPeople(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	people, err := db.GetPeople()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(people)
	return
}

func getPerson(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		http.Error(writer, fmt.Sprintf("%s is not a valid beer ID, it must be a number", params.ByName("id")), http.StatusBadRequest)
		return
	}

	result, err := db.GetPerson(Person{ID: id})

	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(result)
	return
}
