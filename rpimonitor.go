package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index")
}
func ProcIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Proc Index")
}
func ProcShow(w http.ResponseWriter, r *http.Request) {
	proc := mux.Vars(r)["procID"]
	fmt.Fprintf(w, "Proc Show: %s", proc)
}
func TempShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Temp Show")
}
func CPUShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "CPU Show")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/proc", ProcIndex)
	router.HandleFunc("/proc/{procID}", ProcShow)
	router.HandleFunc("/hw/temp", TempShow)
	router.HandleFunc("/hw/cpu", CPUShow)

	n := negroni.Classic()
	n.UseHandler(router)

	n.Run(":8080")
}
