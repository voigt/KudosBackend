package main

import (
	"fmt"
	"os"
	"playground/KudosBackend/models"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	models.InitDB("./kudos.db")
	selectMode(os.Args)

}

func selectMode(args []string) {

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
