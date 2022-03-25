//	main_test.go

package main

import (
	"testing"
)

func TestHandler(t *testing.T) {
	// Here, we form a new HTTP request. This is the request that's going to be passed to our handler.
	// The first argument is the method, the second argument is the route (which
	// we leave blank for now, and will get back to soon), and the third is the
	// request body, which we don't have in this case.
	// req, err := http.NewRequest("GET", "", nil)

	// // In case there is an error in forming the request, we fail and stop the test
	// if err != nil {
	// 	t.Fatal(err)
	// }

}
