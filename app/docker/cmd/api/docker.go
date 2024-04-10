package main

import (
	"flag"
	"fmt"

	"go-zero-container/app/docker/cmd/api/internal/config"
	"go-zero-container/app/docker/cmd/api/internal/handler"
	"go-zero-container/app/docker/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "D:\\projects\\go-zero-container\\app\\docker\\cmd\\api\\etc\\docker.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
