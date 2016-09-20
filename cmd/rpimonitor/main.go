package main

import (
	"log"
	"time"

	"github.com/andreandradecosta/rpimonitor/daemon"
	"github.com/andreandradecosta/rpimonitor/device"
	"github.com/andreandradecosta/rpimonitor/echo"
	"github.com/andreandradecosta/rpimonitor/mongo"
	"github.com/namsral/flag"
)

var (
	commit  string
	builtAt string
)

func main() {
	log.Printf("Build info: %s @ %s", commit, builtAt)

	config := flag.String("config", "", "Config file path")
	// redisHost := flag.String("REDIS_HOST", "localhost:6379", "Redis host:port")
	// redisPasswd := flag.String("REDIS_PASSWD", "", "Redis password")
	mongoURL := flag.String("MONGO_URL", "localhost", "mongodb://user:pass@host:port/database")
	sampleInterval := flag.Duration("SAMPLE_INTERVAL", time.Second*10, "Sampling interval")

	flag.Parse()

	if *config != "" {
		log.Println("Using ", *config)
	}

	device := &device.Device{}
	mongo, err := mongo.NewSampleService(*mongoURL)

	log.Println("Starting HTTP server...")
	echo := &echo.Server{
		StatusReader:  device,
		SampleFetcher: mongo,
	}
	go echo.Start()

	log.Println("Starting monitor...")
	if err != nil {
		log.Println("Mongo:", err)
	}
	daemon := &daemon.Daemon{
		Interval: *sampleInterval,
		Reader:   device,
		Writer:   mongo,
	}
	daemon.Start()

}
