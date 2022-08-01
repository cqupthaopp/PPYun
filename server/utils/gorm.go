package utils

import (
	"PPYun/server/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
 * @Author: Hao_pp
 * @Data: 2022-7-30 10:05
 * @Desc: GormUtil
 */

var db *gorm.DB

//ConnectDB ConnectMysql
func ConnectDB() error {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@%v(%v:%v)"+"/%v?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.ConnectWay, config.Address, config.Port, config.DBname))
	db.SingularTable(true)
	return err
}

//ConnectTable Connect to table
func ConnectTable(table_name string) *gorm.DB {
	return db.Table(table_name)
}

//JudgeUser Is user exist?
func JudgeUser(username string) bool {
	user_db := ConnectTable(config.Users_table)
	var user config.User
	user_db.Where("username = ?", username).First(&user)

	return user.Username == username
}

//JudgePWD judge user's password
func JudgePWD(username string, pwd string) bool {
	user_db := ConnectTable(config.Users_table)
	var user config.User
	user_db.Where("username = ?", username).First(&user)

	return user.Password == pwd
}

//CreateUser New User in SQL
func CreateUser(username string, pwd string) error {
	user_db := ConnectTable(config.Users_table)
	result := user_db.Create(&config.User{username, pwd})
	return result.Error
}

//IsFileExist JudgeFileExist
func IsFileExist(MD5 string) bool {
	var ans config.MD5File
	md5_db := ConnectTable(config.File_MD5_table)
	md5_db.Where("md5 = ?", MD5).First(&ans)
	return len(ans.MD5) > 0
}

//MD5CountAdd MD5文件使用量自增
func MD5CountAdd(MD5 string) {
	var ans config.MD5File
	md5_db := ConnectTable(config.File_MD5_table)
	md5_db.Where("md5 = ?", MD5).First(&ans)
	md5_db.Where("md5 = ?", MD5).Update("count", ans.Count+1)
}

//MD5CountAdd MD5文件使用量自减
func MD5CountDel(MD5 string) {
	var ans config.MD5File
	md5_db := ConnectTable(config.File_MD5_table)
	md5_db.Where("md5 = ?", MD5).First(&ans)
	if ans.Count > 1 {
		md5_db.Where("md5 = ?", MD5).Update("count", ans.Count-1)
	} else {
		md5_db.Delete(&ans)
	}
}

//GetMD5Count 获取该文件的剩余使用量
func GetMD5Count(MD5 string) int {
	var ans config.MD5File
	md5_db := ConnectTable(config.File_MD5_table)
	md5_db.Where("md5 = ?", MD5).First(&ans)
	return ans.Count
}

//AddMD5InDB 将MD5添加入数据库
func AddMD5InDB(MD5 string, Size int64) {
	md5_db := ConnectTable(config.File_MD5_table)
	md5_db.Create(&config.MD5File{
		Count: 1,
		MD5:   MD5,
		Size:  int(Size),
	})
}

//CloseDB close database
func CloseDB() {
	db.Close()
}

//GetMD5ByPath 通过path查找MD5值
func GetMD5ByPath(path string) string {
	var ans config.PathMD5
	md5_db := ConnectTable(config.Path_MD5_table)
	md5_db.Where("path = ?", path).First(&ans)
	return ans.MD5
}

//SetPathMD5 写入该路径的MD5值
func SetPathMD5(path string, MD5 string) {
	md5_db := ConnectTable(config.Path_MD5_table)
	md5_db.Create(&config.PathMD5{MD5, path})
}

//GetFileSizeByMD5 通过md5查找文件大小
func GetFileSizeByMD5(md5 string) int {
	var ans config.MD5File
	md5_db := ConnectTable(config.File_MD5_table)
	md5_db.Where("md5 = ?", md5).First(&ans)
	return ans.Size
}

//UpDatePathMD5 更新path
func UpDatePathMD5(oldPath string, newPath string) {
	md5_db := ConnectTable(config.Path_MD5_table)
	md5_db.Where("path = ?", oldPath).Update("path", newPath)
}
