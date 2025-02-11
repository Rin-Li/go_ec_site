package main

import (
	"gin-mall-tmp/conf"
	"gin-mall-tmp/routes"
)

func main() {
    conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}

