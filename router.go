package main

import "github.com/julienschmidt/httprouter"

func NewRouter() *httprouter.Router {
	apiVersion := "v0"
	router := httprouter.New()
	for _, route := range routes {
		if route.Pattern == "index"{
			pattern := "/"
			router.Handle(route.Method, pattern, route.Handle)
			continue
		}
		pattern := "/" + apiVersion + "/" + route.Pattern
		router.Handle(route.Method, pattern, route.Handle)
	}
	return router
}
