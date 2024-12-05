package server

import (
	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeLoginQRCodeRefreshEvent 处理登陆二维码刷新事件
func (srv *Server) handleMsgTypeLoginQRCodeRefreshEvent(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	//表示登陆二维码刷新事件
	wcMsgItem.EventType = message.PCLoginQrcodeRefreshEvent
	wcMsgItem.MsgItem = message.QRCode{
		QRCodeBase64: msgInfo.Get("QRCode_Base64").String(),
	}
	return nil
}
