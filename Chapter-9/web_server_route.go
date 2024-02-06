/**
*	@author Elia Renzoni
*	@date 06/02/2024
*	@brief basic web server written in Golang with routing func.
**/

package main

import (
	"net/http"
	"time"
)

func handleRequest(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Hello World"))
}

func main() {
	// associo un API endpoint alla funzione da chiamare
	http.HandleFunc("/hello", handleRequest)

	server_info := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second, // session info
		WriteTimeout:   10 * time.Second, // session info
		MaxHeaderBytes: 1 << 20,
	}

	server_info.ListenAndServe()
}
