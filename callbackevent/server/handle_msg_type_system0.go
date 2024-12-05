package server

import (
	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeSystem0 处理系统0消息
func (srv *Server) handleMsgTypeSystem0(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	return nil
}
