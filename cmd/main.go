package main

import (
	"backend_case/pkg/routes"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	router :=mux.NewRouter()
	routes.CreateGetTeamRouters(router)
	http.Handle("/", router)
	log.Println("Listenin on port 8000")
	log.Fatal(http.ListenAndServe(":8000",router))
}