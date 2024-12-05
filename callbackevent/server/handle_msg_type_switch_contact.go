package server

import (
	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeSwitchContact 处理切换联系人消息
func (srv *Server) handleMsgTypeSwitchContact(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	//表示切换联系人
	wcMsgItem.EventType = message.PCSwitchContactEvent
	return nil
}
