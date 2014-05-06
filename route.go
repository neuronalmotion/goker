package gocker

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func HttpHandler() (h *rest.ResourceHandler) {
	h = &rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}

	user := User{}
	h.SetRoutes(
		// users
		rest.RouteObjectMethod("GET", "/users/", &user, "GetAll"),
		rest.RouteObjectMethod("POST", "/users/", &user, "Post"),
		rest.RouteObjectMethod("GET", "/users/:id", &user, "Get"),
		rest.RouteObjectMethod("PUT", "/users/:id", &user, "Put"),
		rest.RouteObjectMethod("DELETE", "/users/:id", &user, "Delete"),
	)
	return
}
