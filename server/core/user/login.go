package user

import (
	"PPYun/server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Hao_pp
 * @Data: 2022-7-30 10:05
 * @Desc: Login
 */

const (
	name_limit = 10
	pwd_limit  = 15
) //用户名长度限制，密码长度限制

//Login user login
func Login(c *gin.Context) {

	username := c.PostForm("username")
	pwd := c.PostForm("password")

	if !utils.JudgeStrVaild(username, name_limit) || !utils.JudgeStrVaild(pwd, pwd_limit) {
		c.JSON(401, map[string]interface{}{
			"staus": "10005",
			"info":  "Arguments invaild",
		})
		return
	} //参数非法

	if !utils.JudgeUser(username) || !utils.JudgePWD(username, pwd) {
		c.JSON(403, map[string]interface{}{
			"staus": "10005",
			"info":  "User is not exist or PWD not match",
		})
		return
	}

	token, err := utils.GetToken(username)
	fmt.Println(token, err)
	if err != nil {
		c.JSON(403, map[string]interface{}{
			"staus": "10004",
			"info":  "Create Token error",
		})
		return
	}

	c.JSON(200, map[string]interface{}{
		"staus": "10000",
		"info":  "success",
		"data": map[string]string{
			"token": token,
		},
	})
	return

}
