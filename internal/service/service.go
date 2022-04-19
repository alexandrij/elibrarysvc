package service

import (
	"context"
	"elibrarysvc/internal/domain"
	"elibrarysvc/internal/repository"
	"elibrarysvc/pkg/cache"
)

type Articles interface {
	GetArticles(ctx context.Context) ([]domain.Article, error)
	GetArticle(ctx context.Context, id domain.ArticleID) (domain.Article, error)
	PostArticle(ctx context.Context, a domain.Article) error
	DeleteArticle(ctx context.Context, id domain.ArticleID) error
}

type Services struct {
	Articles Articles
}

type Deps struct {
	Repos    *repository.Repositories
	Cache    cache.Cache
	CacheTTL int64
}

func NewServices(deps Deps) *Services {
	articlesService := NewArticlesService(deps.Repos.Articles, deps.Cache, deps.CacheTTL)
	return &Services{
		Articles: articlesService,
	}
}
