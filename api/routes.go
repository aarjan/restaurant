package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	handlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"vat",
		"GET",
		"/api/v1/vats",
		VATGet,
	},
	Route{
		"index",
		"GET",
		"/",
		index,
	},
}
