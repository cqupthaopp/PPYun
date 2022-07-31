package utils

import (
	"PPYun/server/config"
	"encoding/json"
)

//MarshalFiles 将dfs出来的文件树序列化成json
func MarshalFiles(datas *config.Folder) ([]byte, error) {
	ans, err := json.Marshal(*datas)
	return ans, err
}
