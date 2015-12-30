package main

import (
	"os"

	"github.com/andreandradecosta/rpimonitor/controllers/monitor"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	renderer := render.New(render.Options{
		IndentJSON: true,
	})

	c := monitor.New(renderer)
	c.Register(router)

	n := negroni.Classic()
	n.UseHandler(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	n.Run(":" + port)
}
