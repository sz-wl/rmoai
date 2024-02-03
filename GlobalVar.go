package rmoai

import (
	"context"
	"sync"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
)

// 全局context.Background()
var Ctx context.Context

// 机器人全局token
var Token *token.Token

// 全局Mutex
var Mutex *sync.Mutex

// 机器人全局api
var Api openapi.OpenAPI

// 机器人全局ws
var Ws *dto.WebsocketAP

// 基础解释Message
type NormalMessage struct {
	ChannelID string
	ID        string
	Author    *dto.User
}
