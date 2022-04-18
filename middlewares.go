package elibrarysvc

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) GetArticles(ctx context.Context) (as []Article, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetArticles", "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetArticles(ctx)
}

func (mw loggingMiddleware) GetArticle(ctx context.Context, id uint64) (a Article, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetArticle", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetArticle(ctx, id)
}
