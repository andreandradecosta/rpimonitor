package main

import (
	"log"
	"time"

	"github.com/andreandradecosta/rpimonitor/daemon"
	"github.com/andreandradecosta/rpimonitor/echo"
	"github.com/andreandradecosta/rpimonitor/hw"
	"github.com/andreandradecosta/rpimonitor/mongo"
	"github.com/andreandradecosta/rpimonitor/redis"
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
	sampleInterval := flag.Duration("SAMPLE_INTERVAL", time.Second*10, "Sampling interval")
	jwtSigningKey := flag.String("JWT_SIGNING_KEY", "", "JWT Signing Key")
	staticDir := flag.String("STATIC_DIR", "web", "Web content dir")
	debug := flag.Bool("DEBUG", false, "HTTP Server debug")

	flag.Parse()

	if *config != "" {
		log.Println("Using ", *config)
	}

	hardware := &hw.Hardware{}
	mongo, err := mongo.NewSampleService(*mongoURL)
	if err != nil {
		log.Println("Mongo:", err)
	}
	redis := redis.NewUserService(*redisHost, *redisPasswd)

	log.Println("Starting HTTP server...")
	e := echo.New(*jwtSigningKey,
		echo.WithDevice(hardware),
		echo.WithSampleFetcher(mongo),
		echo.WithUserManager(redis),
		echo.WithStaticDir(*staticDir),
		echo.WithDebug(*debug))
	go e.Start()

	log.Println("Starting monitor...")
	daemon := daemon.New(*sampleInterval, hardware, mongo)
	daemon.Start()

}
