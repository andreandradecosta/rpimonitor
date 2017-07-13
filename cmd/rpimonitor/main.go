package main

import (
	"log"
	"time"

	"github.com/andreandradecosta/rpimonitor/daemon"
	"github.com/andreandradecosta/rpimonitor/echo"
	"github.com/andreandradecosta/rpimonitor/hw"
	"github.com/andreandradecosta/rpimonitor/mongo"
	"github.com/andreandradecosta/rpimonitor/usersdb"
	"github.com/namsral/flag"
)

var (
	commit  string
	builtAt string
)

func main() {
	log.Printf("Build info: %s @ %s", commit, builtAt)

	config := flag.String("config", "", "Config file path")
	mongoURL := flag.String("MONGO_URL", "localhost", "mongodb://user:pass@host:port/database")
	sampleInterval := flag.Duration("SAMPLE_INTERVAL", time.Second*10, "Sampling interval")
	jwtSigningKey := flag.String("JWT_SIGNING_KEY", "", "JWT Signing Key")
	staticDir := flag.String("STATIC_DIR", "web", "Web content dir")
	usersFile := flag.String("USERS_DB", "users.db", "Users db file")
	certFile := flag.String("CERT_FILE", "./cert.pem", "Certificate file path")
	keyFile := flag.String("KEY_FILE", "./key.pem", "Key file path")
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
	userService := usersdb.NewUserService(*usersFile)

	log.Println("Starting HTTP server...")
	e := echo.New(*jwtSigningKey, *certFile, *keyFile,
		echo.WithDevice(hardware),
		echo.WithSampleFetcher(mongo),
		echo.WithUserManager(userService),
		echo.WithStaticDir(*staticDir),
		echo.WithDebug(*debug))
	go e.Start()

	log.Println("Starting monitor...")
	daemon := daemon.New(*sampleInterval, hardware, mongo)
	daemon.Start()

}
