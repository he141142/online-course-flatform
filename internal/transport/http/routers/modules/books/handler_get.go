package books

import "net/http"

func (g *bookRouter) GetByID() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//TODO implement me
		panic("implement)")
	}
}

func (g *bookRouter) ListAll() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//TODO implement me
		panic("implement)")
	}
}

func (g *bookRouter) FindByFilter() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//TODO implement me
		panic("implement)")
	}
}
