/**
*	@author Elia Renzoni
*	@date 06/02/2024
*	@brief web server with a custom multiplexer that switch the API endpoints written in Go
**/

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type CustomServerMultiplexer struct {
	// empty
}

// implementation of a custom multiplexer
func (myMux *CustomServerMultiplexer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		basicRandom(w, req)
	} else {
		http.NotFound(w, req)
	}
}

func basicRandom(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	fmt.Fprintf(w, "Numero Random Basico : %d", rand.Intn(10))
}

func main() {
	customMux := &CustomServerMultiplexer{}
	http.ListenAndServe(":8080", customMux)
}
