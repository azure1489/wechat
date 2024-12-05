package server

import (
	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeVoice 处理语音消息
func (srv *Server) handleMsgTypeVoice(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	fromType := wcMsgItem.FromType
	// "msgtype":"34", PC收到语音消息
	wcMsgItem.MsgItem = message.Voice{
		VoiceLen:  msgInfo.Get("voice_len").String(),
		VoiceData: msgInfo.Get("voice_data").String(),
		VoiceHex:  msgInfo.Get("voice_hex").String(),
	}
	// fromtype: 1=个人消息,2=群消息
	if fromType == "1" {
		wcMsgItem.EventType = message.PCRecvVoiceMsgEvent
	} else if fromType == "2" {
		wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
			FromGname: msgInfo.Get("fromgname").String(), // 群名称
			FromGid:   msgInfo.Get("fromgid").String(),   // 群ID
		}
		wcMsgItem.EventType = message.PCRecvGroupVoiceMsgEvent
	}
	return nil
}
