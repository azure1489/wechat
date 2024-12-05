package server

import (
	"encoding/xml"
	"fmt"

	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeSystem2 处理系统2消息
func (srv *Server) handleMsgTypeSystem2(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	fromType := wcMsgItem.FromType
	msgContent := wcMsgItem.Msg

	sysMsgXml := message.SysMsgXml{}
	err := xml.Unmarshal([]byte(msgContent), &sysMsgXml)
	if err != nil {
		fmt.Printf("xml.Unmarshal err: %v\n", err)
		return err
	}
	if sysMsgXml.Type == "revokemsg" {
		// 撤回消息
		revokeMsgXml := message.RevokeMsgXml{}
		err = xml.Unmarshal([]byte(msgContent), &sysMsgXml)
		if err != nil {
			return err
		}

		revoke := message.Revoke{
			RevokeMsg:  msgInfo.Get("revoke_msg").String(),
			Session:    revokeMsgXml.RevokeMsg.Session,
			MsgId:      revokeMsgXml.RevokeMsg.MsgId,
			NewMsgId:   revokeMsgXml.RevokeMsg.NewMsgId,
			ReplaceMsg: revokeMsgXml.RevokeMsg.ReplaceMsg,
		}
		wcMsgItem.MsgItem = revoke
		if fromType == "1" {
			wcMsgItem.EventType = message.PCRecvRevokeMsgEvent
		} else if fromType == "2" {
			wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
				FromGname: msgInfo.Get("fromgname").String(), // 群名称
				FromGid:   msgInfo.Get("fromgid").String(),   // 群ID
			}
			wcMsgItem.EventType = message.PCRecvGroupRevokeMsgEvent
		}
	} else if sysMsgXml.Type == "paymsg" {
		// 付款事件

		if fromType == "1" {
			wcMsgItem.EventType = message.PCRecvPayMsgEvent
		} else if fromType == "2" {
			wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
				FromGname: msgInfo.Get("fromgname").String(), // 群名称
				FromGid:   msgInfo.Get("fromgid").String(),   // 群ID
			}
			wcMsgItem.EventType = message.PCRecvGroupPayMsgEvent
		}
	}
	return nil
}
