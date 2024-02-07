/**
*	@author Elia Renzoni
*	@date 07/02/2024
*	@brief Gorilla Mux route
*	Gorilla Mux is a Go pakcage written to parse and set custom routes
**/

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func articleHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Category : %v\n", vars["category"])
	fmt.Fprintf(w, "ID : %v", vars["id"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{category}/{id:[0-9]+}", articleHandler)
	serverInfo := &http.Server{
		Handler:        r,
		Addr:           ":8080",
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	serverInfo.ListenAndServe()
}
