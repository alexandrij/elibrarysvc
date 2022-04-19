package middleware

import (
	"context"
	"elibrarysvc/internal/domain"
	"elibrarysvc/internal/service"
	"github.com/go-kit/kit/log"
	"time"
)

type ArticleMiddleware func(service.Articles) service.Articles

func ArticlesLoggingMiddleware(logger log.Logger) ArticleMiddleware {
	return func(next service.Articles) service.Articles {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func (mw loggingMiddleware) GetArticles(ctx context.Context) (a []domain.Article, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetArticles", "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetArticles(ctx)
}

func (mw loggingMiddleware) GetArticle(ctx context.Context, id domain.ArticleID) (a domain.Article, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetArticle", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetArticle(ctx, id)
}

func (mw loggingMiddleware) PostArticle(ctx context.Context, a domain.Article) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "PostArticle", "id", a.ID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PostArticle(ctx, a)
}

func (mw loggingMiddleware) DeleteArticle(ctx context.Context, id domain.ArticleID) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "DeleteArticle", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.DeleteArticle(ctx, id)
}
