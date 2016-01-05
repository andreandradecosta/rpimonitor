package monitor

import (
	"net/http"

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
		Path("/snapshot").
		Name("Snapshot").
		Handler(m.Action(m.Snapshot))
}

func (m *Monitor) Index(w http.ResponseWriter, r *http.Request) error {
	m.JSON(w, http.StatusOK, models.NewIndex())
	return nil
}

func (m *Monitor) Snapshot(w http.ResponseWriter, r *http.Request) error {
	m.JSON(w, http.StatusOK, models.NewSample())
	return nil
}
