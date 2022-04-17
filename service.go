package elibrarysvc

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"sync"
)

type Service interface {
	GetArticles(ctx context.Context) ([]Article, error)
	GetArticle(ctx context.Context, id uuid.UUID) (Article, error)
	PostArticle(ctx context.Context, a Article) error
	DeleteArticle(ctx context.Context, id uuid.UUID) error
}

type Article struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"context,omitempty"`
	Author  string    `json:"author,omitempty"`
}

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

type inmemService struct {
	mtx sync.RWMutex
	m   map[uuid.UUID]Article
}

func NewInmemService() Service {
	return &inmemService{
		m: map[uuid.UUID]Article{},
	}
}

func (s *inmemService) GetArticles(ctx context.Context) ([]Article, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	articles := make([]Article, 0, len(s.m))
	for _, a := range s.m {
		articles = append(articles, a)
	}
	return articles, nil
}

func (s *inmemService) GetArticle(ctx context.Context, id uuid.UUID) (Article, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	a, ok := s.m[id]
	if !ok {
		return Article{}, ErrNotFound
	}
	return a, nil
}

func (s *inmemService) PostArticle(ctx context.Context, a Article) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[a.Id]; !ok {
		return ErrAlreadyExists
	}
	s.m[a.Id] = a
	return nil
}

func (s *inmemService) DeleteArticle(ctx context.Context, id uuid.UUID) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[id]; !ok {
		return ErrNotFound
	}
	delete(s.m, id)
	return nil
}
