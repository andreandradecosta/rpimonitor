package main

import (
	"log"
	"time"

	"github.com/andreandradecosta/rpimonitor/db"
	"github.com/andreandradecosta/rpimonitor/monitor"
	"github.com/namsral/flag"
)

func main() {
	config := flag.String("config", "", "Config file path")
	sampleInterval := flag.Duration("SAMPLE_INTERVAL", time.Second*5, "Sampling interval")
	redisHost := flag.String("REDIS_HOST", "localhost:6379", "Redis host:port")
	redisPasswd := flag.String("REDIS_PASSWD", "", "Redis password")
	mongoURL := flag.String("MONGO_URL", "localhost", "mongodb://user:pass@host:port/database")

	flag.Parse()

	log.Println("Starting monitor...")
	if *config != "" {
		log.Println("Using ", *config)
	}

	db := db.NewDB(*mongoURL, *redisHost, *redisPasswd)
	m := &monitor.Monitor{
		Interval:     *sampleInterval,
		RedisPool:    db.RedisPool,
		MongoSession: db.MongoSession,
	}
	m.Start()
}
