package user

import (
	"PPYun/server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	username := c.PostForm("username")
	pwd := c.PostForm("password")

	if !utils.JudgeStrVaild(username, name_limit) || !utils.JudgeStrVaild(pwd, pwd_limit) {
		c.JSON(401, map[string]interface{}{
			"staus": "10005",
			"info":  "Arguments invaild",
		})
		return
	} //参数非法

	if utils.JudgeUser(username) {
		c.JSON(403, map[string]interface{}{
			"staus": "10004",
			"info":  "User Exist",
		})
		return
	}

	err := utils.CreateUser(username, pwd)
	if err != nil {
		c.JSON(403, map[string]interface{}{
			"staus": "10004",
			"info":  "Register error",
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
