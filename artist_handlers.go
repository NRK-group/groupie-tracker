package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"
)

//GetAPI receive an api url and return the api into []byte.
//This function will get the url that you give, read it and return it as a []byte
func GetAPI(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	return body
}

func getArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Status not found", http.StatusNotFound)
		return
	}
	url := "https://groupietrackers.herokuapp.com/api"
	// let's declare an Articles array
	// that we can then populate in our main function
	// to simulate a database
	var Artists []Artist
	if err := json.Unmarshal(GetAPI(url+"/artists"), &Artists); err != nil {
		http.Error(w, "500 Internal error", http.StatusInternalServerError)
		return
	}

	//unpacking the relation api
	var Relation Relations
	if err := json.Unmarshal(GetAPI(url+"/relation"), &Relation); err != nil {
		http.Error(w, "500 Internal error", http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case "GET":
		// data will be all the information that will be in the inforamtion section
		data := Artists[0]
		data.Concerts = Relation.Index[0] // add the concert from relation api into the data struct

		info := map[string]interface{}{
			"Artists": Artists,
			"Data":    data,
		}

		t, _ := template.ParseFiles("static/index.html")
		t.Execute(w, info)
	case "POST":
		id := r.FormValue("artist")
		idNum, err := strconv.Atoi(id)
		if err != nil {
			idNum = 1
		}

		// data will be all the information that will be in the inforamtion section
		data := Artists[idNum-1]
		data.Concerts = Relation.Index[idNum-1] // add the concert from relation api into the data struct

		info := map[string]interface{}{
			"Artists": Artists,
			"Data":    data,
		}

		t, _ := template.ParseFiles("static/index.html")
		t.Execute(w, info)
	default:
		http.Error(w, "Bad request", http.StatusBadRequest) // if the request method is not GET or POST
	}

}
