package elibrarysvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetArticlesEndpoint   endpoint.Endpoint
	GetArticleEndpoint    endpoint.Endpoint
	PostArticleEndpoint   endpoint.Endpoint
	DeleteArticleEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetArticlesEndpoint:   MakeGetArticlesEndpoint(s),
		GetArticleEndpoint:    MakeGetArticleEndpoint(s),
		PostArticleEndpoint:   MakePostArticleEndpoint(s),
		DeleteArticleEndpoint: MakeDeleteArticleEndpoint(s),
	}
}

//func MakeClientEndpoints(instance string) (Endpoints, error) {
//	if !strings.HasPrefix(instance, "http") {
//		instance = "http://" + instance
//	}
//	tgt, err := url.Parse(instance)
//	if err != nil {
//		return Endpoints{}, err
//	}
//	tgt.Path = ""
//
//	options := []httptransport.ClientOption{}
//
//	return Endpoints{}
//}

func (e Endpoints) GetArticles(ctx context.Context) ([]Article, error) {
	request := getArticlesRequest{}
	response, err := e.GetArticlesEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	resp := response.(getArticlesResponse)
	return resp.Articles, resp.Err
}

func (e Endpoints) GetArticle(ctx context.Context, id int) (Article, error) {
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

func (e Endpoints) DeleteArticle(ctx context.Context, id int) error {
	request := deleteArticleRequest{Id: id}
	response, err := e.DeleteArticleEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(deleteArticleResponse)
	return resp.Err
}

func MakeGetArticlesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		as, e := s.GetArticles(ctx)
		return getArticlesResponse{Articles: as, Err: e}, nil
	}
}

func MakeGetArticleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getArticleRequest)
		a, e := s.GetArticle(ctx, req.Id)
		return getArticleResponse{Article: a, Err: e}, nil
	}
}

func MakePostArticleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postArticleRequest)
		e := s.PostArticle(ctx, req.Article)
		return postArticleResponse{Err: e}, nil
	}
}

func MakeDeleteArticleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteArticleRequest)
		e := s.DeleteArticle(ctx, req.Id)
		return deleteArticleResponse{Err: e}, nil
	}
}

type getArticlesRequest struct{}

type getArticlesResponse struct {
	Articles []Article `json:"articles,omitempty"`
	Err      error     `json:"err,omitempty"`
}

func (r getArticlesResponse) error() error { return r.Err }

type getArticleRequest struct {
	Id int
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
	Id int
}

type deleteArticleResponse struct {
	Err error `json:"err,omitempty"`
}

func (r deleteArticleResponse) error() error { return r.Err }