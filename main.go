package main

import (
	"L0/cache"
	//"L0/server"
)

func main() {
	cache.Init()
	cache.Recover()
	//subscriber.Subscribe()
	server.Start()
}
