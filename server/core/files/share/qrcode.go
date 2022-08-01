package share

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

/**
 * @Author: Hao_pp
 * @Data: 2022年8月1日-14点58分
 * @Desc: 二维码
 */

func CreateQRCode(c *gin.Context) {

	url := c.Query("url")
	png, err := qrcode.Encode(url, qrcode.Medium, 256)

	fmt.Println(err)

	if err != nil {
		c.JSON(401, map[string]interface{}{
			"statu": "10003",
			"info":  "failed",
		})
		return
	} //生成二维码失败

	c.JSON(200, map[string]interface{}{
		"statu": "10000",
		"info":  "success",
		"data":  png,
	})
	//qrcode.WriteFile(url, qrcode.Medium, 256, "./golang_qrcode.png")
	return

}
