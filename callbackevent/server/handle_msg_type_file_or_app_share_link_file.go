package server

import (
	"encoding/xml"
	"fmt"

	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeFileOrAppShareLinkFile 处理文件或应用分享链接消息
func (srv *Server) handleMsgTypeFileOrAppShareLinkFile(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	msgContent := msgInfo.Get("msg").String()
	fromType := wcMsgItem.FromType

	var appMsgXml message.AppMsgXml
	// 字符串转换为xml
	err := xml.Unmarshal([]byte(msgContent), &appMsgXml)
	if err != nil {
		fmt.Printf("xml.Unmarshal err: %v\n", err)
		return nil
	}

	switch appMsgXml.AppMsg.Type {
	case "57":
		// 引用消息
		var quoteMsgXml message.QuoteMsgXml
		// 字符串转换为xml
		err = xml.Unmarshal([]byte(msgContent), &appMsgXml)
		if err != nil {
			return err
		}
		// quote :=
		wcMsgItem.MsgItem = message.Quote{
			MsgSource:    quoteMsgXml.AppMsg.ReferMsg.MsgSource, // 消息源内容
			QuoteMsg:     quoteMsgXml.AppMsg.ReferMsg.Content,   // 引用的消息内容
			QuoteMsgType: quoteMsgXml.AppMsg.ReferMsg.Type,      // 引用的消息类型
			QuoteMsgId:   quoteMsgXml.AppMsg.ReferMsg.Svrid,     // 引用的消息id
			ReplyMsg:     quoteMsgXml.AppMsg.Title,              // 回复的消息内容
		}

		// fromtype: 1=个人消息,2=群消息
		if fromType == "1" {
			wcMsgItem.EventType = message.PCRecvQuoteMsgEvent
		} else if fromType == "2" {
			wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
				FromGname: msgInfo.Get("fromgname").String(), // 群名称
				FromGid:   msgInfo.Get("fromgid").String(),   // 群ID
			}
			wcMsgItem.EventType = message.PCRecvGroupQuoteMsgEvent
		}
	case "51":
		// 视频号消息
		var channelsMsgXml message.ChannelsMsgXml
		// 字符串转换为xml
		err = xml.Unmarshal([]byte(msgContent), &channelsMsgXml)
		if err != nil {
			return err
		}

		wcMsgItem.MsgItem = message.ChannelsMsg{
			ObjectId:      channelsMsgXml.AppMsg.FinderFeed.ObjectId,
			ObjectNonceId: channelsMsgXml.AppMsg.FinderFeed.ObjectNonceId,
		}

		// fromtype: 1=个人消息,2=群消息
		if fromType == "1" {
			wcMsgItem.EventType = message.PCRecvChannelsMsgEvent
		} else if fromType == "2" {
			wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
				FromGname: msgInfo.Get("fromgname").String(), // 群名称
				FromGid:   msgInfo.Get("fromgid").String(),   // 群ID
			}
			wcMsgItem.EventType = message.PCRecvGroupChannelsMsgEvent
		}
	case "5":
		// 服务通知消息
		var serviceNoticeXml message.ServiceNoticeXml
		// 字符串转换为xml
		err = xml.Unmarshal([]byte(msgContent), &serviceNoticeXml)
		if err != nil {
			return err
		}

		serviceNotice := message.ServiceNotice{
			Title:       serviceNoticeXml.AppMsg.Title,
			Description: serviceNoticeXml.AppMsg.Des,
			AppName:     serviceNoticeXml.AppMsg.MMReader.Category.Item.Title,
			WeappPath:   serviceNoticeXml.AppMsg.MMReader.Category.Item.WeappPath,
			WeappUser:   serviceNoticeXml.AppMsg.MMReader.Publisher.Username,
		}

		wcMsgItem.MsgItem = serviceNotice
		wcMsgItem.EventType = message.PCServiceNoticeEvent
	}

	return nil
}
