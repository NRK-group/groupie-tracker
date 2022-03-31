package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"
)

type Data struct {
	Index interface{}
}

func getArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Get response to the API

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// let's declare an Articles array
	// that we can then populate in our main function
	// to simulate a database
	var Artists []Artist
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &Artists); err != nil {
		fmt.Println("ERR")
	}

	//unpacking the relation api
	resp2, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println(err)
	}
	defer resp2.Body.Close()
	var Relation Relations
	body2, _ := ioutil.ReadAll(resp2.Body)
	if err := json.Unmarshal(body2, &Relation); err != nil {
		fmt.Println("err")
	}

	// the code below gets the ID of the button that has been clicked and convert it to interger
	b := r.FormValue("submit")
	idNum, err := strconv.Atoi(b)
	if err != nil {
		idNum = 1
	}

	// data will be all the information that will be in the inforamtion section
	data := Artists[idNum-1]
	data.Concerts = Relation.Index[idNum-1] // add the concert from relation api into the data

	info := map[string]interface{}{
		"Artists": Artists,
		"Data":    data,
	}

	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, info)
}
