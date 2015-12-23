package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()
	n := negroni.Classic()
	n.UseHandler(mux)

	n.Run(":8080")
}
