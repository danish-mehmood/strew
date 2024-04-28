package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/danish-mehmood/strew/db"
)

var (
	db_path = flag.String("db-location", "", "path to the db")
	http_address = flag.String("http-addr" , ":8080" , "the server address")
)

func parse_flags() {
	flag.Parse()
	if *db_path == "" {
		log.Fatal("must pass the db path")
	}
}

func main() {
	parse_flags()

	database , close ,err := db.NewDatabase(*db_path)
    if err!=nil{
		log.Fatal(err)
	}
	defer close()

	fmt.Println(database)

	http.HandleFunc("/get" , func(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w , "accessing the get function ")})

	http.HandleFunc("/set" , func(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w , "accessing the set function ")})

	log.Fatal(http.ListenAndServe(*http_address , nil))

}


