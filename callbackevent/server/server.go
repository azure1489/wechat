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
	messageHandler func(*message.MsgBody) error
}

func NewServer(body []byte) *Server {
	srv := new(Server)
	srv.RequestRawMsg = body
	// msgBody := message.MsgBody{}
	// err := json.Unmarshal(srv.RequestRawMsg, &msgBody)
	// if err != nil {
	// 	return nil
	// }

	return srv
}

// Serve 处理微信的请求消息
func (srv *Server) Serve() error {
	return srv.handleRequest()
}

// SetMessageHandler 设置用户自定义的回调方法
func (srv *Server) SetMessageHandler(handler func(*message.MsgBody) error) {
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

	return srv.messageHandler(msgBody)
}
