package server

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"

	"github.com/azure1489/wechat/callbackevent/message"
)

type Server struct {

	// Request *http.Request

	RequestRawMsg []byte
	RequestMsg    *message.MsgBody
	// WcMsgBody      *message.WcMsgBody
	// CommonMsg 子类为入参
	messageHandler func(msgList []message.WcMsgItem) error
}

func NewServer(body []byte) *Server {
	srv := new(Server)
	srv.RequestRawMsg = body
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

// 判断是否是群消息
func (srv *Server) isGroupMsg(username string) bool {
	return strings.HasSuffix(username, "@chatroom")
}

// interface{} 转 sring
func interfaceToString(i interface{}) string {
	if i == nil {
		return ""
	}
	return i.(string)
}

// interface{} 转 sring
func interfaceToFloat64(i interface{}) float64 {
	if i == nil {
		return 0
	}
	return i.(float64)
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

	// wcMsgBody := message.WcMsgBody{

	// 	Sendorrecv: msgBody.Sendorrecv,
	// 	MsgNumber:  msgBody.MsgNumber,
	// }

	var msgItemList []message.WcMsgItem

	// sendorrecv: 1=收到的消息,2=发送的消息
	sendorrecv := msgBody.Sendorrecv

	if sendorrecv != "2" {
		return nil
	}

	msgList := msgBody.Msglist

	if len(msgList) == 0 {
		return errors.New("消息列表为空")
	}

	// if sendorrecv == "1" {
	// 收到的消息
	// for _, msgItem := range msgList {

	// 	msgType := message.MsgType(msgItem.MsgType)

	// 	switch msgType {
	// 	case message.MsgTypeText:
	// 		// PC发送文本消息
	// 		eventType = message.PCSendTextMsgEvent
	// 	case message.MsgTypeImage:
	// 		// PC发送图片消息
	// 		eventType = message.PCSendImgMsgEvent
	// 	case message.MsgTypeFileOrAppShareLinkFile:
	// 		// PC发送文件/app消息
	// 		eventType = message.PCSendFileOrAppMsgEvent
	// 	case message.MsgTypeGif:
	// 		// PC发动态图片消息成功
	// 		eventType = message.PCSendGifImgMsgEvent
	// 	}

	// }
	// } else
	// if sendorrecv == "2" {
	// 定义一个可以存放 CommonMsg的子类的数组

	for _, msgItem := range msgList {

		fromid := interfaceToString(msgItem["fromid"])

		// log.Println(" ---------------- msgBody.SelfWxid:", msgBody.SelfWxid, " ---------------- ")
		// log.Println(" ---------------- fromid:", fromid, " ---------------- ")

		msgType := message.MsgType(interfaceToString(msgItem["msgtype"]))

		fromtype := interfaceToString(msgItem["fromtype"])

		// if msgBody.SelfWxid == fromid {
		// 	log.Println(" ---------------- 自己发送的消息 ---------------- ")

		// }

		toid := interfaceToString(msgItem["toid"])

		if fromtype == "1" && srv.isGroupMsg(toid) {
			fromtype = "2"
		}

		msgContent := interfaceToString(msgItem["msg"])

		wcMsgItem := message.WcMsgItem{
			ServerPort: msgBody.ServerPort,
			SelfWxid:   msgBody.SelfWxid,
			CommonMsg: message.CommonMsg{
				MsgSvrid: interfaceToString(msgItem["msgsvrid"]),
				Time:     interfaceToString(msgItem["time"]),
				Msg:      msgContent,
				MsgType:  interfaceToString(msgItem["msgtype"]),  // 消息类型代码
				FromType: fromtype,                               // 个人消息=1 群消息=2
				FromId:   fromid,                                 // 发送方微信ID
				FromName: interfaceToString(msgItem["fromname"]), // 发送方昵称
				Index:    interfaceToString(msgItem["index"]),
			},
			ToCommonMsg: message.ToCommonMsg{
				ToId:   interfaceToString(msgItem["toid"]),
				ToName: interfaceToString(msgItem["toname"]),
			},
		}

		switch msgType {
		case message.MsgTypeText:
			if msgContent == "PC发文本消息成功" {
				continue
			}

			text := message.Text{
				MsgSource: interfaceToString(msgItem["msgsource"]), // 消息源内容
			}
			if fromtype == "1" {
				wcMsgItem.MsgItem = text
				wcMsgItem.EventType = message.PCRecvTextMsgEvent
			} else if fromtype == "2" {
				groupText := message.GroupText{
					Text: text, // 消息源内容
					CommonGroupMsg: message.CommonGroupMsg{
						FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
						FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
					},
				}
				wcMsgItem.MsgItem = groupText
				wcMsgItem.EventType = message.PCRecvGroupTextMsgEvent
			}
		case message.MsgTypeImage: // PC收到图片消息
			if msgContent == "PC发图片消息成功" {
				continue
			}

			image := message.Image{
				Info:      interfaceToString(msgItem["info"]),       // 消息源内容
				ImgLen:    interfaceToFloat64(msgItem["img_len"]),   // 消息源内容
				ImgPath:   interfaceToString(msgItem["img_path"]),   // 消息源内容
				ImgBase64: interfaceToString(msgItem["img_base64"]), // 消息源内容
			}

			if fromtype == "1" {
				wcMsgItem.MsgItem = image
				wcMsgItem.EventType = message.PCRecvImgMsgEvent
			} else if fromtype == "2" {
				wcMsg := message.GroupImage{
					Image: image,
					CommonGroupMsg: message.CommonGroupMsg{
						FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
						FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
					},
				}
				wcMsgItem.MsgItem = wcMsg
				wcMsgItem.EventType = message.PCRecvGroupImgMsgEvent
			}
			// log.Println(" ---------------- PC收到图片消息处理结束 ---------------- ")
		case message.MsgTypeFileOrAppShareLinkFile: // "msgtype":"49",
			if msgContent == "PC发app/文件消息成功" {
				continue
			}
			var appMsgXml message.AppMsgXml
			// 字符串转换为xml
			err = xml.Unmarshal([]byte(msgContent), &appMsgXml)
			if err != nil {
				return err
			}
			if appMsgXml.AppMsg.Type == "57" { // 引用消息
				quote := message.Quote{
					MsgSource:    appMsgXml.AppMsg.ReferMsg.MsgSource, // 消息源内容
					QuoteMsg:     appMsgXml.AppMsg.ReferMsg.Content,   // 引用的消息内容
					QuoteMsgType: appMsgXml.AppMsg.ReferMsg.Type,      // 引用的消息类型
					QuoteMsgId:   appMsgXml.AppMsg.ReferMsg.Svrid,     // 引用的消息id
					ReplyMsg:     appMsgXml.AppMsg.Title,              // 回复的消息内容
				}
				if fromtype == "1" {
					wcMsgItem.MsgItem = quote
					wcMsgItem.EventType = message.PCRecvQuoteMsgEvent
				} else if fromtype == "2" {
					wcMsg := message.GroupQuote{ // 回复的消息内容
						Quote: quote, // 引用的消息内容
						CommonGroupMsg: message.CommonGroupMsg{
							FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
							FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
						},
					}
					wcMsgItem.MsgItem = wcMsg
					wcMsgItem.EventType = message.PCRecvGroupQuoteMsgEvent
				}
			}
		case message.MsgTypeGif: // "msgtype":"47",
			if msgContent == "PC发动态图片消息成功" {
				continue
			}
			gif := message.Gif{
				GifPath: interfaceToString(msgItem["gif_path"]),
			}
			if fromtype == "1" {
				wcMsgItem.MsgItem = gif
				wcMsgItem.EventType = message.PCRecvGifImgMsgEvent
			} else if fromtype == "2" {
				wcMsg := message.GroupGif{
					Gif: gif,
					CommonGroupMsg: message.CommonGroupMsg{
						FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
						FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
					},
				}
				wcMsgItem.MsgItem = wcMsg
				wcMsgItem.EventType = message.PCRecvGroupGifImgMsgEvent
			}
		case message.MsgTypeVideo: // "msgtype":"43", PC收到视频消息

			video := message.Video{
				Info:      interfaceToString(msgItem["info"]), // 消息源内容
				VideoPath: interfaceToString(msgItem["video_path"]),
			}
			if fromtype == "1" {
				wcMsgItem.MsgItem = video
				wcMsgItem.EventType = message.PCRecvVideoMsgEvent
			} else if fromtype == "2" {
				wcMsg := message.GroupVideo{
					Video: video,
					CommonGroupMsg: message.CommonGroupMsg{
						FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
						FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
					},
				}
				wcMsgItem.MsgItem = wcMsg
				wcMsgItem.EventType = message.PCRecvGroupVideoMsgEvent
			}
		case message.MsgTypeVoice: // "msgtype":"34", PC收到语音消息
			voice := message.Voice{
				VoiceLen: interfaceToString(msgItem["voice_len"]),
				VoiceHex: interfaceToString(msgItem["voice_hex"]),
			}
			if fromtype == "1" {
				wcMsgItem.MsgItem = voice
				wcMsgItem.EventType = message.PCRecvVoiceMsgEvent
			} else if fromtype == "2" {
				wcMsg := message.GroupVoice{
					Voice: voice,
					CommonGroupMsg: message.CommonGroupMsg{
						FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
						FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
					},
				}
				wcMsgItem.MsgItem = wcMsg
				wcMsgItem.EventType = message.PCRecvGroupVoiceMsgEvent
			}
		case message.MsgTypeShareCard: // "msgtype":"42",
			// PC收到名片消息
			// eventType = message.PCRecvShareCardMsgEvent
		case message.MsgTypeLocation: // "msgtype":"48",
			// PC收到位置消息
			// eventType = message.PCRecvLocationMsgEvent
		case message.MsgTypeShareLocation: // "msgtype":"58",
			// PC收到共享位置消息
			// eventType = message.PCRecvShareLocationMsgEvent
		case message.MsgTypeSystem: // "msgtype":"1000", 系统消息
			// if msgItem.Msg == "位置共享结束" {
			// PC收到共享位置结束消息
			// eventType = message.PCRecvEndShareLocationMsgEvent
			// }
		case message.MsgTypeRecall: // "msgtype":"10002", 撤回消息
			sysMsgXml := message.SysMsgXml{}
			err = xml.Unmarshal([]byte(msgContent), &sysMsgXml)
			if err != nil {
				return err
			}
			if sysMsgXml.Type == "revokemsg" { // 撤回消息
				revoke := message.Revoke{
					RevokeMsg:  interfaceToString(msgItem["revoke_msg"]),
					Session:    sysMsgXml.RevokeMsg.Session,
					MsgId:      sysMsgXml.RevokeMsg.MsgId,
					NewMsgId:   sysMsgXml.RevokeMsg.NewMsgId,
					ReplaceMsg: sysMsgXml.RevokeMsg.ReplaceMsg,
				}
				if fromtype == "1" {
					wcMsgItem.MsgItem = revoke
					wcMsgItem.EventType = message.PCRecvRevokeMsgEvent
				} else if fromtype == "2" {
					wcMsg := message.GroupRevoke{
						Revoke: revoke,
						CommonGroupMsg: message.CommonGroupMsg{
							FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
							FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
						},
					}
					wcMsgItem.MsgItem = wcMsg
					wcMsgItem.EventType = message.PCRecvGroupRevokeMsgEvent
				}
			}
		}
		msgItemList = append(msgItemList, wcMsgItem)

	}
	// }

	return srv.messageHandler(msgItemList)
}
