package main

import (
	auth "github.com/DiscoFighter47/gAuth"
	config "github.com/DiscoFighter47/gConfig"
	"github.com/DiscoFighter47/todo/backend/api"
	cache "github.com/DiscoFighter47/todo/backend/cache/inmemory"
	data "github.com/DiscoFighter47/todo/backend/data/inmemory"
	"github.com/DiscoFighter47/todo/backend/server"
)

func main() {
	appCfg := config.App()
	authCfg := config.Auth()

	store := data.NewDatastore()
	cache := cache.NewCache()
	auth := auth.NewAuth(authCfg.Secret, authCfg.TokenExpireTimeout)
	auth.SetBlackListStore(cache)
	api := api.NewAPI(store, cache, auth)

	server.NewServer(api, appCfg).Serve()
}
