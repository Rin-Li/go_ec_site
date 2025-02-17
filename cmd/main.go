package main

import (
	"gin-mall-tmp/conf"
	"gin-mall-tmp/pkg/util"
	"gin-mall-tmp/routes"
	"log"
)

func main() {
    conf.Init()
	
	go func (){
		log.Println("Starting preheating...")
		util.PreheatSeckillProducts()
		log.Println("Preheating finished.")
	}()

	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}

