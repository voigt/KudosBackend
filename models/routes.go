package models

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetAllKudos",
		"GET",
		"/",
		GetAllKudos,
	},
	Route{
		"GetKudos",
		"GET",
		"/{url:blog.christophvoigt.com/[a-z|A-Z|-]+}",
		GetKudos,
	},
	Route{
		"PostKudos",
		"POST",
		"/{url:blog.christophvoigt.com/[a-z|A-Z|-]+}",
		PostKudos,
	},
}
