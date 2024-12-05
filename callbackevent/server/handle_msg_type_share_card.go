package server

import (
	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeShareCard 处理名片消息
func (srv *Server) handleMsgTypeShareCard(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	// PC收到名片消息
	wcMsgItem.EventType = message.PCRecvShareCardMsgEvent
	return nil
}
