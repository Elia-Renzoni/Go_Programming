/**
*	@author Elia Renzoni
*	@date 08/02/2024
*	@brief Query Based Web Server
*
**/

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	queryID, queryCategory string = "id", "category"
)

func handleQuery(w http.ResponseWriter, r *http.Request) {
	queryContent := r.URL.Query()
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Query 1 : %s", queryContent[queryID][0])
	fmt.Fprintf(w, "Query 2 : %s", queryContent[queryCategory][0])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles", handleQuery)
	r.Queries(queryID, queryCategory)

	http.ListenAndServe(":8080", nil)
}
