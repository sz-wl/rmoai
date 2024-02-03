package rmoai

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/token"
)

// 此函数初始化rmoai的一切前置，在使用任何函数前，必须调用该函数
//
// appid：填入qq机器人的appid
//
// tokenl：填入qq机器人的token
func Init(appid uint64, tokenl string) {
	Mutex = new(sync.Mutex)
	Ctx = context.Background()
	Token = token.BotToken(appid, tokenl)
	Api = botgo.NewOpenAPI(Token)
	var Wserr error
	Ws, Wserr = Api.WS(Ctx, nil, "")
	if Wserr != nil {
		log.Printf("%+v, err:%v", Ws, Wserr)
	}
	fmt.Printf("\n>>>>>>>>>>>>>>>>>>>>\nBy: 侍者\nDevlop: golang\n<<<<<<<<<<<<<<<<<<<<\n")
}

// 此函数将会启动rmoai，在进行全部工作后，运行函数以启动rmoai
//
// handlers：注册事件
func Running(handlers ...interface{}) {
	inter := event.RegisterHandlers(handlers...)
	SeManager := botgo.NewSessionManager()
	SeManager.Start(Ws, Token, &inter)

}
