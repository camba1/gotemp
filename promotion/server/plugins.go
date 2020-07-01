package main

import (
	"github.com/micro/go-micro/v2/config/cmd"
	redis "github.com/micro/go-plugins/store/redis/v2"
)

func initPlugins() {
	cmd.DefaultStores["redis"] = redis.NewStore
}
