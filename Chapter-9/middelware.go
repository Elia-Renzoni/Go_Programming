/**
*	@author Elia Renzoni
*	@date 07/02/2024
*	@brief Basic Middleware Implementation
*
**/

package main

import (
	"fmt"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("Middleware ON...")
		handler.ServeHTTP(w, req)
		fmt.Printf("Middleare OFF...")
	})
}

func mainLogic(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Sto eseguendo l'handler")
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "OK")
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	http.ListenAndServe(":8080", nil)
}
