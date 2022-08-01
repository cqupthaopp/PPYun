package files

import (
	"PPYun/server/config"
	"PPYun/server/utils"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	_ "path"
	"strconv"
)

func UploadFile(c *gin.Context) {

	head := c.GetHeader("Authorization")
	username, ok := utils.JudgeAccessToken(head)
	if !ok {
		c.JSON(401, map[string]interface{}{
			"staus": "10004",
			"info":  "AccessToken Error",
		})
		return
	} //judge token

	path := c.PostForm("path")                          //上传到的相对路径
	md5 := c.PostForm("md5")                            //文件的md5值
	chunks, cserr := strconv.Atoi(c.PostForm("chunks")) //分片总数
	chunk, cerr := strconv.Atoi(c.PostForm("chunk"))    //当前分片数
	filename := c.PostForm("name")                      //上传的文件名
	file, err := c.FormFile("file")                     //上传的文件
	file.Filename = strconv.Itoa(chunk) + ".download"   //重命名分块文件

	if err != nil {
		c.JSON(403, map[string]interface{}{
			"statu": 10003,
			"info":  "Upload File Failed",
		})
		return
	} //上传分块文件失败
	if cserr != nil || cerr != nil {
		c.JSON(403, map[string]interface{}{
			"statu": 10003,
			"info":  "Arguments invaild",
		})
		return
	} //参数非法

	file_path := config.MD5_path + md5 //文件存放路径
	if ok, err := utils.PathExists(file_path); !ok || nil != err {
		os.MkdirAll(file_path, os.ModePerm)
	} //不存在或者异常就创建文件夹

	flag := make([]bool, chunks+1) //记录已上传的切片数
	count := 0                     //成功上传的数量
	fileSons, err := ioutil.ReadDir(config.MD5_path + md5)
	var sizeSum int64 //文件大小
	sizeSum = 0

	if err == nil {
		for _, fi := range fileSons {
			if utils.GetSuffix(fi.Name()) == "download" {
				sizeSum += fi.Size()
				now, err := strconv.Atoi(fi.Name()[:(len(fi.Name()) - 9)])
				if err != nil {
					continue
				}
				flag[now] = true
			} //文件分块
		}
	} //能读取就遍历

	for i := 1; i <= chunks; i++ {
		if flag[i] {
			count++
		}
	} //记录已上传成功的分块数

	if count == chunks {
		nowPath := config.File_path + username + "\\" + path + "\\" + filename
		if ok, _ := utils.PathExists(nowPath); !ok {
			os.Create(nowPath)
		} else {
			c.JSON(401, map[string]interface{}{
				"statu": 10004,
				"info":  "File Exist",
			})
			return
		}
		utils.SetPathMD5(config.File_path+username+"\\"+path+"\\"+filename, md5)
		utils.UpDateUserFilesPath(username)
		//全部上传完成就更新用户的文件树
		c.JSON(200, map[string]interface{}{
			"statu": 10000,
			"info":  "success",
			"data":  "100.00%",
		})
		utils.MD5CountAdd(md5) //引用量+1
		return
	} //已存在相同文件，无需上传

	if !flag[chunk] {
		err := c.SaveUploadedFile(file, config.MD5_path+md5+"\\"+file.Filename)
		if err == nil {
			flag[chunk] = true
			count++
			sizeSum += file.Size
		} else {
			c.JSON(401, map[string]interface{}{
				"statu": 10004,
				"info":  "Upload File Failed",
			}) //上传失败
			return
		}
	} else {
		c.JSON(401, map[string]interface{}{
			"statu": 10004,
			"info":  "Upload File Failed",
		})
		return
	} //不需要保存文件

	if count == chunks {
		nowPath := config.File_path + username + "\\" + path + "\\" + filename
		if ok, _ := utils.PathExists(nowPath); !ok {
			os.Create(nowPath)
			utils.SetPathMD5(config.File_path+username+"\\"+path+"\\"+filename, md5)
		}
		utils.UpDateUserFilesPath(username)
	} //全部上传完成就更新用户的文件树

	{
		os.Create(config.MD5_path + md5 + "\\" + filename)
		sum, err := os.OpenFile(config.MD5_path+md5+"\\"+filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			c.JSON(401, map[string]interface{}{
				"statu": 10004,
				"info":  "Saved File Faild",
			})
			return
		} //合并文件失败
		writer := bufio.NewWriter(sum)
		for fiNum := 1; fiNum <= chunks; fiNum++ {
			now, err := os.ReadFile(config.MD5_path + md5 + "\\" + strconv.Itoa(fiNum) + ".download")
			if err != nil {
				c.JSON(401, map[string]interface{}{
					"statu": 10004,
					"info":  "Saved File Faild",
				})
				return
			} //合并文件失败
			writer.Write(now)
			writer.Flush()
		} //遍历文件序号

		for fiNum := 1; fiNum <= chunks; fiNum++ {
			os.Remove(config.MD5_path + md5 + "\\" + strconv.Itoa(fiNum) + ".download") // 删除分块文件
		}

	} //合并文件

	c.JSON(200, map[string]interface{}{
		"statu": 10000,
		"info":  "success",
		"data":  fmt.Sprintf("%.2f%v", 100*float64(1.00*count)/(float64(chunks)), "%"),
	})
	utils.AddMD5InDB(md5, sizeSum) //第一次存文件，引入引用
	return

}
