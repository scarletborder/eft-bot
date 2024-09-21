package main

import (
	"eftbot/bot"
	"eftbot/config"
	"eftbot/svc"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	c := config.NewConfig()

	conf.MustLoad(*configFile, &c)
	//TODO: check config file

	ctx := svc.NewServiceContext(c)
	defer svc.SaveUinInfo(ctx.QQCli) // 保存账号信息

	bot.Subscribe(ctx)

	// setup the main stop channel
	mc := make(chan os.Signal, 2)
	signal.Notify(mc, os.Interrupt, syscall.SIGTERM)
	for {
		switch <-mc {
		case os.Interrupt, syscall.SIGTERM:
			return
		}
	}
}
