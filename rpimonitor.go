package main

import (
	"log"
	"time"

	"github.com/andreandradecosta/rpimonitor/server"
	"github.com/garyburd/redigo/redis"
	"github.com/namsral/flag"
	"gopkg.in/mgo.v2"
)

func main() {
	config := flag.String("config", "", "Config file path")
	startServer := flag.Bool("START_SERVER", false, "Start HTTP Server")
	host := flag.String("HOST", "localhost", "Domain")
	httpPort := flag.String("HTTP_PORT", "8080", "HTTP port")
	httpsPort := flag.String("HTTPS_PORT", "443", "HTTPS port")
	isDev := flag.Bool("IS_DEVELOPMENT", false, "Is Dev Env.")
	cert := flag.String("CERT", "cert.pem", "Certification path")
	key := flag.String("KEY", "key.pem", "Private Key path")
	sampleInterval := flag.Duration("SAMPLE_INTERVAL", time.Second*5, "Sampling interval")
	redisHost := flag.String("REDIS_HOST", "localhost:6379", "Redis host:port")
	redisPasswd := flag.String("REDIS_PASSWD", "", "Redis password")
	mongoURL := flag.String("MONGO_URL", "localhost", "mongodb://user:pass@host:port/database")

	flag.Parse()

	log.Println("Starting...")
	if *config != "" {
		log.Println("Using ", *config)
	}

	if *startServer {
		log.Println("... server")
		s := &server.HTTPServer{
			Host:      *host,
			HTTPPort:  *httpPort,
			HTTPSPort: *httpsPort,
			IsDev:     *isDev,
			Cert:      *cert,
			Key:       *key,
		}
		s.Start()
	}

	redisPool := newRedisPool(*redisHost, *redisPasswd)
	mongoSession := newMongoSession(*mongoURL)

	m := &server.Monitor{
		Interval:     *sampleInterval,
		RedisPool:    redisPool,
		MongoSession: mongoSession,
	}
	log.Println("... monitor")
	m.Start()
}

func newRedisPool(redisHost, redisPasswd string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				return nil, err
			}
			if redisPasswd != "" {
				if _, err := c.Do("AUTH", redisPasswd); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     3,
		MaxActive:   10,
		IdleTimeout: 5 * time.Hour,
		Wait:        true,
	}
}

func newMongoSession(mongoURL string) *mgo.Session {
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Fatalln("MongoDial:", err)
	}
	return session
}
