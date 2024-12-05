package server

import (
	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeShareLocation 处理共享位置消息
func (srv *Server) handleMsgTypeShareLocation(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	// PC收到共享位置消息
	wcMsgItem.EventType = message.PCRecvShareLocationMsgEvent
	return nil
}
