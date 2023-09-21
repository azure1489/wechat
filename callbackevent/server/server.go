package server

import (
	"encoding/json"

	"github.com/azure1489/wechat/callbackevent/message"
)

type Server struct {
	body           []byte
	MsgList        []message.MsgItem
	messageHandler func() ([]message.MsgItem, error)
}

func NewServer(body []byte) *Server {
	srv := new(Server)
	srv.body = body

	msgBody := message.MsgBody{}
	err := json.Unmarshal(srv.body, &msgBody)
	if err != nil {
		return nil
	}

	var msgList []message.MsgItem
	// sendorrecv: 1=收到的消息,2=发送的消息
	sendorrecv := msgBody.Sendorrecv

	if msgBody.Msglist != nil && len(msgBody.Msglist) > 0 {
		msgList = msgBody.Msglist
	} else {
		msgList = append(msgList, msgBody.MsgItem)
	}

	for _, msgItem := range msgList {

		msgType := message.MsgType(msgItem.Msgtype)
		eventType := message.UnknownEvent

		switch msgType {
		case message.MsgTypeText:
			if sendorrecv == "1" {
				// PC发送文本消息
				eventType = message.PCSendTextMsgEvent
			} else if sendorrecv == "2" {
				if msgItem.Msg == "PC发文本消息成功" {
					// PC发文本消息成功
					eventType = message.PCSendTextMsgSuccessEvent
				} else {
					// PC收到文本消息
					eventType = message.PCRecvTextMsgEvent
				}
			}
		case message.MsgTypeImage:
			if sendorrecv == "1" {
				// PC发送图片消息
				eventType = message.PCSendImgMsgEvent
			} else if sendorrecv == "2" {
				if msgItem.Msg == "PC发图片消息成功" {
					// PC发图片消息成功
					eventType = message.PCSendImgMsgSuccessEvent
				} else {
					// PC收到图片消息
					eventType = message.PCRecvImgMsgEvent
				}
			}
		case message.MsgTypeFileOrAppShareLinkFile:
			if sendorrecv == "1" {
				// PC发送文件/app消息
				eventType = message.PCSendFileOrAppMsgEvent
			} else if sendorrecv == "2" {
				if msgItem.Msg == "PC发app/文件消息成功" {
					// PC发文件/app消息成功
					eventType = message.PCSendFileOrAppMsgSuccessEvent
				} else {
					// PC收到文件/app消息
					eventType = message.PCRecallMsgEvent
				}
			}
		case message.MsgTypeGif:

		}

		msgItem.EventType = eventType
		msgItem.MsgType = msgType

	}

	srv.messageHandler = func() ([]message.MsgItem, error) {

		// msgBody.Msgtype 转换为 message.MsgType

		return msgList, nil
	}
	return srv
}

// // SetMessageHandler 设置用户自定义的回调方法
// func (srv *Server) SetMessageHandler(handler func(body string)) {
// 	srv.messageHandler = handler
// }
