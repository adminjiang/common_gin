package routes

import (
	"context"
	"common_gin/common"
	"log"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	r := InitApiRouter()
	HttpSrvHandler = &http.Server{
		Addr:           common.GetStringConf("base.http.addr"),
		Handler:        r,

		ReadTimeout:    time.Duration(common.GetIntConf("base.http.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(common.GetIntConf("base.http.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(common.GetIntConf("base.http.max_header_bytes")),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n",common.GetStringConf("base.http.addr"))
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", common.GetStringConf("base.http.addr"), err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
