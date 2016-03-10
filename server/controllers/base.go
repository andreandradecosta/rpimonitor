package controllers

import (
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"

	"gopkg.in/mgo.v2"
	"gopkg.in/unrolled/render.v1"
)

type action func(w http.ResponseWriter, r *http.Request) error

type controller struct {
	*render.Render
	redisPool    *redis.Pool
	mongoSession *mgo.Session
}

func (c *controller) handleAction(a action) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := a(w, r); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
