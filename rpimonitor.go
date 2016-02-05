package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"gopkg.in/unrolled/render.v1"

	"github.com/andreandradecosta/rpimonitor/controllers/monitor"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/secure"
)

func main() {
	host := flag.String("HOST", "localhost", "Domain")
	httpPort := flag.String("HTTP_PORT", "", "HTTP port")
	httpsPort := flag.String("HTTPS_PORT", "443", "HTTPS port")
	isDev := flag.Bool("IsDevelopment", false, "Is Dev Env.")
	cert := flag.String("CERT", "cert.pem", "Certification path")
	key := flag.String("KEY", "key.pem", "Private Key path")
	flag.Parse()

	secureOptions := secure.Options{
		SSLRedirect:           true,
		SSLHost:               *host + ":" + *httpsPort,
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IsDevelopment:         *isDev,
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

	l := log.New(os.Stdout, "[negroni] ", 0)

	// HTTP
	if *httpPort != "" {
		addr := ":" + *httpPort
		l.Printf("listening on http://%s%s", *host, addr)
		go func() {
			log.Fatal(http.ListenAndServe(addr, n))
		}()

	}
	// HTTPS
	httpsAddr := ":" + *httpsPort
	l.Printf("listening on https://%s%s", *host, httpsAddr)
	l.Fatal(http.ListenAndServeTLS(httpsAddr, *cert, *key, n))
}
