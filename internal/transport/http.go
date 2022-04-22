package articles

import (
	"bytes"
	"context"
	"elibrarysvc/internal/domain"
	"elibrarysvc/internal/endpoint"
	service "elibrarysvc/internal/service"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func MakeHTTPHandler(s service.Services, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := endpoint.MakeServerEndpoints(s)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/articles/").Handler(httptransport.NewServer(
		e.Articles.GetArticlesEndpoint,
		decodeGetArticlesRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/articles/{id}").Handler(httptransport.NewServer(
		e.Articles.GetArticleEndpoint,
		decodeGetArticleRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/articles/").Handler(httptransport.NewServer(
		e.Articles.PostArticleEndpoint,
		decodePostArticleRequest,
		encodeResponse,
	))
	r.Methods("DELETE").Path("/articles/").Handler(httptransport.NewServer(
		e.Articles.DeleteArticleEndpoint,
		decodeDeleteArticleRequest,
		encodeResponse,
	))
	return r
}

func decodeGetArticlesRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return endpoint.GetArticlesRequest{}, nil
}

func decodeGetArticleRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	sid, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		return nil, err
	}
	return endpoint.GetArticleRequest{ID: domain.ArticleID(id)}, nil
}

func decodePostArticleRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	sid, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		return nil, err
	}
	var article domain.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		return nil, err
	}
	return endpoint.PostArticleRequest{ID: domain.ArticleID(id), Article: article}, nil
}

func decodeDeleteArticleRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	sid, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		return nil, err
	}
	return endpoint.DeleteArticleRequest{ID: domain.ArticleID(id)}, nil
}

func encodeGetArticlesRequest(ctx context.Context, req *http.Request, request interface{}) error {
	req.URL.Path = "/articles"
	return encodeRequest(ctx, req, request)
}

func encodeGetArticleRequest(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(endpoint.GetArticleRequest)
	req.URL.Path = "/articles/" + string(r.ID)
	return encodeRequest(ctx, req, request)
}

func encodePostArticleRequest(ctx context.Context, req *http.Request, request interface{}) error {
	req.URL.Path = "/articles/"
	return encodeRequest(ctx, req, request)
}

func encodeDeleteArticleRequest(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(endpoint.DeleteArticleRequest)
	req.URL.Path = "/articles/" + string(r.ID)
	return encodeRequest(ctx, req, request)
}

func decodeGetArticlesResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response endpoint.GetArticlesResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeGetArticleResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response endpoint.GetArticleResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodePostArticleResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response endpoint.PostArticleResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeDeleteArticleResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response endpoint.DeleteArticleResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrAlreadyExists, domain.ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}

}
