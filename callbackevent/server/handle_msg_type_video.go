package server

import (
	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeVideo 处理视频消息
func (srv *Server) handleMsgTypeVideo(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	// "msgtype":"43", PC收到视频消息
	wcMsgItem.MsgItem = message.Video{
		Info:      msgInfo.Get("info").String(), // 消息源内容
		VideoPath: msgInfo.Get("video_path").String(),
	}

	fromType := wcMsgItem.FromType

	// fromtype: 1=个人消息,2=群消息
	if fromType == "1" {
		wcMsgItem.EventType = message.PCRecvVideoMsgEvent
	} else if fromType == "2" {
		wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
			FromGname: msgInfo.Get("fromgname").String(), // 群名称
			FromGid:   msgInfo.Get("fromgid").String(),   // 群ID
		}
		wcMsgItem.EventType = message.PCRecvGroupVideoMsgEvent
	}
	return nil
}
