package daemon

import (
	"log"
	"time"

	"github.com/andreandradecosta/rpimonitor"
)

//Daemon controls a ticker that periodically sample metrics and persist it.
type Daemon struct {
	Interval time.Duration
	Reader   rpimonitor.SampleReader
	Writer   rpimonitor.SampleWriter
}

//Start dispatches the ticker
func (d *Daemon) Start() {
	ticker := time.NewTicker(d.Interval)
	for {
		d.sampleData()
		<-ticker.C
	}
}

func (d *Daemon) sampleData() {
	s, err := d.Reader.ReadSample()
	if err != nil {
		log.Println("Error reading device data:", err)
	}
	err = d.Writer.Write(s)
	if err != nil {
		log.Println("Error persisting data:", err)
	}
}
