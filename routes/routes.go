package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	Auth(r)
	Film(r)
	Category(r)
	Transaction(r)
	Episode(r)
}