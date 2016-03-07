package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/andreandradecosta/rpimonitor/models"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
)

type status struct {
	controller
}

//NewStatus creates this controller and register it with router.
func NewStatus(renderer *render.Render, router *mux.Router, redisPool *redis.Pool) {
	s := &status{controller{Render: renderer, redisPool: redisPool}}
	router.
		Methods("GET").
		Path("/").
		Name("Index").
		Handler(s.handleAction(s.index))
}

func (s *status) index(w http.ResponseWriter, r *http.Request) error {
	conn := s.redisPool.Get()
	defer conn.Close()
	b, err := redis.Bytes(conn.Do("GET", "status"))
	if err != nil {
		return err
	}
	var d models.Info
	err = json.Unmarshal(b, &d)
	if err != nil {
		return err
	}
	s.JSON(w, http.StatusOK, d)
	return nil
}
