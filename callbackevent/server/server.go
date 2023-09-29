package server

import (
	"encoding/json"
	"errors"

	"github.com/azure1489/wechat/callbackevent/message"
)

type Server struct {

	// Request *http.Request

	RequestRawMsg  []byte
	RequestMsg     *message.MsgBody
	messageHandler func(*message.MsgBody)
}

func NewServer(body []byte) *Server {
	srv := new(Server)
	srv.RequestRawMsg = body
	// msgBody := message.MsgBody{}
	// err := json.Unmarshal(srv.RequestRawMsg, &msgBody)
	// if err != nil {
	// 	return nil
	// }

	// var msgList []message.CommonMsg

	// if msgBody.Msglist != nil && len(msgBody.Msglist) > 0 {
	// 	msgList = msgBody.Msglist
	// } else {
	// 	msgList = append(msgList, msgBody.CommonMsg)
	// }

	// // sendorrecv: 1=收到的消息,2=发送的消息
	// sendorrecv := msgBody.Sendorrecv
	// if sendorrecv == "1" {
	// 	// 收到的消息
	// 	for _, msgItem := range msgList {

	// 		msgType := message.MsgType(msgItem.MsgType)
	// 		eventType := message.UnknownEvent

	// 		switch msgType {
	// 		case message.MsgTypeText:
	// 			// PC发送文本消息
	// 			eventType = message.PCSendTextMsgEvent
	// 		case message.MsgTypeImage:
	// 			// PC发送图片消息
	// 			eventType = message.PCSendImgMsgEvent
	// 		case message.MsgTypeFileOrAppShareLinkFile:
	// 			// PC发送文件/app消息
	// 			eventType = message.PCSendFileOrAppMsgEvent
	// 		case message.MsgTypeGif:
	// 			// PC发动态图片消息成功
	// 			eventType = message.PCSendGifImgMsgEvent
	// 		}

	// 	}
	// } else if sendorrecv == "2" {
	// 	// 发送的消息

	// 	for _, msgItem := range msgList {

	// 		msgType := message.MsgType(msgItem.Msgtype)
	// 		eventType := message.UnknownEvent

	// 		switch msgType {
	// 		case message.MsgTypeText:
	// 			if msgItem.Msg == "PC发文本消息成功" {
	// 				// PC发文本消息成功
	// 				eventType = message.PCSendTextMsgSuccessEvent
	// 			} else {
	// 				// PC收到文本消息
	// 				eventType = message.PCRecvTextMsgEvent
	// 			}
	// 		case message.MsgTypeImage:
	// 			if msgItem.Msg == "PC发图片消息成功" {
	// 				// PC发图片消息成功
	// 				eventType = message.PCSendImgMsgSuccessEvent
	// 			} else {
	// 				// PC收到图片消息
	// 				eventType = message.PCRecvImgMsgEvent
	// 			}
	// 		case message.MsgTypeFileOrAppShareLinkFile:
	// 			if msgItem.Msg == "PC发app/文件消息成功" {
	// 				// PC发文件/app消息成功
	// 				eventType = message.PCSendFileOrAppMsgSuccessEvent
	// 			} else {
	// 				// PC收到文件/app消息
	// 				eventType = message.PCRecvFileOrAppMsgEvent
	// 			}
	// 		case message.MsgTypeGif: // "msgtype":"47",
	// 			if msgItem.Msg == "PC发动态图片消息成功" {
	// 				// PC发动态图片消息成功
	// 				eventType = message.PCSendGifImgMsgSuccessEvent
	// 			} else {
	// 				// PC收到动态图片消息
	// 				eventType = message.PCRecvGifImgMsgEvent
	// 			}
	// 		case message.MsgTypeVoice: // "msgtype":"34",
	// 			// PC收到语音消息
	// 			eventType = message.PCRecvVoiceMsgEvent
	// 		case message.MsgTypeShareCard: // "msgtype":"42",
	// 			// PC收到名片消息
	// 			eventType = message.PCRecvShareCardMsgEvent
	// 		case message.MsgTypeLocation: // "msgtype":"48",
	// 			// PC收到位置消息
	// 			eventType = message.PCRecvLocationMsgEvent
	// 		case message.MsgTypeShareLocation: // "msgtype":"58",
	// 			// PC收到共享位置消息
	// 			eventType = message.PCRecvShareLocationMsgEvent
	// 		case message.MsgTypeSystem: // "msgtype":"1000", 系统消息
	// 			if msgItem.Msg == "位置共享结束" {
	// 				// PC收到共享位置结束消息
	// 				eventType = message.PCRecvEndShareLocationMsgEvent
	// 			}

	// 		}

	// 		msgItem.EventType = eventType
	// 		msgItem.MsgType = msgType

	// 	}
	// }

	return srv
}

// Serve 处理微信的请求消息
func (srv *Server) Serve() error {
	return srv.handleRequest()
}

// SetMessageHandler 设置用户自定义的回调方法
func (srv *Server) SetMessageHandler(handler func(*message.MsgBody)) {
	srv.messageHandler = handler
}

// getMessage 解析微信返回的消息
func (srv *Server) getMessage() (interface{}, error) {
	msg := &message.MsgBody{}
	// parse json
	err := json.Unmarshal(srv.RequestRawMsg, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

// HandleRequest 处理微信的请求
func (srv *Server) handleRequest() error {

	var msg interface{}
	msg, err := srv.getMessage()
	if err != nil {
		return err
	}
	msgBody, success := msg.(*message.MsgBody)
	if !success {
		err = errors.New("消息类型转换失败")
		return err
	}
	srv.RequestMsg = msgBody
	srv.messageHandler(msgBody)

	return nil
}
