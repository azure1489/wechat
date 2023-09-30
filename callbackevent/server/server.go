package server

import (
	"encoding/json"
	"encoding/xml"
	"errors"

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

	if msgBody.Msglist == nil || len(msgBody.Msglist) == 0 {
		return errors.New("消息列表为空")
	}

	msgList := msgBody.Msglist

	// sendorrecv: 1=收到的消息,2=发送的消息
	sendorrecv := msgBody.Sendorrecv
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
	var msgItemList []message.WcMsgItem

	// }
	// } else
	if sendorrecv == "2" {
		// 定义一个可以存放 CommonMsg的子类的数组

		for _, msgItem := range msgList {

			msgtype := msgItem["msgtype"]

			msgType := message.MsgType(msgtype)

			switch msgType {
			case message.MsgTypeText:
				if msgItem["msg"] != "PC发文本消息成功" {
					if msgItem["fromtype"] == "1" {
						wcMsg := message.Text{
							CommonMsg: message.CommonMsg{
								MsgSvrid:   msgItem["msgsvrid"],
								Time:       msgItem["time"],
								Msg:        msgItem["msg"],
								MsgType:    msgItem["msgtype"],  // 消息类型代码
								FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
								FromId:     msgItem["fromid"],   // 发送方微信ID
								FromName:   msgItem["fromname"], // 发送方昵称
								Index:      msgItem["index"],
								ServerPort: msgBody.ServerPort,
								SelfWxid:   msgBody.SelfWxid,
							},
							ToCommonMsg: message.ToCommonMsg{
								ToId:   msgItem["toid"],
								ToName: msgItem["toname"],
							},
							MsgSource: msgItem["msgsource"], // 消息源内容
						}
						wcMsgItem := message.WcMsgItem{
							EventType: message.PCRecvTextMsgEvent,
							MsgItem:   wcMsg,
						}
						msgItemList = append(msgItemList, wcMsgItem)
					} else if msgItem["fromtype"] == "2" {
						wcMsg := message.GroupText{
							CommonMsg: message.CommonMsg{
								MsgSvrid:   msgItem["msgsvrid"],
								Time:       msgItem["time"],
								Msg:        msgItem["msg"],
								MsgType:    msgItem["msgtype"],  // 消息类型代码
								FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
								FromId:     msgItem["fromid"],   // 发送方微信ID
								FromName:   msgItem["fromname"], // 发送方昵称
								Index:      msgItem["index"],
								ServerPort: msgBody.ServerPort,
								SelfWxid:   msgBody.SelfWxid,
							},
							ToCommonMsg: message.ToCommonMsg{
								ToId:   msgItem["toid"],
								ToName: msgItem["toname"],
							},
							MsgSource: msgItem["msgsource"], // 消息源内容
							CommonGroupMsg: message.CommonGroupMsg{
								FromGname: msgItem["fromgname"], // 群名称
								FromGid:   msgItem["fromgid"],   // 群ID
							},
						}
						wcMsgItem := message.WcMsgItem{
							EventType: message.PCRecvTextMsgEvent,
							MsgItem:   wcMsg,
						}
						msgItemList = append(msgItemList, wcMsgItem)
					}
				}
			case message.MsgTypeImage:
				if msgItem["msg"] != "PC发图片消息成功" {
					// PC收到图片消息
					// eventType = message.PCRecvImgMsgEvent
					if msgItem["fromtype"] == "1" {
						wcMsg := message.Image{
							CommonMsg: message.CommonMsg{
								MsgSvrid:   msgItem["msgsvrid"],
								Time:       msgItem["time"],
								Msg:        msgItem["msg"],
								MsgType:    msgItem["msgtype"],  // 消息类型代码
								FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
								FromId:     msgItem["fromid"],   // 发送方微信ID
								FromName:   msgItem["fromname"], // 发送方昵称
								Index:      msgItem["index"],
								ServerPort: msgBody.ServerPort,
								SelfWxid:   msgBody.SelfWxid,
							},
							ToCommonMsg: message.ToCommonMsg{
								ToId:   msgItem["toid"],
								ToName: msgItem["toname"],
							},
							Info:      msgItem["info"],       // 消息源内容
							ImgLen:    msgItem["img_len"],    // 消息源内容
							ImgPath:   msgItem["img_path"],   // 消息源内容
							ImgBase64: msgItem["img_base64"], // 消息源内容
						}
						wcMsgItem := message.WcMsgItem{
							EventType: message.PCRecvImgMsgEvent,
							MsgItem:   wcMsg,
						}
						msgItemList = append(msgItemList, wcMsgItem)
					} else if msgItem["fromtype"] == "2" {
						wcMsg := message.GroupImage{
							CommonMsg: message.CommonMsg{
								MsgSvrid:   msgItem["msgsvrid"],
								Time:       msgItem["time"],
								Msg:        msgItem["msg"],
								MsgType:    msgItem["msgtype"],  // 消息类型代码
								FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
								FromId:     msgItem["fromid"],   // 发送方微信ID
								FromName:   msgItem["fromname"], // 发送方昵称
								Index:      msgItem["index"],
								ServerPort: msgBody.ServerPort,
								SelfWxid:   msgBody.SelfWxid,
							},
							ToCommonMsg: message.ToCommonMsg{
								ToId:   msgItem["toid"],
								ToName: msgItem["toname"],
							},
							Info:      msgItem["info"],       // 消息源内容
							ImgLen:    msgItem["img_len"],    // 消息源内容
							ImgPath:   msgItem["img_path"],   // 消息源内容
							ImgBase64: msgItem["img_base64"], // 消息源内容
							CommonGroupMsg: message.CommonGroupMsg{
								FromGname: msgItem["fromgname"], // 群名称
								FromGid:   msgItem["fromgid"],   // 群ID
							},
						}
						wcMsgItem := message.WcMsgItem{
							EventType: message.PCRecvImgMsgEvent,
							MsgItem:   wcMsg,
						}
						msgItemList = append(msgItemList, wcMsgItem)
					}
				}
			case message.MsgTypeFileOrAppShareLinkFile: // "msgtype":"49",
				msg := msgItem["msg"]
				if msg != "PC发app/文件消息成功" {
					var appMsgXml message.AppMsgXml
					// 字符串转换为xml
					err = xml.Unmarshal([]byte(msg), &appMsgXml)
					if err != nil {
						return err
					}

					if appMsgXml.Msg.AppMsg.Type == "57" && appMsgXml.Msg.AppMsg.ReferMsg.Type == "47" { // 引用消息
						if msgItem["fromtype"] == "1" {
							wcMsg := message.Quote{
								CommonMsg: message.CommonMsg{
									MsgSvrid:   msgItem["msgsvrid"],
									Time:       msgItem["time"],
									Msg:        msgItem["msg"],
									MsgType:    msgItem["msgtype"],  // 消息类型代码
									FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
									FromId:     msgItem["fromid"],   // 发送方微信ID
									FromName:   msgItem["fromname"], // 发送方昵称
									Index:      msgItem["index"],
									ServerPort: msgBody.ServerPort,
									SelfWxid:   msgBody.SelfWxid,
								},
								ToCommonMsg: message.ToCommonMsg{
									ToId:   msgItem["toid"],
									ToName: msgItem["toname"],
								},
								MsgSource:  appMsgXml.Msg.AppMsg.ReferMsg.MsgSource, // 消息源内容
								QuoteMsg:   appMsgXml.Msg.AppMsg.ReferMsg.Content,   // 引用的消息内容
								QuoteMsgId: appMsgXml.Msg.AppMsg.ReferMsg.Svrid,     // 引用的消息id
								ReplyMsg:   appMsgXml.Msg.AppMsg.Title,              // 回复的消息内容
							}
							wcMsgItem := message.WcMsgItem{
								EventType: message.PCRecvQuoteMsgEvent,
								MsgItem:   wcMsg,
							}
							msgItemList = append(msgItemList, wcMsgItem)
						} else if msgItem["fromtype"] == "2" {
							wcMsg := message.GroupQuote{
								CommonMsg: message.CommonMsg{
									MsgSvrid:   msgItem["msgsvrid"],
									Time:       msgItem["time"],
									Msg:        msgItem["msg"],
									MsgType:    msgItem["msgtype"],  // 消息类型代码
									FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
									FromId:     msgItem["fromid"],   // 发送方微信ID
									FromName:   msgItem["fromname"], // 发送方昵称
									Index:      msgItem["index"],
									ServerPort: msgBody.ServerPort,
									SelfWxid:   msgBody.SelfWxid,
								},
								ToCommonMsg: message.ToCommonMsg{
									ToId:   msgItem["toid"],
									ToName: msgItem["toname"],
								},
								MsgSource:  appMsgXml.Msg.AppMsg.ReferMsg.MsgSource, // 消息源内容
								QuoteMsg:   appMsgXml.Msg.AppMsg.ReferMsg.Content,   // 引用的消息内容
								QuoteMsgId: appMsgXml.Msg.AppMsg.ReferMsg.Svrid,     // 引用的消息id
								ReplyMsg:   appMsgXml.Msg.AppMsg.Title,              // 回复的消息内容
								CommonGroupMsg: message.CommonGroupMsg{
									FromGname: msgItem["fromgname"], // 群名称
									FromGid:   msgItem["fromgid"],   // 群ID
								},
							}
							wcMsgItem := message.WcMsgItem{
								EventType: message.PCRecvQuoteMsgEvent,
								MsgItem:   wcMsg,
							}
							msgItemList = append(msgItemList, wcMsgItem)
						}
					}
				}
			case message.MsgTypeGif: // "msgtype":"47",
				if msgItem["msg"] != "PC发动态图片消息成功" {
					if msgItem["fromtype"] == "1" {
						wcMsg := message.Gif{
							CommonMsg: message.CommonMsg{
								MsgSvrid:   msgItem["msgsvrid"],
								Time:       msgItem["time"],
								Msg:        msgItem["msg"],
								MsgType:    msgItem["msgtype"],  // 消息类型代码
								FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
								FromId:     msgItem["fromid"],   // 发送方微信ID
								FromName:   msgItem["fromname"], // 发送方昵称
								Index:      msgItem["index"],
								ServerPort: msgBody.ServerPort,
								SelfWxid:   msgBody.SelfWxid,
							},
							ToCommonMsg: message.ToCommonMsg{
								ToId:   msgItem["toid"],
								ToName: msgItem["toname"],
							},
							GifPath: msgItem["gif_path"],
						}
						wcMsgItem := message.WcMsgItem{
							EventType: message.PCRecvGifImgMsgEvent,
							MsgItem:   wcMsg,
						}
						msgItemList = append(msgItemList, wcMsgItem)
					} else if msgItem["fromtype"] == "2" {
						wcMsg := message.GroupGif{
							CommonMsg: message.CommonMsg{
								MsgSvrid:   msgItem["msgsvrid"],
								Time:       msgItem["time"],
								Msg:        msgItem["msg"],
								MsgType:    msgItem["msgtype"],  // 消息类型代码
								FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
								FromId:     msgItem["fromid"],   // 发送方微信ID
								FromName:   msgItem["fromname"], // 发送方昵称
								Index:      msgItem["index"],
								ServerPort: msgBody.ServerPort,
								SelfWxid:   msgBody.SelfWxid,
							},
							ToCommonMsg: message.ToCommonMsg{
								ToId:   msgItem["toid"],
								ToName: msgItem["toname"],
							},
							GifPath: msgItem["gif_path"],
							CommonGroupMsg: message.CommonGroupMsg{
								FromGname: msgItem["fromgname"], // 群名称
								FromGid:   msgItem["fromgid"],   // 群ID
							},
						}
						wcMsgItem := message.WcMsgItem{
							EventType: message.PCRecvGifImgMsgEvent,
							MsgItem:   wcMsg,
						}
						msgItemList = append(msgItemList, wcMsgItem)
					}
				}
			case message.MsgTypeVideo: // "msgtype":"43",
				if msgItem["fromtype"] == "1" {
					wcMsg := message.Video{
						CommonMsg: message.CommonMsg{
							MsgSvrid:   msgItem["msgsvrid"],
							Time:       msgItem["time"],
							Msg:        msgItem["msg"],
							MsgType:    msgItem["msgtype"],  // 消息类型代码
							FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
							FromId:     msgItem["fromid"],   // 发送方微信ID
							FromName:   msgItem["fromname"], // 发送方昵称
							Index:      msgItem["index"],
							ServerPort: msgBody.ServerPort,
							SelfWxid:   msgBody.SelfWxid,
						},
						ToCommonMsg: message.ToCommonMsg{
							ToId:   msgItem["toid"],
							ToName: msgItem["toname"],
						},
						Info:      msgItem["info"], // 消息源内容
						VideoPath: msgItem["video_path"],
					}
					wcMsgItem := message.WcMsgItem{
						EventType: message.PCRecvVideoMsgEvent,
						MsgItem:   wcMsg,
					}
					msgItemList = append(msgItemList, wcMsgItem)
				} else if msgItem["fromtype"] == "2" {
					wcMsg := message.GroupVideo{
						CommonMsg: message.CommonMsg{
							MsgSvrid:   msgItem["msgsvrid"],
							Time:       msgItem["time"],
							Msg:        msgItem["msg"],
							MsgType:    msgItem["msgtype"],  // 消息类型代码
							FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
							FromId:     msgItem["fromid"],   // 发送方微信ID
							FromName:   msgItem["fromname"], // 发送方昵称
							Index:      msgItem["index"],
							ServerPort: msgBody.ServerPort,
							SelfWxid:   msgBody.SelfWxid,
						},
						ToCommonMsg: message.ToCommonMsg{
							ToId:   msgItem["toid"],
							ToName: msgItem["toname"],
						},
						Info:      msgItem["info"], // 消息源内容
						VideoPath: msgItem["video_path"],
						CommonGroupMsg: message.CommonGroupMsg{
							FromGname: msgItem["fromgname"], // 群名称
							FromGid:   msgItem["fromgid"],   // 群ID
						},
					}
					wcMsgItem := message.WcMsgItem{
						EventType: message.PCRecvVideoMsgEvent,
						MsgItem:   wcMsg,
					}
					msgItemList = append(msgItemList, wcMsgItem)
				}
			case message.MsgTypeVoice: // "msgtype":"34",
				// PC收到语音消息
				if msgItem["fromtype"] == "1" {
					wcMsg := message.Voice{
						CommonMsg: message.CommonMsg{
							MsgSvrid:   msgItem["msgsvrid"],
							Time:       msgItem["time"],
							Msg:        msgItem["msg"],
							MsgType:    msgItem["msgtype"],  // 消息类型代码
							FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
							FromId:     msgItem["fromid"],   // 发送方微信ID
							FromName:   msgItem["fromname"], // 发送方昵称
							Index:      msgItem["index"],
							ServerPort: msgBody.ServerPort,
							SelfWxid:   msgBody.SelfWxid,
						},
						ToCommonMsg: message.ToCommonMsg{
							ToId:   msgItem["toid"],
							ToName: msgItem["toname"],
						},
						VoiceLen: msgItem["voice_len"],
						VoiceHex: msgItem["voice_hex"],
					}
					wcMsgItem := message.WcMsgItem{
						EventType: message.PCRecvVoiceMsgEvent,
						MsgItem:   wcMsg,
					}
					msgItemList = append(msgItemList, wcMsgItem)
				} else if msgItem["fromtype"] == "2" {
					wcMsg := message.GroupVoice{
						CommonMsg: message.CommonMsg{
							MsgSvrid:   msgItem["msgsvrid"],
							Time:       msgItem["time"],
							Msg:        msgItem["msg"],
							MsgType:    msgItem["msgtype"],  // 消息类型代码
							FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
							FromId:     msgItem["fromid"],   // 发送方微信ID
							FromName:   msgItem["fromname"], // 发送方昵称
							Index:      msgItem["index"],
							ServerPort: msgBody.ServerPort,
							SelfWxid:   msgBody.SelfWxid,
						},
						ToCommonMsg: message.ToCommonMsg{
							ToId:   msgItem["toid"],
							ToName: msgItem["toname"],
						},
						VoiceLen: msgItem["voice_len"],
						VoiceHex: msgItem["voice_hex"],
						CommonGroupMsg: message.CommonGroupMsg{
							FromGname: msgItem["fromgname"], // 群名称
							FromGid:   msgItem["fromgid"],   // 群ID
						},
					}
					wcMsgItem := message.WcMsgItem{
						EventType: message.PCRecvVoiceMsgEvent,
						MsgItem:   wcMsg,
					}
					msgItemList = append(msgItemList, wcMsgItem)
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
				msg := msgItem["msg"]

				sysMsgXml := message.SysMsgXml{}
				err = xml.Unmarshal([]byte(msg), &sysMsgXml)
				if err != nil {
					return err
				}

				if sysMsgXml.SysMsg.Type == "revokemsg" {
					if msgItem["fromtype"] == "1" {
						wcMsg := message.Revoke{
							CommonMsg: message.CommonMsg{
								MsgSvrid:   msgItem["msgsvrid"],
								Time:       msgItem["time"],
								Msg:        msgItem["msg"],
								MsgType:    msgItem["msgtype"],  // 消息类型代码
								FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
								FromId:     msgItem["fromid"],   // 发送方微信ID
								FromName:   msgItem["fromname"], // 发送方昵称
								Index:      msgItem["index"],
								ServerPort: msgBody.ServerPort,
								SelfWxid:   msgBody.SelfWxid,
							},
							ToCommonMsg: message.ToCommonMsg{
								ToId:   msgItem["toid"],
								ToName: msgItem["toname"],
							},
							RevokeMsg:  msgItem["revoke_msg"],
							Session:    sysMsgXml.SysMsg.RevokeMsg.Session,
							MsgId:      sysMsgXml.SysMsg.RevokeMsg.MsgId,
							NewMsgId:   sysMsgXml.SysMsg.RevokeMsg.NewMsgId,
							ReplaceMsg: sysMsgXml.SysMsg.RevokeMsg.ReplaceMsg,
						}
						wcMsgItem := message.WcMsgItem{
							EventType: message.PCRecvRevokeMsgEvent,
							MsgItem:   wcMsg,
						}
						msgItemList = append(msgItemList, wcMsgItem)
					} else if msgItem["fromtype"] == "2" {
						wcMsg := message.GroupRevoke{
							CommonMsg: message.CommonMsg{
								MsgSvrid:   msgItem["msgsvrid"],
								Time:       msgItem["time"],
								Msg:        msgItem["msg"],
								MsgType:    msgItem["msgtype"],  // 消息类型代码
								FromType:   msgItem["fromtype"], // 个人消息=1 群消息=2
								FromId:     msgItem["fromid"],   // 发送方微信ID
								FromName:   msgItem["fromname"], // 发送方昵称
								Index:      msgItem["index"],
								ServerPort: msgBody.ServerPort,
								SelfWxid:   msgBody.SelfWxid,
							},
							ToCommonMsg: message.ToCommonMsg{
								ToId:   msgItem["toid"],
								ToName: msgItem["toname"],
							},
							RevokeMsg:  msgItem["revoke_msg"],
							Session:    sysMsgXml.SysMsg.RevokeMsg.Session,
							MsgId:      sysMsgXml.SysMsg.RevokeMsg.MsgId,
							NewMsgId:   sysMsgXml.SysMsg.RevokeMsg.NewMsgId,
							ReplaceMsg: sysMsgXml.SysMsg.RevokeMsg.ReplaceMsg,
							CommonGroupMsg: message.CommonGroupMsg{
								FromGname: msgItem["fromgname"], // 群名称
								FromGid:   msgItem["fromgid"],   // 群ID
							},
						}
						wcMsgItem := message.WcMsgItem{
							EventType: message.PCRecvRevokeMsgEvent,
							MsgItem:   wcMsg,
						}
						msgItemList = append(msgItemList, wcMsgItem)
					}
				}
			}

			// msgItem.EventType = eventType
			// msgItem.MsgType = msgType

		}
	}

	return srv.messageHandler(msgItemList)
}
