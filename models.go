package main

// Creating models of API structs
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Concerts     Relations
}

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations`
}
