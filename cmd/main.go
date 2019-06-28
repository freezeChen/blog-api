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
	"github.com/freezeChen/studio-library/zlog"
	"github.com/micro/go-web"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}

	zlog.InitLogger(cfg.Log)

	svc := web.NewService(
		web.Name("go.micro.web.hello"),
		web.Address(":8080"))
	svc.Init()
	//helloService := proto.NewHelloService("go.micro.srv.hello", client.DefaultClient)
	s := service.New(cfg)
	//s.HelloService = helloService
	engine := http.New(s)
	svc.Handle("/", engine)
	if err := svc.Run(); err != nil {
		return
	}
}
