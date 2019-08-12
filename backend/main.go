package main

import (
	auth "github.com/DiscoFighter47/gAuth"
	config "github.com/DiscoFighter47/gConfig"
	"github.com/DiscoFighter47/todo/backend/api"
	"github.com/DiscoFighter47/todo/backend/data/inmemory"
	"github.com/DiscoFighter47/todo/backend/server"
)

func main() {
	appCfg := config.App()
	authCfg := config.Auth()

	store := inmemory.NewDatastore()
	auth := auth.NewAuth(authCfg.Secret, authCfg.TokenExpireTimeout)
	api := api.NewAPI(store, auth)

	server.NewServer(api, appCfg).Serve()
}
