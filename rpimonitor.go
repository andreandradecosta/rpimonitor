package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/andreandradecosta/rpimonitor/server"
)

func main() {
	startServer := flag.Bool("StartServer", false, "Start HTTP Server")
	host := flag.String("HOST", "localhost", "Domain")
	httpPort := flag.String("HTTP_PORT", "8080", "HTTP port")
	httpsPort := flag.String("HTTPS_PORT", "443", "HTTPS port")
	isDev := flag.Bool("IsDevelopment", false, "Is Dev Env.")
	cert := flag.String("CERT", "cert.pem", "Certification path")
	key := flag.String("KEY", "key.pem", "Private Key path")
	flag.Parse()

	if *startServer {
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
	startMonitor()
}

func startMonitor() {
	ticker := time.NewTicker(time.Second * 5)
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}
