package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

type Server struct {

	// Request *http.Request

	RequestRawMsg []byte
	RequestMsg    *message.MsgBody
	// WcMsgBody      *message.WcMsgBody
	Ctx context.Context
	// EventType message.EventType
	// CommonMsg 子类为入参
	messageHandler func(msgList []message.WcMsgItem) error
}

func NewServer(body []byte) *Server {
	srv := new(Server)
	srv.RequestRawMsg = body
	// srv.Ctx = ctx
	// srv.EventType = message.UnknownEvent
	// log.Println(" ---------------- 收到消息开始 ---------------- ")
	// log.Println("收到消息内容:\n", string(body))
	// log.Println(" ---------------- 收到消息结束 ---------------- ")
	return srv
}

// Serve 处理微信的请求消息
func (srv *Server) Serve() error {
	return srv.handleRequest()
}

// SetMessageHandler 设置用户自定义的回调方法
func (srv *Server) SetMessageHandler(handler func([]message.WcMsgItem) error) {
	srv.messageHandler = handler
}

// interface{} 转 sring
func stringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return f
}

// handleRequest 处理微信的请求
func (srv *Server) handleRequest() error {
	jsonText := string(srv.RequestRawMsg)
	var msgItemList []message.WcMsgItem

	// 获取基础信息
	selfwxid := gjson.Get(jsonText, "selfwxid").String()
	serverPort := gjson.Get(jsonText, "ServerPort").String()

	// 处理消息列表
	msgList := gjson.Get(jsonText, "msglist")
	for _, msgInfo := range msgList.Array() {
		wcMsgItem, err := srv.processMessage(msgInfo, selfwxid, serverPort)
		if err != nil {
			// 记录错误但继续处理其他消息
			fmt.Printf("Error processing message: %v\n", err)
			continue
		}
		msgItemList = append(msgItemList, wcMsgItem)
	}

	return srv.messageHandler(msgItemList)
}

// processMessage 处理单条消息
func (srv *Server) processMessage(msgInfo gjson.Result, selfwxid, serverPort string) (message.WcMsgItem, error) {
	// 提取基础消息字段
	wcMsgItem := srv.extractBasicMessageFields(msgInfo, selfwxid, serverPort)

	// 根据消息类型处理具体内容
	msgType := message.MsgType(msgInfo.Get("msgtype").String())
	if err := srv.handleMessageByType(msgType, &wcMsgItem, msgInfo); err != nil {
		return wcMsgItem, fmt.Errorf("handle message type %s: %w", msgType, err)
	}

	return wcMsgItem, nil
}

// extractBasicMessageFields 提取消息基础字段
func (srv *Server) extractBasicMessageFields(msgInfo gjson.Result, selfwxid, serverPort string) message.WcMsgItem {
	return message.WcMsgItem{
		ServerPort: serverPort,
		SelfWxid:   selfwxid,
		CommonMsg: message.CommonMsg{
			MsgSvrid: msgInfo.Get("msgsvrid").String(),
			Time:     msgInfo.Get("time").String(),
			Msg:      msgInfo.Get("msg").String(),
			MsgType:  msgInfo.Get("msgtype").String(),
			FromType: msgInfo.Get("fromtype").String(),
			FromId:   msgInfo.Get("fromid").String(),
			FromName: msgInfo.Get("fromname").String(),
			Index:    msgInfo.Get("index").String(),
		},
		ToCommonMsg: message.ToCommonMsg{
			ToId:   msgInfo.Get("toid").String(),
			ToName: msgInfo.Get("toname").String(),
		},
		EventType: message.UnknownEvent,
	}
}

