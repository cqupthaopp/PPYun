package utils

import (
	"go.uber.org/zap"
	_ "go.uber.org/zap"
)

/**
 * @Author: Hao_pp
 * @Data: 2022-7-30 10:05
 * @Desc: ZapUtil
 */

var logger *zap.Logger

//InitZap 初始化zap对象
func InitZap() error {
	var err error
	logger, err = zap.NewDevelopment()
	return err
}

//Zap 获取zap对象
func Zap() (l *zap.Logger) {
	return logger
}
