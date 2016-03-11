package main

import (
	"log"

	"github.com/andreandradecosta/rpimonitor/db"
	"github.com/andreandradecosta/rpimonitor/server"
	"github.com/namsral/flag"
)

var (
	commit  string
	builtAt string
)

func main() {
	log.Printf("Build info: %s @ %s", commit, builtAt)
	
	config := flag.String("config", "", "Config file path")
	host := flag.String("HOST", "localhost", "Domain")
	httpPort := flag.String("HTTP_PORT", "8080", "HTTP port")
	httpsPort := flag.String("HTTPS_PORT", "443", "HTTPS port")
	isDev := flag.Bool("IS_DEVELOPMENT", false, "Is Dev Env.")
	cert := flag.String("CERT", "cert.pem", "Certification path")
	key := flag.String("KEY", "key.pem", "Private Key path")
	redisHost := flag.String("REDIS_HOST", "localhost:6379", "Redis host:port")
	redisPasswd := flag.String("REDIS_PASSWD", "", "Redis password")
	mongoURL := flag.String("MONGO_URL", "localhost", "mongodb://user:pass@host:port/database")

	flag.Parse()

	log.Println("Starting server...")
	if *config != "" {
		log.Println("Using ", *config)
	}
	db := db.NewDB(*mongoURL, *redisHost, *redisPasswd)

	s := &server.HTTPServer{
		Host:         *host,
		HTTPPort:     *httpPort,
		HTTPSPort:    *httpsPort,
		IsDev:        *isDev,
		Cert:         *cert,
		Key:          *key,
		RedisPool:    db.RedisPool,
		MongoSession: db.MongoSession,
	}
	s.Start()

}
