package daemon

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/andreandradecosta/rpimonitor/device"
)

type Daemon struct {
	Interval time.Duration
	Device   *device.Device
}

func (d *Daemon) Start() {
	ticker := time.NewTicker(d.Interval)
	for {
		d.sampleData()
		<-ticker.C
	}
}

func (d *Daemon) sampleData() {
	s, _ := d.Device.ReadSample()
	b, _ := json.Marshal(s)
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	log.Println(string(out.Bytes()))
}
