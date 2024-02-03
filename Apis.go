package rmoai

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/tencent-connect/botgo/dto"
)

// 被动消息发送包装
//
// ChannelId: 子频道Id
//
// Message: Message，可使用BuildMessage构建
func SendMessage(ChannelId string, Message *dto.MessageToCreate) error {
	_, err := Api.PostMessage(Ctx, ChannelId, Message)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// 构建被动消息
//
// content：将发送的消息
//
// msgId：回复消息Id
func BuildMessage(content, msgId string) *dto.MessageToCreate {
	return &dto.MessageToCreate{
		Content: content,
		MsgID:   msgId,
	}
}

// 处理at消息
//
// s: AtMessage事件体中的data.Content，去除At文本，保留纯消息文本
func FatContent(s string) string {
	re, _ := regexp.Compile("> (.*?)::")
	return strings.Split(strings.Split(string(re.Find([]byte(s+"::"))), "> ")[1], "::")[0]
}

// 将事件的消息体转换为NormalMessage,目前仅支持WSMessageData，dto.WSATMessageData
func BuildNormalMessage(data interface{}) *NormalMessage {
	_, ok := data.(*dto.WSMessageData)
	if ok {
		dat := data.(*dto.WSMessageData)
		return &NormalMessage{
			ChannelID: dat.ChannelID,
			ID:        dat.ID,
			Author:    dat.Author,
		}
	} else {
		dat := data.(*dto.WSATMessageData)
		return &NormalMessage{
			ChannelID: dat.ChannelID,
			ID:        dat.ID,
			Author:    dat.Author,
		}
	}

}
