package main

import (
	router "ginchat/router"
	"ginchat/utils"
)

func main() {
	//拆分成2块：GET放在service   路有地址放在router

	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run(":8081")
}
