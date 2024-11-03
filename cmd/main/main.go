package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rohanfizz/go-bookstore/pkg/routes"
)

func main(){
	r := mux.NewRouter()	// Creating a custom mux
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":8080", r))
}