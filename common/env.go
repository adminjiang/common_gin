package common

import (
	"bytes"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

type Env struct {
	ConfEnvPath string //配置文件夹
	ConfEnv     string //配置环境名 比如：dev prod test
}

// 解析配置文件目录
//
// 配置文件必须放到一个文件夹中
// 如：config=conf/dev/base.json 	ConfEnvPath=conf/dev	ConfEnv=dev
// 如：config=conf/base.json		ConfEnvPath=conf		ConfEnv=conf
func (env *Env) ParseConfPath(config string) error {
	path := strings.Split(config, "/")
	prefix := strings.Join(path[:len(path)-1], "/")
	env.ConfEnvPath = prefix
	env.ConfEnv = path[len(path)-2]
	return nil
}

//获取配置环境名
func (env *Env) GetConfEnv() string {
	return env.ConfEnv
}

func (env *Env) GetConfPath(fileName string) string {
	return env.ConfEnvPath + "/" + fileName + ".toml"
}

func (env *Env) GetConfFilePath(fileName string) string {
	return env.ConfEnvPath + "/" + fileName
}

//解析toml 文件
func (env *Env) ParseConfig(path string, conf interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	v := viper.New()
	v.SetConfigType("toml")
	err = v.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	err = v.Unmarshal(conf)
	if err != nil {
		return err
	}
	return nil
}
