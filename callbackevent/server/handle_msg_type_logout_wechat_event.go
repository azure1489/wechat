package server

import (
	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeLogoutWeChatEvent 处理退出登陆微信事件
func (srv *Server) handleMsgTypeLogoutWeChatEvent(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	//表示退出登陆微信
	wcMsgItem.EventType = message.PCLogoutWxEvent
	return nil
}
