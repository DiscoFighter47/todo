package main

import (
	gauth "github.com/DiscoFighter47/gAuth"
	gconfig "github.com/DiscoFighter47/gConfig"
	"github.com/DiscoFighter47/todo/backend/api"
	"github.com/DiscoFighter47/todo/backend/data/inmemory"
	"github.com/DiscoFighter47/todo/backend/server"
)

func main() {
	server.NewServer(api.NewAPI(inmemory.NewDatastore(), gauth.NewAuth(gconfig.Auth().Secret, gconfig.Auth().TokenExpireTimeout)), gconfig.App()).Serve()
}
