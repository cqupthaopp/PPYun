package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

/**
 * @Author: Hao_pp
 * @Data: 2022年7月31日-13点50分
 * @Desc: file——MD5
 */

//GetFileMd5 获取文件的MD5编码
func GetFileMd5(path string) string {
	pFile, err := os.Open(path)
	if err != nil {
		Zap().Info("打开文件失败，filename=%v, err=%v" + path + " " + err.Error())
		return ""
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)

	return hex.EncodeToString(md5h.Sum(nil))
}
