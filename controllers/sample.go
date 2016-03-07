package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/andreandradecosta/rpimonitor/models"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/unrolled/render.v1"
)

type sample struct {
	controller
}

//NewSample creates this controller and register it with router.
func NewSample(renderer *render.Render, router *mux.Router, redisPool *redis.Pool, mongoSession *mgo.Session) {
	s := &sample{
		controller{
			Render:       renderer,
			redisPool:    redisPool,
			mongoSession: mongoSession,
		},
	}
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

func (s *sample) snapshot(w http.ResponseWriter, r *http.Request) error {
	conn := s.redisPool.Get()
	defer conn.Close()
	b, err := redis.Bytes(conn.Do("GET", "snapshot"))
	if err != nil {
		return err
	}
	var d models.Sample
	err = json.Unmarshal(b, &d)
	if err != nil {
		return err
	}
	s.JSON(w, http.StatusOK, d)
	return nil
}

func (s *sample) history(w http.ResponseWriter, r *http.Request) error {
	start, err1 := time.Parse("2006-01-02", r.FormValue("start"))
	stop, err2 := time.Parse("2006-01-02", r.FormValue("stop"))
	log.Println(stop)
	if err1 != nil || err2 != nil {
		return fmt.Errorf("Inavlid parameters: %v", r.Form)
	}

	session := s.mongoSession.Copy()
	defer session.Close()
	c := session.DB("rpimonitor").C("samples")
	result := make([]models.Sample, 1)

	err := c.
		Find(bson.M{
		"localtime": bson.M{
			"$gte": start,
			"$lte": stop,
		},
	}).
		Select(bson.M{
		"localtime":              1,
		"timestamp":              1,
		"metrics.virtual_memory": 1,
	}).
		Sort("-localtime").
		All(&result)
	if err != nil {
		return err
	}
	s.JSON(w, http.StatusOK, result)
	return nil
}
