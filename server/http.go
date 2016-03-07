package server

import (
	"log"
	"net/http"
	"os"

	"github.com/andreandradecosta/rpimonitor/controllers"
	"github.com/codegangsta/negroni"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/unrolled/secure"
	"gopkg.in/mgo.v2"
	"gopkg.in/unrolled/render.v1"
)

//HTTPServer represents the HTTP part of this app.
type HTTPServer struct {
	Host         string
	HTTPPort     string
	HTTPSPort    string
	IsDev        bool
	Cert         string
	Key          string
	RedisPool    *redis.Pool
	MongoSession *mgo.Session
}

//Start configures the http server and starts listening requests.
func (h *HTTPServer) Start() {
	secureOptions := secure.Options{
		SSLRedirect:           true,
		SSLHost:               h.Host + ":" + h.HTTPSPort,
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IsDevelopment:         h.IsDev,
	}
	secureMiddleware := secure.New(secureOptions)

	router := mux.NewRouter().StrictSlash(true)
	renderer := render.New(render.Options{
		IndentJSON: true,
	})
	controllers.NewStatus(renderer, router, h.RedisPool)
	controllers.NewSample(renderer, router, h.RedisPool, h.MongoSession)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext))
	n.UseHandler(router)

	l := log.New(os.Stdout, "[negroni] ", 0)

	// HTTP
	if h.HTTPPort != "" {
		addr := ":" + h.HTTPPort
		l.Printf("listening on http://%s%s", h.Host, addr)
		go func() {
			l.Fatal(http.ListenAndServe(addr, n))
		}()

	}
	// HTTPS
	if h.HTTPSPort != "" {
		httpsAddr := ":" + h.HTTPSPort
		l.Printf("listening on https://%s%s", h.Host, httpsAddr)
		go func() {
			l.Fatal(http.ListenAndServeTLS(httpsAddr, h.Cert, h.Key, n))
		}()
	}

}
