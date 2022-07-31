package files

import (
	"PPYun/server/core/files/create"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Hao_pp
 * @Data: 2022年7月31日-00点04分
 * @Desc: AboutFiles
 */

func Routers(r *gin.Engine) {

	r.POST("/files/upload", UploadFile)                 //上传文件
	r.POST("/files/create/folder", create.CreateFolder) //新建文件夹
	r.POST("/files/rename", Rename)                     //重命名文件

	{
		r.GET("/files/download", DownloadFile) //普通下载
	} //下载文件

}
