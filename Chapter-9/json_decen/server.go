/**
*	@author Elia Renzoni
*	@date 28/02/2024
*	@brief Simple RESTful API
**/

package main

import (
	"encoding/json"
	"net/http"
)

const (
	_get, post string = "GET", "POST"
)

type Sum struct {
	FirstNumb  int `json:"first_number"`
	SecondNumb int `json:"second_number"`
}

func sum(w http.ResponseWriter, r *http.Request) {
	var sum Sum = Sum{}
	defer r.Body.Close()
	if r.Method == post {
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&sum)
		effective_sum := sum.FirstNumb + sum.SecondNumb
		w.WriteHeader(http.StatusAccepted)
		encoder := json.NewEncoder(w)
		encoder.Encode(effective_sum)

		//fmt.Fprintf(w, "Sum : %d", sum.FirstNumb+sum.SecondNumb)
	}
}

func main() {
	http.HandleFunc("/makesum", sum)

	http.ListenAndServe(":8080", nil)
}
