package models

import "time"

type Sample struct {
	Name  string    `json:"name"`
	Time  time.Time `json:"time"`
	Value string    `json:"value"`
}
