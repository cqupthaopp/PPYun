package create

import (
	"PPYun/server/config"
	"PPYun/server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func CreateFolder(c *gin.Context) {

	head := c.GetHeader("Authorization")

	username, ok := utils.JudgeAccessToken(head)

	if !ok {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "AccessToken Error",
		})
		return
	}

	var datas config.Folder
	utils.GetUserData(username, &datas)
	path := config.File_path + username + "\\" + c.PostForm("path") //用户的文件路径

	for k := 0; k < 10000; k++ {
		nowPath := path + "\\新建文件夹"

		if k > 0 {
			nowPath += strconv.Itoa(k)
		}
		ok, err := utils.PathExists(nowPath)

		if ok || err != nil {
			continue
		}
		err = os.MkdirAll(nowPath, os.ModePerm)
		if err != nil {
			utils.Zap().Info(err.Error())
			break
		}
		err = utils.UpDateUserFilesPath(username)
		fmt.Println(err)
		c.JSON(200, map[string]interface{}{
			"statu": "10000",
			"info":  "success",
			"path":  nowPath,
		})
		return
	} //创建

	c.JSON(401, map[string]interface{}{
		"statu": "10004",
		"info":  "Create folder error",
	})
	return

}
