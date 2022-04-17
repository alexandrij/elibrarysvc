package elibrarysvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
)

type Endpoints struct {
	GetArticlesEndpoint   endpoint.Endpoint
	GetArticleEndpoint    endpoint.Endpoint
	PostArticleEndpoint   endpoint.Endpoint
	DeleteArticleEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{}
}

func (e Endpoints) GetArticles(ctx context.Context) ([]Article, error) {
	request := getArticlesRequest{}
	response, err := e.GetArticlesEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	resp := response.(getArticlesResponse)
	return resp.Articles, resp.Err
}

func (e Endpoints) GetArticle(ctx context.Context, id uuid.UUID) (Article, error) {
	request := getArticleRequest{Id: id}
	response, err := e.GetArticleEndpoint(ctx, request)
	if err != nil {
		return Article{}, err
	}
	resp := response.(getArticleResponse)
	return resp.Article, resp.Err
}

func (e Endpoints) PostArticle(ctx context.Context, article Article) error {
	request := postArticleRequest{Article: article}
	response, err := e.PostArticleEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(postArticleResponse)
	return resp.Err
}

func (e Endpoints) DeleteArticle(ctx context.Context, id uuid.UUID) error {
	request := deleteArticleRequest{Id: id}
	response, err := e.DeleteArticleEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(deleteArticleResponse)
	return resp.Err
}

type getArticlesRequest struct{}

type getArticlesResponse struct {
	Articles []Article `json:"articles,omitempty"`
	Err      error     `json:"err,omitempty"`
}

func (r getArticlesResponse) error() error { return r.Err }

type getArticleRequest struct {
	Id uuid.UUID
}

type getArticleResponse struct {
	Article Article `json:"article,omitempty"`
	Err     error   `json:"err,omitempty"`
}

func (r getArticleResponse) error() error { return r.Err }

type postArticleRequest struct {
	Article Article
}

type postArticleResponse struct {
	Err error `json:"err,omitempty"`
}

func (r postArticleResponse) error() error { return r.Err }

type deleteArticleRequest struct {
	Id uuid.UUID
}

type deleteArticleResponse struct {
	Err error `json:"err,omitempty"`
}

func (r deleteArticleResponse) error() error { return r.Err }
