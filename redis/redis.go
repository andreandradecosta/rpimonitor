package redis

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/andreandradecosta/rpimonitor"
	"github.com/garyburd/redigo/redis"
)

type UserService struct {
	RedisPool *redis.Pool
}

func NewUserService(redisHost, redisPassword string) *UserService {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", redisPassword); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     1,
		MaxActive:   5,
		IdleTimeout: 5 * time.Hour,
		Wait:        true,
	}
	return &UserService{pool}
}

func (u *UserService) Fetch(login string) (rpimonitor.User, error) {
	c := u.RedisPool.Get()
	defer c.Close()
	name, err := redis.String(c.Do("GET", fmt.Sprintf("user:%s:name", login)))
	if err != nil {
		return rpimonitor.User{}, err
	}
	if name != "" {
		return rpimonitor.User{Login: login, Name: name}, nil
	}
	return rpimonitor.User{}, rpimonitor.NotFound
}

func (u *UserService) Authenticate(login, password string) (bool, error) {
	c := u.RedisPool.Get()
	defer c.Close()
	hash, err := redis.String(c.Do("GET", fmt.Sprintf("user:%s:hash", login)))
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, nil
}
