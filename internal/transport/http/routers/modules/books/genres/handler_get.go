package genres

import "net/http"

func (g *genresRouter) GetGenresByID() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//TODO implement me
		panic("implement)")
	}
}

func (g *genresRouter) ListAll() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//TODO implement me
		panic("implement)")
	}
}

func (g *genresRouter) FindByFilter() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//TODO implement me
		panic("implement)")
	}
}
