package models

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Website struct {
	Id    int    `json:"id"`
	Url   string `json:"url"`
	Kudos int    `json:"kudos"`
}

type Websites []Website

func GetAllKudos(w http.ResponseWriter, r *http.Request) {

	var website Website
	var websites Websites

	rows, err := db.Query("SELECT id, url, kudos FROM kudos")
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&website.Id, &website.Url, &website.Kudos)
		checkErr(err)

		websites = append(websites, website)

	}
	json.NewEncoder(w).Encode(websites)

	err = rows.Err()
	checkErr(err)

}

func GetKudoCount(url string) Website {

	rows, err := db.Query("SELECT * FROM kudos WHERE kudos.url = (?)", url)
	checkErr(err)

	var website Website

	for rows.Next() {
		err = rows.Scan(&website.Id, &website.Url, &website.Kudos)
		checkErr(err)
	}

	log.Printf("Query URL %s for Kudos... %d Kudos.", url, website.Kudos)

	return website

}

func PostKudoCount(url string) Website {

	website := GetKudoCount(url)

	if website.Id != 0 {
		log.Printf("%s is already known.", url)
		_, err := db.Exec("UPDATE kudos set kudos = kudos + 1 WHERE kudos.url = (?)", url)
		//stmt, err := db.Prepare("UPDATE kudos set kudos = kudos + 1 WHERE kudos.url = (?)")
		checkErr(err)

		website := GetKudoCount(url)

		log.Printf("Kudos for URL: %s with ID #%d given.", website.Url, website.Id)

		return website
	} else {
		website.Url = url
		website.Kudos = 1

		stmt, err := db.Prepare("INSERT INTO kudos(url, kudos) values (?, ?)")
		checkErr(err)

		res, err := stmt.Exec(website.Url, website.Kudos)
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)
		website.Id = int(id)

		log.Printf("Created URL: %s with ID #%d", url, website.Id)
		return website
	}

}

func GetKudos(w http.ResponseWriter, r *http.Request) {

	var website Website

	params := mux.Vars(r)
	url := params["url"]

	website = GetKudoCount(url)

	json.NewEncoder(w).Encode(website)

}

func PostKudos(w http.ResponseWriter, r *http.Request) {

	var website Website

	params := mux.Vars(r)
	website.Url = params["url"]

	website = PostKudoCount(website.Url)

	//w.Write([]byte("Kudos for URL: " + website.Url + " with ID #" + id + " given."))
	json.NewEncoder(w).Encode(website)

}
