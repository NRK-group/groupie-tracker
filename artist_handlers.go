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
		fmt.Println("err")
	}

	//unpacking the relation api
	var Relation Relations
	if err := json.Unmarshal(GetAPI(url+"/relation"), &Relation); err != nil {
		fmt.Println("err")
	}

	// the code below gets the ID of the button that has been clicked and convert it to interger
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
}
