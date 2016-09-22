package redis

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/andreandradecosta/rpimonitor"
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
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

func (u *UserService) Authenticate(login, password string) (*rpimonitor.User, error) {
	c := u.RedisPool.Get()
	defer c.Close()
	userData, err := redis.StringMap(c.Do("HGETALL", fmt.Sprintf("user:%s", login)))
	if err != nil {
		return nil, errors.Wrapf(err, "Fetch %s user failed", login)
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData["hash"]), []byte(password))
	if err != nil {
		return nil, nil
	}
	return &rpimonitor.User{Login: login, Name: userData["name"]}, nil
}
