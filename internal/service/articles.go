package service

import (
	"context"
	"elibrarysvc/internal/domain"
	"elibrarysvc/internal/repository"
	"elibrarysvc/pkg/cache"
)

type ArticlesService struct {
	repo  repository.Articles
	cache cache.Cache
	ttl   int64
}

func NewArticlesService(repo repository.Articles, cache cache.Cache, ttl int64) *ArticlesService {
	return &ArticlesService{repo: repo, cache: cache, ttl: ttl}
}

func (s *ArticlesService) GetArticles(ctx context.Context) ([]domain.Article, error) {
	return s.repo.GetArticles(ctx)
}

func (s *ArticlesService) GetArticle(ctx context.Context, id domain.ArticleID) (domain.Article, error) {
	if value, err := s.cache.Get(id); err == nil {
		return value.(domain.Article), nil
	}

	article, err := s.repo.GetArticle(ctx, id)
	if err != nil {
		return domain.Article{}, nil
	}

	err = s.cache.Set(id, article, s.ttl)
	return article, err
}

func (s *ArticlesService) PostArticle(ctx context.Context, a domain.Article) error {
	return s.repo.PostArticle(ctx, a)
}

func (s *ArticlesService) DeleteArticle(ctx context.Context, id domain.ArticleID) error {
	return s.repo.DeleteArticle(ctx, id)
}
