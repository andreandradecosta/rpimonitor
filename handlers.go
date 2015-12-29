package main

import (
	"log"
	"net/http"
	"time"

	"gopkg.in/unrolled/render.v1"

	"github.com/andreandradecosta/rpimonitor/models"
)

var rd *render.Render

func init() {
	log.Print("Init")
	rd = render.New(render.Options{
		IndentJSON: true,
	})

}

func Index(w http.ResponseWriter, r *http.Request) {
	rd.Text(w, http.StatusOK, "Index")
}

func TempShow(w http.ResponseWriter, r *http.Request) {
	temp := models.Sample{
		Name:  "Temperature",
		Time:  time.Now(),
		Value: "1",
	}
	rd.JSON(w, http.StatusOK, temp)
}

func CPUShow(w http.ResponseWriter, r *http.Request) {
	cpu := models.Sample{
		Name:  "CPU Load",
		Time:  time.Now(),
		Value: "1",
	}
	rd.JSON(w, http.StatusOK, cpu)
}

func MemShow(w http.ResponseWriter, r *http.Request) {
	mem := models.Sample{
		Name:  "Free Memory",
		Time:  time.Now(),
		Value: "1",
	}
	rd.JSON(w, http.StatusOK, mem)
}

func SysShow(w http.ResponseWriter, r *http.Request) {
	sys := models.Sample{
		Name:  "Sys Uptime",
		Time:  time.Now(),
		Value: "1",
	}
	rd.JSON(w, http.StatusOK, sys)
}
