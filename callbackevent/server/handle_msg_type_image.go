package server

import (
	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeImage 处理图片消息
func (srv *Server) handleMsgTypeImage(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	fromType := wcMsgItem.FromType
	// PC收到图片消息
	image := message.Image{
		Info:      msgInfo.Get("info").String(),                     // 消息源内容
		ImgLen:    stringToFloat64(msgInfo.Get("img_len").String()), // 消息源内容
		ImgPath:   msgInfo.Get("img_path").String(),                 // 消息源内容
		ImgBase64: msgInfo.Get("img_base64").String(),               // 消息源内容
	}

	wcMsgItem.MsgItem = image

	// fromtype: 1=个人消息,2=群消息
	if fromType == "1" {
		wcMsgItem.EventType = message.PCRecvImgMsgEvent
	} else if fromType == "2" {
		wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
			FromGname: msgInfo.Get("fromgname").String(), // 群名称
			FromGid:   msgInfo.Get("fromgid").String(),   // 群ID
		}
		wcMsgItem.EventType = message.PCRecvGroupImgMsgEvent
	}
	return nil
}
