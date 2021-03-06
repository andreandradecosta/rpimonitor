package echo

import (
	"github.com/andreandradecosta/rpimonitor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

// Server is responsible to start the echo HTTP server.
type Server struct {
	device        rpimonitor.Device
	sampleFetcher rpimonitor.SampleFetcher
	userManager   rpimonitor.UserManager
	jwtSigningKey string
	certFile      string
	keyFile       string
	staticDir     string
	debug         bool
}

// Option is functional argument
type Option func(*Server)

// New creates and configures a Echo HTTP Server
func New(key string, certFile string, keyFile string, options ...Option) *Server {
	s := &Server{
		jwtSigningKey: key,
		certFile:      certFile,
		keyFile:       keyFile,
	}
	for _, opt := range options {
		opt(s)
	}
	return s
}

// WithStaticDir sets the static files path to the echo HTTP Server
func WithStaticDir(d string) Option {
	return func(s *Server) {
		s.staticDir = d
	}
}

// WithDebug sets the debug options of echo HTTP Server
func WithDebug(d bool) Option {
	return func(s *Server) {
		s.debug = d
	}
}

// WithDevice sets the component responsible for collecting device status. (e.g. Hardware)
func WithDevice(d rpimonitor.Device) Option {
	return func(s *Server) {
		s.device = d
	}
}

// WithSampleFetcher sets the component responsible for searching for samples.
func WithSampleFetcher(sf rpimonitor.SampleFetcher) Option {
	return func(s *Server) {
		s.sampleFetcher = sf
	}
}

// WithUserManager sets the component for fetching and authenticating users.
func WithUserManager(um rpimonitor.UserManager) Option {
	return func(s *Server) {
		s.userManager = um
	}
}

// Start configures the echo framework and starts the HTTP server.
func (s *Server) Start() {
	e := echo.New()
	e.SetLogLevel(log.ERROR)
	e.SetDebug(s.debug)
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Secure())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} - ${method}, ${uri}, [${status}]\n",
	}))

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  s.staticDir,
		HTML5: true,
	}))
	e.POST("/auth", s.login)

	r := e.Group("/api")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(s.jwtSigningKey),
		SigningMethod: "HS256",
	}))
	r.GET("/status", s.status)
	r.GET("/history", s.history)
	r.GET("/snapshot", s.snapshot)
	e.Run(standard.WithTLS(
		":8443",
		s.certFile,
		s.keyFile,
	))
}
