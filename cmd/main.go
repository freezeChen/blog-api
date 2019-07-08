/*
   @Time : 2019-06-28 13:49:25
   @Author :
   @File : main
   @Software: blog
*/
package main

import (
	"blog/conf"
	"blog/server/http"
	"blog/service"
	"fmt"
	"github.com/freezeChen/studio-library/zlog"
	"github.com/micro/go-micro/web"
	"time"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v",cfg))
	zlog.InitLogger(cfg.Log)

	svc := web.NewService(
		web.Name("go.micro.web.hello"),
		web.Address(":8081"),
		web.RegisterTTL(30*time.Second),
		web.RegisterInterval(20*time.Second))
	svc.Init()
	s := service.New(cfg)
	engine := http.New(s)
	svc.Handle("/", engine)
	if err := svc.Run(); err != nil {
		return
	}
}
