package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name            string
	Method          string
	Pattern         string
	HandlerFunction http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Subscribe",
		"POST",
		"/subscribe",
		Subscribe,
	},
	Route{
		"sendmsg",
		"POST",
		"/send",
		SendMsg,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunction)
	}
	return router
}
