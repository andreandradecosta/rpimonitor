package server

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/andreandradecosta/rpimonitor/models"
	"github.com/garyburd/redigo/redis"
)

type Monitor struct {
	Interval  time.Duration
	RedisPool *redis.Pool
	log       *log.Logger
}

func (m *Monitor) Start() {
	m.log = log.New(os.Stdout, "[monitor] ", 0)
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
	sample := m.toJSON(s)
	conn.Send("SET", "snapshot", sample)
	status := m.toJSON(models.NewIndex())
	conn.Send("SET", "status", status)
	_, err := conn.Do("EXEC")
	if err != nil {
		m.log.Println(err)
	}
}

func (m *Monitor) toJSON(data interface{}) []byte {
	json, err := json.Marshal(data)
	if err != nil {
		m.log.Println(err)
		return nil
	}
	return json
}
