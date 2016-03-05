package main

import (
	"apretasteconnect/handler"

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
		"GET",
		"index",
		handler.Index,
	},
	Route{
		"POST",
		"login",
		handler.Login,
	},
}
