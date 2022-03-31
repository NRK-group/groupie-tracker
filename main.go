package main

import (
	"net/http"
)

func main() {
	// The "HandleFunc" method accepts a path and a handler function as arguments
	// The handler function has to have the appropriate signature (as described by the "handler" function below)
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/artists", getArtistHandler)
	http.HandleFunc("/", getArtistHandler)
	cssPath := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", cssPath)) // handling the CSS
	favico := http.FileServer(http.Dir("./favicon_io")) // handling the adress bar icon
	http.Handle("/favicon_io/", http.StripPrefix("/favicon_io/", favico)) // handling the icon for the adress bar
	// After defining the server, we "listen and serve" on port 8080
	// The second argument is the handler
	http.ListenAndServe(":8000", nil)
}
