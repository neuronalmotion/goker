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
	prefix := GockerCtx.Cfg.App.UrlPrefix
	h.SetRoutes(
		rest.RouteObjectMethod("GET", prefix+"/users", &user, "GetAll"),
		rest.RouteObjectMethod("POST", prefix+"/users", &user, "Post"),
		rest.RouteObjectMethod("GET", prefix+"/users/:id", &user, "Get"),
		rest.RouteObjectMethod("PUT", prefix+"/users/:id", &user, "Put"),
		rest.RouteObjectMethod("DELETE", prefix+"/users/:id", &user, "Delete"),
		rest.RouteObjectMethod("GET", prefix+"/users/:id/leagues", &user, "GetLeagues"),

		rest.RouteObjectMethod("GET", prefix+"/leagues", &league, "GetAll"),
		rest.RouteObjectMethod("GET", prefix+"/leagues/:id", &league, "Get"),
	)
	return
}
