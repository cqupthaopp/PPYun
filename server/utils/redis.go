package utils

import (
	"PPYun/server/config"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var conn redis.Conn

//InitRedis Connect Redis
func InitRedis() error {
	var err error
	conn, err = redis.Dial(config.Redis_Connect_way, config.Address+":"+config.Redis_port)
	return err
}

//WriteData 更新玩家的文件树
func WriteData(username string, datas []byte) error {
	_, err := conn.Do("Set", username, datas)
	return err
}

//GetUserData GetTheFilesTree in redis
func GetUserData(username string, datas *config.Folder) {

	ans, err := conn.Do("Get", username)
	if err != nil {
		Zap().Info(err.Error())
		return
	}
	json.Unmarshal([]byte(fmt.Sprintf("%s", ans)), &datas)

}

//CloseRedis Close conn
func CloseRedis() {
	conn.Close()
}
