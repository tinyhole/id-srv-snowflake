package main

import (
	"flag"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/sirupsen/logrus"
	"github.com/tinyhole/kit"
	"github.com/tinyhole/id-srv-snowflake/handler"
	"github.com/tinyhole/id-srv-snowflake/idl/platform/id/srv/snowflake"
	"os"
	"time"
)


var (
	GitTag    = "2000.01.01.release"
	BuildTime = "2000-01-01T00:00:00+0800"
)

func main(){
	//显示版本号信息
	version := flag.Bool("v", false, "version")
	flag.Parse()
	if *version {
		fmt.Println("Git Tag: " + GitTag)
		fmt.Println("Build Time: " + BuildTime)
		os.Exit(0)
	}

	kit.LoadConfig()
	srv := micro.NewService(
		micro.Name("platform.id.srv.snowflake"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version("v1"),
		micro.Registry(consul.NewRegistry(registry.Addrs(kit.DefaultRegistryConf.Address))),
	)

	srv.Init()

	if err := snowflake.RegisterSnowFlakeHandler(srv.Server(), &handler.Handle{}); err != nil {
		logrus.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		logrus.Fatalf("service run error: %v", err)
	}
}