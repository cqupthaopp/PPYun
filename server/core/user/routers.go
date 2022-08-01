package user

import "github.com/gin-gonic/gin"

/**
 * @Author: Hao_pp
 * @Data: 2022-7-30 10:05
 * @Desc: Users
 */

func Routers(r *gin.Engine) {

	r.POST("/user/login", Login)  //UserLogin
	r.POST("/user/reg", Register) //UserRegister

}
