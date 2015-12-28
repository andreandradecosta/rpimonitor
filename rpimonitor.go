package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/andreandradecosta/rpimonitor/models"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index")
}

func TempShow(w http.ResponseWriter, r *http.Request) {
	temp := models.Sample{
		Name:  "Temperature",
		Time:  time.Now(),
		Value: "1",
	}
	json.NewEncoder(w).Encode(temp)
}

func CPUShow(w http.ResponseWriter, r *http.Request) {
	cpu := models.Sample{
		Name:  "CPU Load",
		Time:  time.Now(),
		Value: "1",
	}
	json.NewEncoder(w).Encode(cpu)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/hw/temp", TempShow)
	router.HandleFunc("/hw/cpu", CPUShow)

	n := negroni.Classic()
	n.UseHandler(router)

	n.Run(":8080")
}
