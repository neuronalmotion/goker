package goker

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func HttpHandler() (h *rest.ResourceHandler) {
	h = &rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}

	user := User{}
	cup := Cup{}
	prefix := GokerCtx.Cfg.App.UrlPrefix
	h.SetRoutes(
		rest.RouteObjectMethod("GET", prefix+"/users", &user, "GetAll"),
		rest.RouteObjectMethod("POST", prefix+"/users", &user, "Post"),
		rest.RouteObjectMethod("GET", prefix+"/users/:id", &user, "Get"),
		rest.RouteObjectMethod("PUT", prefix+"/users/:id", &user, "Put"),
		rest.RouteObjectMethod("DELETE", prefix+"/users/:id", &user, "Delete"),
		rest.RouteObjectMethod("GET", prefix+"/users/:id/cups", &user, "GetCups"),

		rest.RouteObjectMethod("GET", prefix+"/cups", &cup, "GetAll"),
		rest.RouteObjectMethod("GET", prefix+"/cups/:id", &cup, "Get"),
	)
	return
}
