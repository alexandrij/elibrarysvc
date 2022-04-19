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

type ArticlesRepo struct {
	mtx sync.RWMutex
	m   map[domain.ArticleID]domain.Article
}

func (s *ArticlesRepo) GetArticles(_ context.Context) ([]domain.Article, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	articles := make([]domain.Article, 0, len(s.m))
	for _, a := range s.m {
		articles = append(articles, a)
	}
	return articles, nil
}

func (s *ArticlesRepo) GetArticle(_ context.Context, id domain.ArticleID) (domain.Article, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	a, ok := s.m[id]
	if !ok {
		return domain.Article{}, ErrNotFound
	}
	return a, nil
}

func (s *ArticlesRepo) PostArticle(_ context.Context, a domain.Article) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[a.ID]; !ok {
		return ErrAlreadyExists
	}
	s.m[a.ID] = a
	return nil
}

func (s *ArticlesRepo) DeleteArticle(_ context.Context, id domain.ArticleID) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[id]; !ok {
		return ErrNotFound
	}
	delete(s.m, id)
	return nil
}
