package main

import (
	"apretaste/handler"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Method  string
	Pattern string
	Handle  httprouter.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"POST",
		"login",
		handler.Login,
	},
}
