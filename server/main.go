package main

import (
	"PPYun/server/core"
	"PPYun/server/utils"
	"fmt"
)

func main() {

	zap_err := utils.InitZap()   //初始化Zap
	orm_err := utils.ConnectDB() //gorm连接数据库
	if zap_err != nil || orm_err != nil {
		fmt.Println("Init Panic!!!")
		fmt.Println("Zap: ", zap_err)
		fmt.Println("Gorm: ", orm_err)
		return
	}
	defer utils.CloseDB()

	red_err := utils.InitRedis() //初始化redis
	if red_err != nil {
		utils.Zap().Panic("Init Redis error")
	}
	defer utils.CloseRedis()

	core.RunServer()
}
