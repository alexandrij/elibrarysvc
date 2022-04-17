package elibrarysvc

import (
	"errors"
)

var (
	ErrBadRouting = errors.New("несогласованное сопоставление между маршрутом и обработчиком (ошибка программиста)")
)

//func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {}
