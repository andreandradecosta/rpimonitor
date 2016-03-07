package controllers

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

type action func(w http.ResponseWriter, r *http.Request) error

type controller struct {
	*render.Render
}

func (c *controller) handleAction(a action) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := a(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
