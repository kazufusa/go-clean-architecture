package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var value string

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, value)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	value = string(b)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getHandler).Methods("GET")
	r.HandleFunc("/", putHandler).Methods("PUT")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
