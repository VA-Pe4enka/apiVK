package main

import "modules/internal/server"

func main() {
	service := server.Service{}
	service.Handler()
}
