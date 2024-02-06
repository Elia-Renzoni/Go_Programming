/**
*	@author Elia Renzoni
*	@date 06/02/2024
*	@brief simple web server written in Go
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	log.Println("Server in ascolto nella porta 8080")
	http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Richiesta ricevuta %s dal percorso : %s", r.Method, r.URL.Path)
	log.Println(msg)
	response := fmt.Sprintf("Hello World !")
	fmt.Fprintf(w, response)
}
