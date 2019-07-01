package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kazufusa/go-clean-architecture/domain/model"
)

var value model.Value

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, value.Value)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		// return http status code 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	value.Value = string(b)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getHandler).Methods("GET")
	r.HandleFunc("/", putHandler).Methods("PUT")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
