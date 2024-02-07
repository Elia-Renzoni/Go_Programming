/**
*	@author Elia Renzoni
*	@date 07/02/2024
*	@brief web server with no middleware
**/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CityAPI struct {
	Name string
	Area uint64
}

func mainLogic2(w http.ResponseWriter, req *http.Request) {
	var cityInfo CityAPI
	if req.Method == "POST" {
		decoder := json.NewDecoder(req.Body)              // legge il body della richiesta
		if err := decoder.Decode(&cityInfo); err != nil { // scrive il contenuto del json nel luogo indicato come parametro
			panic(err)
		}
		defer req.Body.Close()
		fmt.Printf("Name : %s, Area : %d\n", cityInfo.Name, cityInfo.Area)
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("201 - create new element"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/", mainLogic2)
	http.ListenAndServe(":8080", nil)
}
