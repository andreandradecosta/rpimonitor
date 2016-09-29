package daemon

import (
	"log"
	"time"

	"github.com/andreandradecosta/rpimonitor"
)

// Daemon controls a ticker that periodically sample metrics and persist it.
type Daemon struct {
	interval time.Duration
	reader   rpimonitor.SampleReader
	writer   rpimonitor.SampleWriter
}

// New creates and inits a Daemon
func New(i time.Duration, r rpimonitor.SampleReader, w rpimonitor.SampleWriter) *Daemon {
	return &Daemon{i, r, w}
}

// Start dispatches the ticker
func (d *Daemon) Start() {
	ticker := time.NewTicker(d.interval)
	for {
		d.sampleData()
		<-ticker.C
	}
}

func (d *Daemon) sampleData() {
	s, err := d.reader.ReadSample()
	if err != nil {
		log.Println("Error reading device data:", err)
	}
	err = d.writer.Write(s)
	if err != nil {
		log.Println("Error persisting data:", err)
	}
}
