package server

import (
	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeLocation 处理位置消息
func (srv *Server) handleMsgTypeLocation(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	// PC收到位置消息
	wcMsgItem.EventType = message.PCRecvLocationMsgEvent
	return nil
}
