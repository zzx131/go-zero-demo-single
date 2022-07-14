package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/config"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/handler"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/svc"
)

var configFile = flag.String("f", "/home/zzx/go-proj/GOPATH/src/go-zero-demo-single/app/usercenter/cmd/api/etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	server.PrintRoutes()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
