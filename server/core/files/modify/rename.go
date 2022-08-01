package modify

import (
	"PPYun/server/config"
	"PPYun/server/utils"
	"github.com/gin-gonic/gin"
	"os"
)

//Rename rename files
func Rename(c *gin.Context) {

	head := c.GetHeader("Authorization")

	username, ok := utils.JudgeAccessToken(head)

	if !ok {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "AccessToken Error",
		})
		return
	}

	basePath := config.File_path + "\\" + username
	path := c.PostForm("path")
	filename := c.PostForm("filename")
	newName := c.PostForm("newname")

	if !utils.JudgeStrVaild(newName, 20) {
		c.JSON(401, map[string]interface{}{
			"statu": 10004,
			"info":  "Name invaild",
		})
		return
	} //判断新文件名是否合法

	ok, err := utils.PathExists(basePath + "\\" + path + "\\" + filename)
	if !ok || err != nil {
		c.JSON(401, map[string]interface{}{
			"statu": 10004,
			"info":  "Rename Failed",
		})
		return
	} //是否可以重命名

	err = os.Rename(basePath+"\\"+path+"\\"+filename, basePath+"\\"+path+"\\"+newName)
	if err != nil {
		c.JSON(401, map[string]interface{}{
			"statu": 10004,
			"info":  "Rename Failed",
		})
		return
	} //重命名失败

	c.JSON(200, map[string]interface{}{
		"statu": 10000,
		"info":  "success",
		"data":  "new name : " + newName,
	})
	utils.UpDatePathMD5(basePath+"\\"+path+"\\"+filename, basePath+"\\"+path+"\\"+newName)
	utils.UpDateUserFilesPath(username)
	return

}
