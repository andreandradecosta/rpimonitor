package main

import (
	"log"
	"time"

	"github.com/andreandradecosta/rpimonitor/db"
	"github.com/andreandradecosta/rpimonitor/monitor"
	"github.com/namsral/flag"
)

var (
	commit  string
	builtAt string
)

func main() {
	log.Printf("Build info: %s @ %s", commit, builtAt)

	config := flag.String("config", "", "Config file path")
	redisHost := flag.String("REDIS_HOST", "localhost:6379", "Redis host:port")
	redisPasswd := flag.String("REDIS_PASSWD", "", "Redis password")
	mongoURL := flag.String("MONGO_URL", "localhost", "mongodb://user:pass@host:port/database")
	sampleInterval := flag.Duration("SAMPLE_INTERVAL", time.Minute*10, "Sampling interval")

	flag.Parse()

	log.Println("starting monitor...")
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
