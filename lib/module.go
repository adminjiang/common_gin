package lib

import (
	"errors"
	"common_gin/common"
)

func InitModule(configPath string, modules []string) error {

	if len(configPath) < 0 {
		return errors.New("config  file is empty")
	}

	// 解析配置文件目录
	env := &common.Env{}
	// 解析配置文件目录
	if err := env.ParseConfPath(configPath); err != nil {
		return err
	}
	//初始化配置文件
	if err := common.InitViperConf(env); err != nil {
		return err
	}

	//加载Logger
	common.InitLogger()

	// 加载mysql配置并初始化实例
	if InArrayString("mysql", modules) {
		err := common.InitDBPool(env)
		if err != nil {
			return err
		}
	}
	//加载redis 配置并初始化
	if InArrayString("redis", modules) {
		err := common.InitRedis(env)
		if err != nil {
			return err
		}
	}

	return nil
}
