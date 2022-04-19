package repository

import (
	"context"
	"elibrarysvc/internal/domain"
)

type Articles interface {
	GetArticles(ctx context.Context) ([]domain.Article, error)
	GetArticle(ctx context.Context, id domain.ArticleID) (domain.Article, error)
	PostArticle(ctx context.Context, a domain.Article) error
	DeleteArticle(ctx context.Context, id domain.ArticleID) error
}
