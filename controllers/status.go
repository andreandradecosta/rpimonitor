package controllers

import (
	"net/http"

	"github.com/andreandradecosta/rpimonitor/models"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
)

type status struct {
	controller
}

func NewStatus(renderer *render.Render, router *mux.Router) {
	s := &status{controller{Render: renderer}}
	router.
		Methods("GET").
		Path("/").
		Name("Index").
		Handler(s.handleAction(s.index))
}

func (s *status) index(w http.ResponseWriter, r *http.Request) error {
	s.JSON(w, http.StatusOK, models.NewStatus())
	return nil
}
