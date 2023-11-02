package main

import (
	"fmt"
	"gin_mall/conf"
	"gin_mall/loading"
	"gin_mall/routes"
)

func main() {
	loading.Loading()
	r := routes.NewRouter()
	err := r.Run(conf.Config.System.HttpPort)
	if err != nil {
		return
	} else {
		fmt.Println("success")
	}

}
