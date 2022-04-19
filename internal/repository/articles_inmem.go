package repository

import (
	"context"
	"elibrarysvc/internal/domain"
	"errors"
	"sync"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type ArticlesInmemRepo struct {
	mtx sync.RWMutex
	m   map[domain.ArticleID]domain.Article
}

func NewArticlesInmemRepo() *ArticlesInmemRepo {
	return &ArticlesInmemRepo{
		m: map[domain.ArticleID]domain.Article{},
	}
}

func (r *ArticlesInmemRepo) GetArticles(_ context.Context) ([]domain.Article, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	articles := make([]domain.Article, 0, len(r.m))
	for _, a := range r.m {
		articles = append(articles, a)
	}
	return articles, nil
}

func (r *ArticlesInmemRepo) GetArticle(_ context.Context, id domain.ArticleID) (domain.Article, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	a, ok := r.m[id]
	if !ok {
		return domain.Article{}, ErrNotFound
	}
	return a, nil
}

func (r *ArticlesInmemRepo) PostArticle(_ context.Context, a domain.Article) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, ok := r.m[a.ID]; !ok {
		return ErrAlreadyExists
	}
	r.m[a.ID] = a
	return nil
}

func (r *ArticlesInmemRepo) DeleteArticle(_ context.Context, id domain.ArticleID) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, ok := r.m[id]; !ok {
		return ErrNotFound
	}
	delete(r.m, id)
	return nil
}
