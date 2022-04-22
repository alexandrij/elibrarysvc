package endpoint

import (
	"context"
	"elibrarysvc/internal/domain"
	"elibrarysvc/internal/service"
	"github.com/go-kit/kit/endpoint"
)

type ArticlesEndpoints struct {
	GetArticlesEndpoint   endpoint.Endpoint
	GetArticleEndpoint    endpoint.Endpoint
	PostArticleEndpoint   endpoint.Endpoint
	DeleteArticleEndpoint endpoint.Endpoint
}

func MakeArticlesServerEndpoints(s service.Articles) ArticlesEndpoints {
	return ArticlesEndpoints{
		GetArticlesEndpoint:   MakeGetArticlesEndpoint(s),
		GetArticleEndpoint:    MakeGetArticleEndpoint(s),
		PostArticleEndpoint:   MakePostArticleEndpoint(s),
		DeleteArticleEndpoint: MakeDeleteArticleEndpoint(s),
	}
}

func MakeGetArticlesEndpoint(s service.Articles) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		as, e := s.GetArticles(ctx)
		return GetArticlesResponse{Articles: as, Err: e}, nil
	}
}

func MakeGetArticleEndpoint(s service.Articles) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetArticleRequest)
		a, e := s.GetArticle(ctx, req.ID)
		return GetArticleResponse{Article: a, Err: e}, nil
	}
}

func MakePostArticleEndpoint(s service.Articles) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PostArticleRequest)
		e := s.PostArticle(ctx, req.Article)
		return PostArticleResponse{Err: e}, nil
	}
}

func MakeDeleteArticleEndpoint(s service.Articles) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteArticleRequest)
		e := s.DeleteArticle(ctx, req.ID)
		return DeleteArticleResponse{Err: e}, nil
	}
}

type GetArticlesRequest struct{}

type GetArticlesResponse struct {
	Articles []domain.Article `json:"articles,omitempty"`
	Err      error            `json:"err,omitempty"`
}

func (r GetArticlesResponse) error() error { return r.Err }

type GetArticleRequest struct {
	ID domain.ArticleID
}

type GetArticleResponse struct {
	Article domain.Article `json:"article,omitempty"`
	Err     error          `json:"err,omitempty"`
}

func (r GetArticleResponse) error() error { return r.Err }

type PostArticleRequest struct {
	ID      domain.ArticleID
	Article domain.Article
}

type PostArticleResponse struct {
	Err error `json:"err,omitempty"`
}

func (r PostArticleResponse) error() error { return r.Err }

type DeleteArticleRequest struct {
	ID domain.ArticleID
}

type DeleteArticleResponse struct {
	Err error `json:"err,omitempty"`
}

func (r DeleteArticleResponse) error() error { return r.Err }
