package controllers

import (
	"errors"
	"net/http"

	"github.com/andreandradecosta/rpimonitor/models"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
)

type sample struct {
	controller
}

func NewSample(renderer *render.Render, router *mux.Router) {
	s := &sample{controller{Render: renderer}}
	router.
		Methods("GET").
		Path("/snapshot").
		Name("Snapshot").
		Handler(s.handleAction(s.snapshot))
	router.
		Methods("GET").
		Path("/history").
		Name("History").
		Handler(s.handleAction(s.history))
}

//Usar session.Copy e depois session.Close()
func (s *sample) snapshot(w http.ResponseWriter, r *http.Request) error {
	s.JSON(w, http.StatusOK, models.NewSample())
	return nil
}

func (s *sample) history(w http.ResponseWriter, r *http.Request) error {
	return errors.New("Not implemented")
}
