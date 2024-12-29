package main

import httpserver "app/internal/http-server"

func main() {
	if err := httpserver.StartHttpServer(); err != nil {
		panic(err)
	}
}
