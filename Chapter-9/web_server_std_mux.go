/**
*	@author Elia Renzoni
*	@date 06/02/2024
*	@brief web server that can handle multiple API endpoint
*
**/

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func getRandomInt(w http.ResponseWriter, req *http.Request) {
	rand.Seed(time.Now().Unix())
	fmt.Fprintf(w, "Intero Random : %d", rand.Intn(100))
}

func getRandomFloat(w http.ResponseWriter, req *http.Request) {
	rand.Seed(time.Now().Unix())
	fmt.Fprintf(w, "Float Random : %f", rand.Float64())
}

func getRandomFloat32(w http.ResponseWriter, req *http.Request) {
	rand.Seed(time.Now().Unix())
	fmt.Fprintf(w, "Float32 Random : %f", rand.Float32())
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/int_random", getRandomInt)
	mux.HandleFunc("/float_random", getRandomFloat)
	mux.HandleFunc("/float32_random", getRandomFloat32)

	http.ListenAndServe(":8080", mux)
}
