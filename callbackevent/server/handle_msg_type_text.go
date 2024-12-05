package server

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeText 处理文本消息
func (srv *Server) handleMsgTypeText(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {

	fromType := wcMsgItem.FromType
	selfwxid := wcMsgItem.SelfWxid

	msgSource := msgInfo.Get("msgsource").String()

	wcMsgItem.MsgItem = message.Text{
		MsgSource: msgSource, // 消息源内容
	}
	// fromtype: 1=个人消息,2=群消息
	if fromType == "1" {
		wcMsgItem.EventType = message.PCRecvTextMsgEvent

	} else if fromType == "2" {

		wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
			FromGname: msgInfo.Get("fromgname").String(), // 群名称
			FromGid:   msgInfo.Get("fromgid").String(),   // 群ID
		}

		hasAt := false
		if msgSource != "" {
			var msgSourceXml message.MsgSourceXml
			// 字符串转换为xml
			err := xml.Unmarshal([]byte(msgSource), &msgSourceXml)
			if err == nil {

				atUserList := strings.Split(msgSourceXml.AtUserList, ",")

				wcMsgItem.AtText = message.AtText{
					AtUserList: atUserList,
				}

				for _, at := range atUserList {
					if at == selfwxid {
						hasAt = true
						break
					}
				}
			} else {
				fmt.Printf("xml.Unmarshal err: %v\n", err)
			}
		}

		if hasAt {
			wcMsgItem.EventType = message.PCRecvAtTextMsgEvent
		} else {
			wcMsgItem.EventType = message.PCRecvGroupTextMsgEvent
		}
	}
	return nil
}
