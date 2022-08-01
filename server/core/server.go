package core

import (
	"PPYun/server/core/files"
	"PPYun/server/core/user"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Hao_pp
 * @Data: 2022-7-30 10:05
 * @Desc: MainServer
 */

func RunServer() {

	r := gin.Default()

	user.Routers(r)  //users
	files.Routers(r) //Files operator

	r.Run(":80")

}
