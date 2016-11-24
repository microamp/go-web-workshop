// With github.com/gorilla/mux!

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func helloName(w http.ResponseWriter, req *http.Request) {
	name := mux.Vars(req)["name"]
	fmt.Fprintf(w, "Hello, %s", name)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc(
		"/hello/{name}", // Variable extraction from path
		helloName,
	).Methods("GET") // GET method

	// handle all requests with the Gorilla router.
	http.Handle("/", r)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
