package server

import (
	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeSwitchChatObject 处理切换聊天对象消息
func (srv *Server) handleMsgTypeSwitchChatObject(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	//表示切换聊天对象
	wcMsgItem.EventType = message.PCSwitchChatObjectEvent
	return nil
}
