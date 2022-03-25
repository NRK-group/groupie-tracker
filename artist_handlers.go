package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

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
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &Artists); err != nil {
		fmt.Println("ERR")
	}

	for i := range Artists {
		resp2, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", Artists[i].ID))
		if err != nil {
			fmt.Println(err)
		}
		defer resp2.Body.Close()
		body2, err := ioutil.ReadAll(resp2.Body)
		if err := json.Unmarshal(body2, &Artists[i].Concerts); err != nil {
			fmt.Println("err")
		}
	}

	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, Artists)
}
