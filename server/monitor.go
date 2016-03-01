package server

import (
	"fmt"
	"time"

	"github.com/andreandradecosta/rpimonitor/models"
)

type Monitor struct{}

func (m *Monitor) Start() {
	ticker := time.NewTicker(time.Second * 5)
	for t := range ticker.C {
		fmt.Println("Tick at", t)
		fmt.Println(models.NewSample())
	}
}
