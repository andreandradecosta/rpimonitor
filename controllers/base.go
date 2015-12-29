package controllers

import "net/http"

type Action func(w http.ResponseWriter, r *http.Request) error

type Controller struct{}

func (c *Controller) Action(a Action) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := a(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
