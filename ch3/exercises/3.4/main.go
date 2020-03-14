package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dartharthur/gopl-study-notes/ch3/exercises/surface"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		// there must be a standard/accepted way to convert query params to proper type
		// this seems janky
		color := strings.Join(params["color"][:], "")
		// ignoring error is...not good
		width, _ := strconv.ParseFloat(strings.Join(params["width"][:], ""), 64)
		height, _ := strconv.ParseFloat(strings.Join(params["height"][:], ""), 64)
		// ------------------------------------------------------------------------
		w.Header().Set("Content-Type", "image/svg+xml")
		surface.Render(w, color, width, height)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
