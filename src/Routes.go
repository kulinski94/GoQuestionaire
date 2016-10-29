package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"Get",
		"/",
		getAllQuestions,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		getAllQuestions,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		getAllQuestions,
	},
}
