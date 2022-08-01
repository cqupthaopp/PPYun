package share

import (
	"PPYun/server/config"
	"PPYun/server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Hao_pp
 * @Data: 2022年8月1日-12点59分
 * @Desc: 生成分享链接
 */

//GetFileLink
func GetFileLink(c *gin.Context) {

	head := c.GetHeader("Authorization")
	username, ok := utils.JudgeAccessToken(head)
	if !ok {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "AccessToken Error",
		})
		return
	} //验证Token

	path := c.PostForm("path")                           //相对地址
	absPath := config.File_path + username + "\\" + path //绝对地址

	if ok, err := utils.PathExists(absPath); !ok || err != nil {
		c.JSON(401, map[string]interface{}{
			"statu": "10004",
			"info":  "Fail is not exist",
		})
		return
	} //文件不存在

	fileToken, err := utils.CreateFileToken(username, path)

	if err != nil {
		c.JSON(401, map[string]interface{}{
			"statu": "10004",
			"info":  "Create Link Failed",
		})
		return
	} //生成链接Token错误

	c.JSON(200, map[string]interface{}{
		"statu": "10000",
		"info":  "success",
		"link":  fmt.Sprintf(config.Web_Address+"/share/download?token=%s", fileToken),
	}) //返回生成的Link
	return

}
