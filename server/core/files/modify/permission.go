package modify

import (
	"PPYun/server/config"
	"PPYun/server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Hao_pp
 * @Data: 2022年8月1日-10点18分
 * @Desc: modify permission
 */

func ModifyPermission(c *gin.Context) {

	head := c.GetHeader("Authorization")
	username, ok := utils.JudgeAccessToken(head)
	if !ok {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "AccessToken Error",
		})
		return
	} //验证Token

	path := config.File_path + username + "\\" + c.PostForm("path")
	allCanDownload := false
	if c.PostForm("allCanDownload") == "1" {
		allCanDownload = true
	}
	downloadOnLink := false
	if c.PostForm("downloadOnLink") == "1" {
		downloadOnLink = true
	}
	onlyDonwloadByOwner := false
	if c.PostForm("onlyDownloadByOwner") == "1" {
		onlyDonwloadByOwner = true
	}
	ok, _ = utils.PathExists(path)

	if !ok {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "Arguments error",
		})
		return
	} //参数错误

	var userdata config.Folder //用户的文件树
	utils.GetUserData(username, &userdata)
	utils.ModifyFileStruct("\\"+c.PostForm("path"), &userdata, allCanDownload, downloadOnLink, onlyDonwloadByOwner)
	
	jsonAns, err := utils.MarshalFiles(&userdata) //序列化
	if err != nil {
		utils.Zap().Info(err.Error())
	}
	utils.WriteData(username, jsonAns) //更新文件树

	c.JSON(200, map[string]interface{}{
		"statu": 10000,
		"info":  "success",
	})
	return

}
