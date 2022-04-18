package elibrarysvc

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Service interface {
	GetArticles(ctx context.Context) ([]Article, error)
	GetArticle(ctx context.Context, id uint64) (Article, error)
	PostArticle(ctx context.Context, a Article) error
	DeleteArticle(ctx context.Context, id uint64) error
}

type Article struct {
	ID      uint64    `json:"id"`
	Alias   string    `json:"alias,omitempty"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	Href    string    `json:"href,omitempty"`
	Author  string    `json:"author,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type inmemService struct {
	mtx sync.RWMutex
	m   map[uint64]Article
}

func NewInmemService() Service {
	return &inmemService{
		m: map[uint64]Article{},
	}
}

func (s *inmemService) GetArticles(_ context.Context) ([]Article, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	articles := make([]Article, 0, len(s.m))
	for _, a := range s.m {
		articles = append(articles, a)
	}
	return articles, nil
}

func (s *inmemService) GetArticle(_ context.Context, id uint64) (Article, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	a, ok := s.m[id]
	if !ok {
		return Article{}, ErrNotFound
	}
	return a, nil
}

func (s *inmemService) PostArticle(_ context.Context, a Article) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[a.ID]; !ok {
		return ErrAlreadyExists
	}
	s.m[a.ID] = a
	return nil
}

func (s *inmemService) DeleteArticle(_ context.Context, id uint64) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[id]; !ok {
		return ErrNotFound
	}
	delete(s.m, id)
	return nil
}
