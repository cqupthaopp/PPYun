package modify

import (
	"PPYun/server/config"
	"PPYun/server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

/**
 * @Author: Hao_pp
 * @Data: 2022年8月1日-09点38分
 * @Desc: Modify Path
 */

//ModifyPath 路径修改
func ModifyPath(c *gin.Context) {

	head := c.GetHeader("Authorization")
	username, ok := utils.JudgeAccessToken(head)
	if !ok {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "AccessToken Error",
		})
		return
	} //验证Token

	basePath := config.File_path + username
	path := c.PostForm("path")       //原相对路径
	newPath := c.PostForm("newpath") //新相对路径

	okp, perr := utils.PathExists(basePath + "\\" + path)
	oknp, nperr := utils.PathExists(basePath + "\\" + newPath)
	if !okp || oknp || perr != nil || nperr != nil {
		c.JSON(401, map[string]interface{}{
			"statu": 10004,
			"info":  "Modify Failed",
		})
		return
	} //判断path是否存在，判断新path是否有文件

	err := os.Rename(basePath+"\\"+path, basePath+"\\"+newPath)
	fmt.Println(err)

	if err != nil {
		c.JSON(401, map[string]interface{}{
			"statu": 10004,
			"info":  "Modify Failed",
		})
		return
	} //修改失败

	c.JSON(200, map[string]interface{}{
		"statu": 10000,
		"info":  "success",
		"data":  "new Path : " + newPath,
	})
	utils.UpDatePathMD5(basePath+"\\"+path, basePath+"\\"+newPath) //修改MD5储值
	utils.UpDateUserFilesPath(username)
	return

}
