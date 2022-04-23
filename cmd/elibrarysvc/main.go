package main

import (
	"elibrarysvc/internal/middleware"
	"elibrarysvc/internal/repository"
	"elibrarysvc/internal/service"
	"elibrarysvc/internal/transport"
	"elibrarysvc/pkg/cache"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var services service.Services
	{
		memCache := cache.NewMemoryCache()
		repos := repository.NewInmemRepositories()
		services = service.NewServices(service.Deps{
			Repos:    repos,
			Cache:    memCache,
			CacheTTL: 0,
		})
		services = middleware.LoggingMiddleware(logger)(services)
	}

	var handler http.Handler
	{
		handler = transport.MakeHTTPHandler(services, nil, log.With(logger, "component", "HTTP"))
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%v", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	logger.Log("exit", <-errs)

}
