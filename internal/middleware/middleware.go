package middleware

import (
	"elibrarysvc/internal/service"
	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	next   service.Articles
	logger log.Logger
}

type Middleware func(services service.Services) service.Services

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(services service.Services) service.Services {
		services.Articles = ArticlesLoggingMiddleware(logger)(services.Articles)
		return services
	}
}
