package main

import (
	"github.com/DiscoFighter47/todo/backend/api"
	"github.com/DiscoFighter47/todo/backend/server"
)

func main() {
	server.NewServer(api.NewAPI()).Serve()
}
