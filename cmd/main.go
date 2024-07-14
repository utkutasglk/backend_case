package main

import (
	"backend_case/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router :=mux.NewRouter()
	routes.CreateRouters(router)
	http.Handle("/", router)
	log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000",router))
}