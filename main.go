package main

import (
	"fmt"
	"common_gin/common"
	"common_gin/lib"
	"common_gin/routes"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("初始化项目")
	err := lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	if err != nil {
		fmt.Println(err)
	}
	defer common.CloseDB()

	//启动服务
	routes.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	//关闭服务
	routes.HttpServerStop()
}
