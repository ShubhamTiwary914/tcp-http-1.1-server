package router

import (
	HTTPSchema "server/types"
)

type RouteHandler func(HTTPSchema.Headers, HTTPSchema.Body) string

func InitializeRoutes(MapRoutes map[string]map[string]RouteHandler) (map[string]map[string]RouteHandler, error) {
	MapRoutes["GET"] = make(map[string]RouteHandler)
	MapRoutes["POST"] = make(map[string]RouteHandler)
	MapRoutes["PUT"] = make(map[string]RouteHandler)
	MapRoutes["DELETE"] = make(map[string]RouteHandler)
	return MapRoutes, nil
}
