package main

import (
	"log"
	"net/http"
	"playground/KudosBackend/models"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	models.InitDB("./kudos.db")

	r := mux.NewRouter()
	r.HandleFunc("/", models.Default).Methods("GET")
	r.HandleFunc("/{url:blog.christophvoigt.com/[a-z|A-Z|-]+}", models.GetKudos).Methods("GET")
	r.HandleFunc("/{url:blog.christophvoigt.com/[a-z|A-Z|-]+}", models.PostKudos).Methods("POST")
	//r.HandleFunc("/products", ProductsHandler)
	http.Handle("/", r)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)

	// SelectMode(os.Args)

}

/*
func SelectMode(args []string) {

	if len(args) > 1 {
		switch args[1] {
		case "post":
			models.PostKudos(args[2])
		case "get":
			if len(args[1:]) > 1 {
				kudoCount := models.GetKudos(args[2])
				fmt.Println(kudoCount)
			} else {
				models.GetAllKudos()
			}
		case "reset":
			models.ResetDB()
		default:
			fmt.Println("Invalid argument.")
		}
	} else {
		fmt.Println("Available commands: post, get, reset")
	}

}
*/
