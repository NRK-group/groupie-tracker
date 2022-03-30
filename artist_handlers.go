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
	b := r.FormValue("submit")
	fmt.Print(b)
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
	for i := range Artists {
		resp2, err := http.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", Artists[i].ID))
		if err != nil {
			fmt.Println(err)
		}
		defer resp2.Body.Close()
		body2, _ := ioutil.ReadAll(resp2.Body)
		if err := json.Unmarshal(body2, &Artists[i].Concerts); err != nil {
			fmt.Println("err")
		}
	}
	idNum, err := strconv.Atoi(b)
	if err != nil {
		idNum = 1
	}
	data := Artists[idNum-1]
	i := map[string]interface{}{
		"Artists": Artists,
		"Data":    data,
	}
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, i)
}
