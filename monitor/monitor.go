package monitor

import (
	"encoding/json"
	"log"
	"time"

	"gopkg.in/mgo.v2"

	"github.com/andreandradecosta/rpimonitor/models"
	"github.com/garyburd/redigo/redis"
)

//Monitor is the daemon that collects and saves data into dbs
type Monitor struct {
	Interval     time.Duration
	RedisPool    *redis.Pool
	MongoSession *mgo.Session
}

//Start starts the daemon at the specified Interval
func (m *Monitor) Start() {
	ticker := time.NewTicker(m.Interval)
	for {
		m.saveData()
		<-ticker.C
	}
}

func (m *Monitor) saveData() {
	defer m.MongoSession.Refresh()
	s := models.NewSample()
	c := m.MongoSession.DB("rpimonitor").C("samples")
	err := c.Insert(s)
	if err != nil {
		log.Println("saveDate:", err)
		return
	}
	m.cacheData(s)
}

func (m *Monitor) cacheData(s models.Sample) {
	conn := m.RedisPool.Get()
	defer conn.Close()
	conn.Send("MULTI")
	conn.Send("SET", "updated", s.Timestamp)
	sample := makeJSON(s)
	conn.Send("SET", "snapshot", sample)
	status := makeJSON(models.NewStatus())
	conn.Send("SET", "status", status)
	_, err := conn.Do("EXEC")
	if err != nil {
		log.Println("saveCache:", err)
	}
}

func makeJSON(data interface{}) []byte {
	json, err := json.Marshal(data)
	if err != nil {
		log.Println("makeJSON:", err)
		return nil
	}
	return json
}
