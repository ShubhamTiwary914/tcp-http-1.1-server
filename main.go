package main

import (
	"fmt"
	"net"

	Router "server/router"
	HTTPSchema "server/types"
	compose "server/utils/compose"
	parser "server/utils/parser"
)

var MapRoutes = make(map[string]map[string]Router.RouteHandler)

func getHandler(header HTTPSchema.Headers, body HTTPSchema.Body) string {
	response := compose.ComposeHttpResponse("HTTP/1.1", "200", "application/json", "Request received at "+header["Host"]+" : Hello "+body["name"])
	return response
}

func socket_init() {
	//Listen TCP Socket
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println("HTTP listening on port 8080")
	defer listener.Close()

	//define routes before tcp event loop
	MapRoutes, err = Router.InitializeRoutes(MapRoutes)
	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	MapRoutes["GET"]["/"] = getHandler

	//tcp event loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error ", err)
			continue
		}
		go connection(conn)
	}

}

func connection(conn net.Conn) {
	//make buffer for request
	defer conn.Close()
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	//REQUEST Receive
	request := string(buff)
	headers, body, info := parser.ParseHttpRequest(request)

	fmt.Println(info.Version, headers["Host"], info.Method, info.Path)

	//Response from Route Handler
	if handler, ok := MapRoutes[info.Method][info.Path]; ok {
		response := handler(headers, body)
		conn.Write([]byte(response))
	}
}

func main() {
	socket_init()
}