// handleMessageByType 根据消息类型处理消息
func (srv *Server) handleMessageByType(msgType message.MsgType, wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {

	handlers := map[message.MsgType]func(*message.WcMsgItem, gjson.Result) error{
		message.MsgTypeText:                    srv.handleMsgTypeText,                    // handle_msg_type_text.go
		message.MsgTypeImage:                   srv.handleMsgTypeImage,                   // handle_msg_type_image.go
		message.MsgTypeFileOrAppShareLinkFile:  srv.handleMsgTypeFileOrAppShareLinkFile,  // handle_msg_type_file_or_app_share_link_file.go
		message.MsgTypeGif:                     srv.handleMsgTypeGif,                     // handle_msg_type_gif.go
		message.MsgTypeFriendConfirmation:      srv.handleMsgTypeFriendConfirmation,      // handle_msg_type_friend_confirmation.go
		message.MsgTypeVideo:                   srv.handleMsgTypeVideo,                   // handle_msg_type_video.go
		message.MsgTypeVoice:                   srv.handleMsgTypeVoice,                   // handle_msg_type_voice.go
		message.MsgTypeShareCard:               srv.handleMsgTypeShareCard,               // handle_msg_type_share_card.go
		message.MsgTypeLocation:                srv.handleMsgTypeLocation,                // handle_msg_type_location.go
		message.MsgTypeShareLocation:           srv.handleMsgTypeShareLocation,           // handle_msg_type_share_location.go
		message.MsgTypeSystem0:                 srv.handleMsgTypeSystem0,                 // handle_msg_type_system0.go
		message.MsgTypeSystem2:                 srv.handleMsgTypeSystem2,                 // handle_msg_type_system2.go
		message.MsgTypeLoginQRCodeRefreshEvent: srv.handleMsgTypeLoginQRCodeRefreshEvent, // handle_msg_type_login_qr_code_refresh_event.go
		message.MsgTypeLoginWeChatEvent:        srv.handleMsgTypeLoginWeChatEvent,        // handle_msg_type_login_wechat_event.go
		message.MsgTypeLogoutWeChatEvent:       srv.handleMsgTypeLogoutWeChatEvent,       // handle_msg_type_logout_wechat_event.go
		message.MsgTypeSwitchChatObject:        srv.handleMsgTypeSwitchChatObject,        // handle_msg_type_switch_chat_object.go
		message.MsgTypeSwitchContact:           srv.handleMsgTypeSwitchContact,           // handle_msg_type_switch_contact.go
	}

	handler, exists := handlers[msgType]
	if !exists {
		return nil
	}

	return handler(wcMsgItem, msgInfo)
}

// getMessage 解析微信返回的消息
// func (srv *Server) getMessage() (interface{}, error) {
// 	msg := &message.MsgBody{}
// 	// parse json
// 	err := json.Unmarshal(srv.RequestRawMsg, msg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return msg, nil
// }

// 判断是否是群消息
// func (srv *Server) isGroupMsg(username string) bool {
// 	return strings.HasSuffix(username, "@chatroom")
// }

// interface{} 转 sring
// func interfaceToString(i interface{}) string {
// 	if i == nil {
// 		return ""
// 	}
// 	return i.(string)
// }

// func (srv *Server) getFromType(result gjson.Result) string {
// 	// fromtype: 1=个人消息,2=群消息
// 	fromtype := result.Get("fromtype")
// 	fromType := ""
// 	if fromtype.Exists() {
// 		fromType = fromtype.String()
// 	}
// 	return fromType
// }

// func (srv *Server) SetEventType(msgInfo gjson.Result) {

// }

// handleMessage 处理单个消息
// func (s *Server) handleMessage(ctx context.Context, selfwxid, serverPort string, msgItem gjson.Result) error {
// 	handler, exists := s.getMessageHandler(msgItem.EventType)
// 	if !exists {
// 		glog.Printf(ctx, "未知消息类型 MsgType:%s\n%s", msgItem.EventType, msgItem.Msg)
// 		return nil
// 	}

// 	return handler(ctx, pcuid, port, msgItem)
// }

// getMessageHandler 获取消息处理器
// func (s *Server) getMessageHandler(msgType message.MsgType) (func(string, string, gjson.Result) error, bool) {

// 	handlers := map[message.MsgType]func(*message.WcMsgItem, gjson.Result) error{
// 		message.MsgTypeText:                    s.handleMsgTypeText,
// 		message.MsgTypeImage:                   s.handleMsgTypeImage,
// 		message.MsgTypeFileOrAppShareLinkFile:  s.handleMsgTypeFileOrAppShareLinkFile,
// 		message.MsgTypeGif:                     s.handleMsgTypeGif,
// 		message.MsgTypeFriendConfirmation:      s.handleMsgTypeFriendConfirmation,
// 		message.MsgTypeVideo:                   s.handleMsgTypeVideo,
// 		message.MsgTypeVoice:                   s.handleMsgTypeVoice,
// 		message.MsgTypeShareCard:               s.handleMsgTypeShareCard,
// 		message.MsgTypeShareLocation:           s.handleMsgTypeShareLocation,
// 		message.MsgTypeSystem0:                 s.handleMsgTypeSystem0,
// 		message.MsgTypeSystem2:                 s.handleMsgTypeSystem2,
// 		message.MsgTypeLoginQRCodeRefreshEvent: s.handleMsgTypeLoginQRCodeRefreshEvent,
// 		message.MsgTypeLoginWeChatEvent:        s.handleMsgTypeLoginWeChatEvent,
// 		message.MsgTypeLogoutWeChatEvent:       s.handleMsgTypeLogoutWeChatEvent,
// 		message.MsgTypeSwitchChatObject:        s.handleMsgTypeSwitchChatObject,
// 		message.MsgTypeSwitchContact:           s.handleMsgTypeSwitchContact,
// 	}

// 	handler, exists := handlers[msgType]
// 	return handler, exists
// }

// HandleRequest 处理微信的请求
// func (srv *Server) handleRequest() error {

// 	jsonText := string(srv.RequestRawMsg)

// 	var msgItemList []message.WcMsgItem

// 	// sendorrecv: 1=收到的消息,2=发送的消息
// 	// sendorrecv := gjson.Get(jsonText, "sendorrecv").String()
// 	// if sendorrecv != "2" {
// 	// 	// 不处理发送的消息
// 	// 	return srv.messageHandler(msgItemList)
// 	// }

// 	selfwxid := gjson.Get(jsonText, "selfwxid").String()
// 	serverPort := gjson.Get(jsonText, "ServerPort").String()

// 	msgList := gjson.Get(jsonText, "msglist")

// 	for _, msgInfo := range msgList.Array() {

// 		msgContent := msgInfo.Get("msg").String()
// 		msgsvrid := msgInfo.Get("msgsvrid").String()
// 		time := msgInfo.Get("time").String()
// 		fromType := msgInfo.Get("fromtype").String()
// 		msgType := msgInfo.Get("msgtype").String()
// 		fromid := msgInfo.Get("fromid").String()
// 		fromname := msgInfo.Get("fromname").String()
// 		toid := msgInfo.Get("toid").String()
// 		toname := msgInfo.Get("toname").String()

// 		wcMsgItem := message.WcMsgItem{
// 			ServerPort: serverPort,
// 			SelfWxid:   selfwxid,
// 			CommonMsg: message.CommonMsg{
// 				MsgSvrid: msgsvrid,
// 				Time:     time,
// 				Msg:      msgContent,
// 				MsgType:  msgType,  // 消息类型代码
// 				FromType: fromType, // 个人消息=1 群消息=2
// 				FromId:   fromid,   // 发送方微信ID
// 				FromName: fromname, // 发送方昵称
// 				Index:    msgInfo.Get("index").String(),
// 			},
// 			ToCommonMsg: message.ToCommonMsg{
// 				ToId:   toid,
// 				ToName: toname,
// 			},
// 			EventType: message.UnknownEvent,
// 		}

// 		switch message.MsgType(msgType) {
// 		case message.MsgTypeText:

// 		case message.MsgTypeImage:

// 		case message.MsgTypeFileOrAppShareLinkFile:
// 			// "msgtype":"49"

// 			// if appMsgXml.AppMsg.Type == "57" {
// 			// } else if appMsgXml.AppMsg.Type == "51" {
// 			// }
// 		case message.MsgTypeGif:

// 		case message.MsgTypeFriendConfirmation:

// 		case message.MsgTypeVideo:

// 		case message.MsgTypeVoice:

// 		case message.MsgTypeShareCard: // "msgtype":"42",

// 		case message.MsgTypeLocation: // "msgtype":"48",

// 		case message.MsgTypeShareLocation: // "msgtype":"58",

// 		case message.MsgTypeSystem0: // "msgtype":"1000", 系统消息
// 			// if msgItem.Msg == "位置共享结束" {
// 			// PC收到共享位置结束消息
// 			// eventType = message.PCRecvEndShareLocationMsgEvent
// 			// }
// 		case message.MsgTypeSystem2:

// 			// "msgtype":"10002", 撤回消息
// 			// msgContent := msgInfo.Get("msg").String()

// 		case message.MsgTypeLoginQRCodeRefreshEvent:

// 		case message.MsgTypeLoginWeChatEvent:

// 		case message.MsgTypeLogoutWeChatEvent:

// 		case message.MsgTypeSwitchChatObject:

// 		case message.MsgTypeSwitchContact:

// 		}

// 		msgItemList = append(msgItemList, wcMsgItem)

// 	}
// 	return srv.messageHandler(msgItemList)
// }
