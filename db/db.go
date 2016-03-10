package db

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
)

//DB hold dbs connections
type DB struct {
	RedisPool    *redis.Pool
	MongoSession *mgo.Session
}

//NewDB configures the dbs connections
func NewDB(mongoURL, redisHost, redisPasswd string) *DB {
	return &DB{
		newRedisPool(redisHost, redisPasswd),
		newMongoSession(mongoURL),
	}
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
