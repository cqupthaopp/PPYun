package files

import (
	"PPYun/server/config"
	"PPYun/server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
	"time"
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
	if ok {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "AccessToken Error",
		})
		return
	} //验证Token
	username = "Hao_pp"
	path := config.File_path + username + "\\" + c.Query("path") //下载的文件的相对路径
	if ok, err := utils.PathExists(path); !ok || err != nil {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "Arguments error",
		})
		return
	} //参数错误

	md5 := utils.GetMD5ByPath(path) // 获取MD5值
	fis, err := os.ReadDir(config.MD5_path + md5)
	if err != nil || len(md5) == 0 {
		c.JSON(401, map[string]interface{}{
			"staus": "10007",
			"info":  "GetFile error",
		})
		return
	} //获取文件失败或者获取md5失败

	file, _ := os.OpenFile(config.MD5_path+md5+"\\"+fis[0].Name(), os.O_RDWR, os.ModePerm)
	path = config.MD5_path + md5 + "\\" + fis[0].Name()
	defer file.Close()

	if info, _ := file.Stat(); info.Size() <= 1024*1024*10 {

		c.Header("Content-Length", strconv.Itoa(int(info.Size())))
		c.Header("Transfer-Encoding", "true")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", info.Name()))
		c.Header("Content-Type", "application/octet-stream")
		c.File(path)
		c.JSON(200, map[string]interface{}{
			"staus": "10000",
			"info":  "Start DownLoad",
		})
		return
	} else /*小文件直接下载 小于10M */ {

		c.Header("Content-Length", strconv.Itoa(int(info.Size())))                           //给前端发送文件大小
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", info.Name())) //文件名
		c.Header("Content-Type", "application/octet-stream")

		downloadLimit := make([]byte, 1024) //每次最多传输1kb
		for {

			n, err := file.Read(downloadLimit)
			if err == io.EOF {
				break
			} //写完了就退出
			c.Writer.Write(downloadLimit[:n]) //写入字节
			time.Sleep(time.Second / 20)      //手动限速 20kb/s

		} //分块传输

		c.Writer.Flush()
		c.JSON(200, map[string]interface{}{
			"staus": "10000",
			"info":  "DownLoad Success",
		})
		return

	} //大文件处理

	return

}

//DownloadFileOnLink 通过链接下载文件
func DownloadFileOnLink() {

}
