package repository

import (
	"context"
	"elibrarysvc/internal/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Articles interface {
	GetArticles(ctx context.Context) ([]domain.Article, error)
	GetArticle(ctx context.Context, id domain.ArticleID) (domain.Article, error)
	PostArticle(ctx context.Context, a domain.Article) error
	DeleteArticle(ctx context.Context, id domain.ArticleID) error
}

type Repositories struct {
	Articles Articles
}

func NewInmemRepositories() *Repositories {
	return &Repositories{
		Articles: NewArticlesInmemRepo(),
	}
}

func NewPgxRepositories(conn *pgxpool.Conn) *Repositories {
	return &Repositories{
		Articles: NewArticlesPgxRepo(conn),
	}
}
