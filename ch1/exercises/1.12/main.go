// Server1 is a minimal "echo" server.
package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		count := r.URL.Query().Get("count")
		convCount, err := strconv.Atoi(count)
		if err != nil {
			log.Fatal(err)
		}
		lissajous(w, float64(convCount))
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
