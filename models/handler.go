package models

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Default(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nothing here :)"))
}

func GetAllKudos() {

	rows, err := db.Query("SELECT id, url, kudos FROM kudos")
	checkErr(err)

	for rows.Next() {
		var id int
		var url string
		var kudos int
		err = rows.Scan(&id, &url, &kudos)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, url, kudos)
	}
	err = rows.Err()
	checkErr(err)

}

func GetKudoCount(url string) int {

	rows, err := db.Query("SELECT * FROM kudos WHERE kudos.url = (?)", url)
	checkErr(err)

	type Result struct {
		ID    int
		URL   string
		Kudos int
	}

	var result Result

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.URL, &result.Kudos)
		checkErr(err)
	}

	log.Printf("Requested URL %s (%d Kudos)", url, result.Kudos)

	return result.Kudos
}

func PostKudoCount(url string) int {
	doesUrlExist := GetKudoCount(url)

	var id int

	if doesUrlExist != 0 {
		log.Printf("%s is already known.", url)
		stmt, err := db.Prepare("UPDATE kudos set kudos = kudos + 1 WHERE kudos.url = (?)")
		checkErr(err)

		res, err := stmt.Exec(url)
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)

		log.Printf("Kudos for URL: %s with ID #%d given.", url, id)
	} else {
		stmt, err := db.Prepare("INSERT INTO kudos(url, kudos) values (?, 1)")
		checkErr(err)

		res, err := stmt.Exec(url)
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)

		log.Printf("Created URL: %s with ID #%d", url, id)
	}

	return id
}

func GetKudos(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := params["url"]

	kudos := GetKudoCount(url)
	kudosStr := strconv.Itoa(kudos)

	log.Printf("Requested URL %s (%d Kudos)", url, kudos)
	w.Write([]byte(kudosStr))
}

func PostKudos(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := params["url"]

	id := PostKudoCount(url)
	idStr := strconv.Itoa(id)

	w.Write([]byte("Kudos for URL: " + url + " with ID #" + idStr + " given."))

}
