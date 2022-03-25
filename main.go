package main

import (
	"fmt"
	"net/http"
)

func main() {
	// The "HandleFunc" method accepts a path and a handler function as arguments
	// The handler function has to have the appropriate signature (as described by the "handler" function below)
	http.HandleFunc("/", handler)
	http.HandleFunc("/artists/", getArtistHandler)

	// After defining the server, we "listen and serve" on port 8080
	// The second argument is the handler
	http.ListenAndServe(":8000", nil)
}

// Our handler function follows the function signature of a ResponseWriter and Request type as arguments
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
