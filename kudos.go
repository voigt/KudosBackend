package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"playground/KudosBackend/models"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	models.InitDB("./kudos.db")

	SelectMode(os.Args)

}

func SelectMode(args []string) {

	if len(args) > 1 {
		switch args[1] {
		case "serve":
			r := mux.NewRouter()
			r.HandleFunc("/", models.GetAllKudos).Methods("GET")
			r.HandleFunc("/{url:blog.christophvoigt.com/[a-z|A-Z|-]+}", models.GetKudos).Methods("GET")
			r.HandleFunc("/{url:blog.christophvoigt.com/[a-z|A-Z|-]+}", models.PostKudos).Methods("POST")
			http.Handle("/", r)

			log.Println("Listening...")
			http.ListenAndServe(":3000", nil)
		case "post":
			models.PostKudoCount(args[2])
		case "get":
			if len(args[1:]) > 1 {
				kudoCount := models.GetKudoCount(args[2])
				fmt.Println(kudoCount)
			} else {
				//models.GetAllKudos()
			}
		case "view":
			//models.GetAllKudos()
		case "reset":
			models.ResetDB()
		default:
			fmt.Println("Invalid argument.")
		}
	} else {
		fmt.Println("Available commands: serve, get, post, view, reset")
	}

}
