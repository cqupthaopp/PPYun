package files

import (
	"PPYun/server/config"
	"PPYun/server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
)

/**
 * @Author: Hao_pp
 * @Data: 2022年7月31日-09点36分
 * @Desc: DownloadFile
 */

//DownloadFile 自己下载自己的文件，即一定有权限
func DownloadFile(c *gin.Context) {

	head := c.GetHeader("Authorization")
	username, ok := utils.JudgeAccessToken(head)
	if !ok {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "AccessToken Error",
		})
		return
	} //验证Token

	path := config.File_path + username + "\\" + c.Query("path") //下载的文件的相对路径
	uploadSize, sErr := strconv.Atoi(c.Query("size"))            //已发送的字节数
	if ok, err := utils.PathExists(path); !ok || err != nil || sErr != nil {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "Arguments error",
		})
		return
	} //参数错误

	file, _ := os.Open(path)
	defer file.Close()

	file.Seek(int64(uploadSize), io.SeekStart) //偏移到未下载的字节
	info, _ := file.Stat()
	size := info.Size()
	sliceSize := 1                      //将文件分块，每块尽量小，这边给1
	sliceNum := size / int64(sliceSize) //块数
	utils.Zap().Info(fmt.Sprintf("%s 开始发送文件 %s ,大小 %v,分片数 %v", username, file.Name(), size, sliceNum))

	for {

	}

	return

}

//DownloadFileOnLink 通过链接下载文件
func DownloadFileOnLink() {

}
