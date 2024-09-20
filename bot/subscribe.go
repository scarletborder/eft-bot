package bot

import (
	"eftbot/query"
	"eftbot/svc"
	"fmt"
	"strconv"
	"strings"

	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
	"github.com/zeromicro/go-zero/core/logx"
)

func Has(arr []uint32, e uint32) bool {
	for _, ele := range arr {
		if ele == e {
			return true
		}
	}
	return false
}

func Subscribe(srvCtx *svc.ServiceContext) {
	// TODO: 优雅的字符串组装
	srvCtx.QQCli.GroupMessageEvent.Subscribe(func(client *client.QQClient, event *message.GroupMessage) {
		if Has(srvCtx.C.SubscribedGroups, event.GroupUin) {
			content := event.ToString()
			if strings.HasPrefix(content, srvCtx.C.BotPrefix) {
				content = strings.TrimPrefix(content, srvCtx.C.BotPrefix)
				content = strings.TrimSpace(content)

				// TODO : 执行逻辑
				if content == "汇率" {
					res, _ := query.QueryExchange()
					text := message.NewText(res)
					_, err := client.SendGroupMessage(event.GroupUin, []message.IMessageElement{text})
					if err != nil {
						logx.Error(err)
					}
				} else if strings.HasPrefix(content, "x87") {
					euros, err := query.QueryEuro()
					if err != nil {
						res := err.Error()
						text := message.NewText("查询失败" + res)
						_, err := client.SendGroupMessage(event.GroupUin, []message.IMessageElement{text})
						if err != nil {
							logx.Error(err)
						}
					}
					content = strings.TrimPrefix(content, "x87")
					content = strings.TrimSpace(content)
					original, err := strconv.Atoi(content)
					if err != nil {
						res := err.Error()
						text := message.NewText("数字转换失败" + res)
						_, err := client.SendGroupMessage(event.GroupUin, []message.IMessageElement{text})
						if err != nil {
							logx.Error(err)
						}
					}
					after := (float64(original)) / (0.87 * float64(euros))
					res := fmt.Sprintf("x87 : %.3f", after)
					text := message.NewText(res)
					_, err = client.SendGroupMessage(event.GroupUin, []message.IMessageElement{text})
					if err != nil {
						logx.Error(err)
					}

				} else {
					res, _ := query.QueryByCn(content)
					text := message.NewText(res)
					_, err := client.SendGroupMessage(event.GroupUin, []message.IMessageElement{text})
					if err != nil {
						logx.Error(err)
					}
				}
			}
		}
	})
}
