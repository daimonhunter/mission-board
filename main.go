package main

import (
	//"mission-board/router"
	"mission-board/libs/env"
	"fmt"
)

func main() {
	//router.Init() // init router
	redis := env.GetSection("redis")
	fmt.Println(redis["REDIS_HOST"])
	mysql := env.GetSection("mysql")
	fmt.Println(mysql["REDIS_HOST"])
}
