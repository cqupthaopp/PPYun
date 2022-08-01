package files

import (
	"PPYun/server/core/files/create"
	"PPYun/server/core/files/download"
	"PPYun/server/core/files/modify"
	"PPYun/server/core/files/share"
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

	{
		r.GET("/files/download", download.DownloadFile)   //普通下载
		r.GET("/share/download", download.DownloadInLink) //通过分享的链接下载
	} //下载文件

	{
		r.POST("/files/rename", modify.Rename)                      //重命名文件
		r.POST("/files/modify/path", modify.ModifyPath)             //修改路径
		r.POST("/files/modify/permission", modify.ModifyPermission) //修改权限
		r.POST("/files/remove", modify.RemoveFile)                  //删除文件
	} //Modify

	{
		r.POST("/share/link", share.GetFileLink) //分享成为链接
		r.GET("/share/qrcode", share.CreateQRCode)
	} //share

}
