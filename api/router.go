package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
			Handler(route.handlerFunc)
	}

	return router
}
