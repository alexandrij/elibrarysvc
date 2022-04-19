package articles

import (
	"context"
	"elibrarysvc/internal/domain"
	"elibrarysvc/internal/repository"
)

type ArticlesService struct {
	repo repository.Articles
}

func NewArticlesService(repo repository.Articles) *ArticlesService {
	return &ArticlesService{repo: repo}
}

func (s *ArticlesService) GetArticles(ctx context.Context) ([]domain.Article, error) {
	return s.repo.GetArticles(ctx)
}

func (s *ArticlesService) GetArticle(ctx context.Context, id domain.ArticleID) (domain.Article, error) {
	return s.repo.GetArticle(ctx, id)
}

func (s *ArticlesService) PostArticle(ctx context.Context, a domain.Article) error {
	return s.repo.PostArticle(ctx, a)
}

func (s *ArticlesService) DeleteArticle(ctx context.Context, id domain.ArticleID) error {
	return s.repo.DeleteArticle(ctx, id)
}
