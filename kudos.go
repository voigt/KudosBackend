package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"KudosBackend/models"

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

			router := models.NewRouter()

			log.Println("Listening...")
			log.Fatal(http.ListenAndServe(":3000", router))

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
