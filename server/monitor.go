package server

import (
	"encoding/json"
	"log"
	"time"

	"github.com/andreandradecosta/rpimonitor/models"
	"github.com/garyburd/redigo/redis"
)

type Monitor struct {
	Interval  time.Duration
	RedisPool *redis.Pool
}

func (m *Monitor) Start() {
	ticker := time.NewTicker(m.Interval)
	for {
		m.tick()
		<-ticker.C
	}
}

func (m *Monitor) tick() {
	conn := m.RedisPool.Get()
	defer conn.Close()
	conn.Send("MULTI")
	s := models.NewSample()
	conn.Send("SET", "updated", s.Timestamp)
	sample := makeJSON(s)
	conn.Send("SET", "snapshot", sample)
	status := makeJSON(models.NewStatus())
	conn.Send("SET", "status", status)
	_, err := conn.Do("EXEC")
	if err != nil {
		log.Println(err)
	}
}

func makeJSON(data interface{}) []byte {
	json, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return nil
	}
	return json
}
