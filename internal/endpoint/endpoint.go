package endpoint

import "elibrarysvc/internal/service"

type Endpoints struct {
	Articles ArticlesEndpoints
}

func MakeServerEndpoints(s service.Services) Endpoints {
	return Endpoints{
		Articles: MakeArticlesServerEndpoints(s.Articles),
	}
}
