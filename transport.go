package elibrarysvc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func decodeGetArticlesRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return getArticlesRequest{}, nil
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
	return getArticleRequest{ID: id}, nil
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
	var article Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		return nil, err
	}
	return postArticleRequest{ID: id, Article: article}, nil
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
	return deleteArticleRequest{ID: id}, nil
}

func encodeGetArticlesRequest(ctx context.Context, req *http.Request, request interface{}) error {
	req.URL.Path = "/articles"
	return encodeRequest(ctx, req, request)
}

func encodeGetArticleRequest(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(getArticleRequest)
	req.URL.Path = "/articles/" + string(r.ID)
	return encodeRequest(ctx, req, request)
}

func encodePostArticleRequest(ctx context.Context, req *http.Request, request interface{}) error {
	req.URL.Path = "/articles/"
	return encodeRequest(ctx, req, request)
}

func encodeDeleteArticleRequest(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(deleteArticleRequest)
	req.URL.Path = "/articles/" + string(r.ID)
	return encodeRequest(ctx, req, request)
}

func decodeGetArticlesResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response getArticlesResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeGetArticleResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response getArticleResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodePostArticleResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response postArticleResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeDeleteArticleResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response deleteArticleResponse
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
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}

}
