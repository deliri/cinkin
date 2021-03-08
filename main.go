package main

import (
	"cinkin/api"
	"cinkin/web"
	"context"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/mattn/go-colorable"
	mw "github.com/neko-neko/echo-logrus/v2"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
)

// Initalize  webrouter
var webRouter = web.WebRoutes()

// Initalize APIrouter
var apiRouter = api.APIRoutes()

// General router
var e = echo.New()

// maximizes the number of CPUs avail on the server
func init() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
}

// create hosts for domain and subdomains
type (
	Host struct {
		Echo *echo.Echo
	}
)

// MatchRoutes combines the subdomain routers into a cohesive single version
func MatchRoutes(c echo.Context) (err error) {
	// Hosts
	var hosts = map[string]*Host{}
	var mutex = &sync.Mutex{}
	// APIs
	mutex.Lock()
	hosts["api.localhost:8080"] = &Host{apiRouter}
	mutex.Unlock()
	// Web
	mutex.Lock()
	hosts["localhost:8080"] = &Host{webRouter}
	mutex.Unlock()
	req := c.Request()
	res := c.Response()
	// how to route requests based on HOST
	host := hosts[req.Host]
	if host == nil {
		err = echo.ErrNotFound
	} else {
		host.Echo.ServeHTTP(res, req)
	}
	return nil
}

func main() {

	// Gather routes according to domain/subdomain
	e.Any("/*", MatchRoutes)
	// set Debug to true during dev and false in prod
	e.Debug = false
	// recover middleware
	e.Use(middleware.Recover())
	// create an HTTP2 server with reasonable limits
	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          100 * time.Second,
	}
	// PAth for the log files
	const logPath = "logs/"
	runID := time.Now().Format("run-2006-01-02-15-04-05")
	logLocation := filepath.Join(logPath, runID+".log")
	logFile, err := os.OpenFile(logLocation, os.O_CREATE|os.O_WRONLY, 0644)
	// close the file so there's no memory leakage
	defer logFile.Close()
	if err != nil {
		log.Fatalf("Failed to open log file %s for output: %s", logLocation, err)
	}
	wrt := io.MultiWriter(os.Stdout, logFile)

	// Logger
	log.Logger().SetLevel(echoLog.DEBUG)
	log.Logger().SetReportCaller(true)
	logrus.SetOutput(colorable.NewColorableStdout())

	log.Logger().SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		TimestampFormat:  time.RFC3339,
		FullTimestamp:    true,
		DisableTimestamp: false,
	})
	log.Logger().SetOutput(wrt)
	// Web router to log its activity
	webRouter.Logger = log.Logger()
	webRouter.Use(mw.Logger())
	// API router to log its activity
	apiRouter.Logger = log.Logger()
	apiRouter.Use(mw.Logger())

	// Start server
	go func() {
		if err := e.StartH2CServer(":8080", s); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(),
		100*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
