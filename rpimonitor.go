package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/andreandradecosta/rpimonitor/controllers/monitor"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/secure"
	"gopkg.in/unrolled/render.v1"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	httpsPort := os.Getenv("HTTPS_PORT")
	if httpsPort == "" {
		httpsPort = "8443"
	}
	isDev, err := strconv.ParseBool(os.Getenv("IsDevelopment"))
	if err != nil {
		isDev = false
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	secureOptions := secure.Options{
		AllowedHosts:          []string{host + ":" + port},
		SSLRedirect:           true,
		SSLHost:               host + ":" + httpsPort,
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		PublicKey:             `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`,
		IsDevelopment:         isDev,
	}
	secureMiddleware := secure.New(secureOptions)

	router := mux.NewRouter().StrictSlash(true)
	renderer := render.New(render.Options{
		IndentJSON: true,
	})

	c := monitor.New(renderer)
	c.Register(router)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext))
	n.UseHandler(router)

	addr := ":" + port
	httpsAddr := ":" + httpsPort
	l := log.New(os.Stdout, "[negroni] ", 0)
	l.Printf("secure options %v", secureOptions)
	l.Printf("listening on http://%s and https://%s", addr, httpsAddr)
	// HTTP
	go func() {
		log.Fatal(http.ListenAndServe(addr, n))
	}()
	l.Fatal(http.ListenAndServeTLS(httpsAddr, "cert.pem", "key.pem", n))
}
