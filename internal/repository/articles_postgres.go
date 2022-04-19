package repository

import (
	"context"
	"elibrarysvc/internal/domain"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrNotImplementation = errors.New("api not implementation")
)

type ArticlesPgxRepo struct {
	conn *pgxpool.Conn
}

func NewArticlesPgxRepo(conn *pgxpool.Conn) *ArticlesPgxRepo {
	return &ArticlesPgxRepo{
		conn: conn,
	}
}

func (r *ArticlesPgxRepo) GetArticles(ctx context.Context) ([]domain.Article, error) {
	return nil, ErrNotImplementation
}

func (r *ArticlesPgxRepo) GetArticle(ctx context.Context, id domain.ArticleID) (domain.Article, error) {
	return domain.Article{}, ErrNotImplementation
}

func (r *ArticlesPgxRepo) PostArticle(ctx context.Context, a domain.Article) error {
	return ErrNotImplementation
}

func (r *ArticlesPgxRepo) DeleteArticle(ctx context.Context, id domain.ArticleID) error {
	return ErrNotImplementation
}
