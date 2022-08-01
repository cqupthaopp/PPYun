package modify

import (
	"PPYun/server/config"
	"PPYun/server/utils"
	"github.com/gin-gonic/gin"
	"os"
)

//RemoveFile 删除文件
func RemoveFile(c *gin.Context) {

	head := c.GetHeader("Authorization")
	username, ok := utils.JudgeAccessToken(head)
	if !ok {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "AccessToken Error",
		})
		return
	} //验证Token

	path := c.PostForm("path")
	md5 := utils.GetMD5ByPath(config.File_path + username + "\\" + path)

	if len(md5) == 0 {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "Arguments error",
		})
		return
	} //文件不存在 即参数错误

	num := utils.GetMD5Count(md5)

	if num == 1 {
		os.Remove(config.MD5_path + md5) //删除文件
	} //还剩一个引用

	os.Remove(config.File_path + username + "\\" + path) //删除文件树中的文件
	utils.UpDateUserFilesPath(username)                  //更新文件树
	utils.MD5CountDel(md5)                               //减少一次引用量

	c.JSON(200, map[string]interface{}{
		"statu": "10000",
		"info":  "success",
	})

	return

}
