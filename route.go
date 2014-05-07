package gocker

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func HttpHandler() (h *rest.ResourceHandler) {
	h = &rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}

	user := User{}
	league := League{}
	h.SetRoutes(
		rest.RouteObjectMethod("GET", "/users/", &user, "GetAll"),
		rest.RouteObjectMethod("POST", "/users/", &user, "Post"),
		rest.RouteObjectMethod("GET", "/users/:id", &user, "Get"),
		rest.RouteObjectMethod("PUT", "/users/:id", &user, "Put"),
		rest.RouteObjectMethod("DELETE", "/users/:id", &user, "Delete"),
		rest.RouteObjectMethod("GET", "/users/:id/leagues/", &user, "GetLeagues"),

        rest.RouteObjectMethod("GET", "/leagues/", &league, "GetAll"),
        rest.RouteObjectMethod("GET", "/leagues/:id", &league, "Get"),
	)
	return
}
