package monitor

import (
	"net/http"
	"time"

	"github.com/andreandradecosta/rpimonitor/controllers"
	"github.com/andreandradecosta/rpimonitor/models"
	"github.com/gorilla/mux"

	"gopkg.in/unrolled/render.v1"
)

type Monitor struct {
	controllers.Controller
	*render.Render
}

func New(renderer *render.Render) *Monitor {
	return &Monitor{Render: renderer}
}

func (m *Monitor) Register(router *mux.Router) {
	router.
		Methods("GET").
		Path("/").
		Name("Index").
		Handler(m.Action(m.Index))
	router.
		Methods("GET").
		Path("/temp").
		Name("Temperature").
		Handler(m.Action(m.Temp))
	router.
		Methods("GET").
		Path("/cpu").
		Name("CPU").
		Handler(m.Action(m.CPU))
	router.
		Methods("GET").
		Path("/mem").
		Name("Memory").
		Handler(m.Action(m.Mem))
	router.
		Methods("GET").
		Path("/sys").
		Name("Sys").
		Handler(m.Action(m.Sys))
}

func (m *Monitor) Index(w http.ResponseWriter, r *http.Request) error {
	m.Text(w, http.StatusOK, "Index")
	return nil
}

func (m *Monitor) Temp(w http.ResponseWriter, r *http.Request) error {
	temp := models.Sample{
		Name:  "Temperature",
		Time:  time.Now(),
		Value: "1",
	}
	m.JSON(w, http.StatusOK, temp)
	return nil
}

func (m *Monitor) CPU(w http.ResponseWriter, r *http.Request) error {
	cpu := models.Sample{
		Name:  "CPU Load",
		Time:  time.Now(),
		Value: "1",
	}
	m.JSON(w, http.StatusOK, cpu)
	return nil
}

func (m *Monitor) Mem(w http.ResponseWriter, r *http.Request) error {
	mem := models.Sample{
		Name:  "Free Memory",
		Time:  time.Now(),
		Value: "1",
	}
	m.JSON(w, http.StatusOK, mem)
	return nil
}

func (m *Monitor) Sys(w http.ResponseWriter, r *http.Request) error {
	sys := models.Sample{
		Name:  "Sys Uptime",
		Time:  time.Now(),
		Value: "1",
	}
	m.JSON(w, http.StatusOK, sys)
	return nil
}
