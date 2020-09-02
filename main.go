package main

import (
	"bluebell/logger"
	"bluebell/settings"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	// 1. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	fmt.Println(settings.Conf)
	fmt.Println(settings.Conf.LogConfig == nil)

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
}
