package server

import (
	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeLoginWeChatEvent 处理登陆微信事件
func (srv *Server) handleMsgTypeLoginWeChatEvent(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	//表示登陆微信事件
	wcMsgItem.EventType = message.PCLoginWxEvent
	return nil
}
