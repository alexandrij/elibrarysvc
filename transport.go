package elibrarysvc

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var (
	ErrBadRouting = errors.New("несогласованное сопоставление между маршрутом и обработчиком (ошибка программиста)")
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
	id, e := strconv.Atoi(sid)
	if e != nil {
		return nil, e
	}
	return getArticleRequest{Id: id}, nil
}
