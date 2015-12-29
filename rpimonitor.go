package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
)

func MyMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("MyMiddleware:", r.RequestURI)
	next(w, r)
}

func main() {
	router := NewRouter()

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(MyMiddleware))
	n.UseHandler(router)

	n.Run(":8080")
}
