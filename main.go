package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/todos", TodoIndex)
	router.GET("/todos/:todoId", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}