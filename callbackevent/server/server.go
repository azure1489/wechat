package server

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

// interface{} 转 sring
func stringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return f
}

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

// HandleRequest 处理微信的请求
func (srv *Server) handleRequest() error {

	jsonText := string(srv.RequestRawMsg)

	var msgItemList []message.WcMsgItem

	// sendorrecv: 1=收到的消息,2=发送的消息
	sendorrecv := gjson.Get(jsonText, "sendorrecv").String()
	if sendorrecv != "2" {
		// 不处理发送的消息
		return srv.messageHandler(msgItemList)
	}

	selfwxid := gjson.Get(jsonText, "selfwxid").String()
	serverPort := gjson.Get(jsonText, "ServerPort").String()

	msgList := gjson.Get(jsonText, "msglist")

	for _, msgInfo := range msgList.Array() {

		msgContent := msgInfo.Get("msg").String()

		msgsvrid := msgInfo.Get("msgsvrid").String()
		time := msgInfo.Get("time").String()
		fromType := msgInfo.Get("fromtype").String()
		msgType := msgInfo.Get("msgtype").String()

		fromid := msgInfo.Get("fromid").String()
		fromname := msgInfo.Get("fromname").String()

		toid := msgInfo.Get("toid").String()
		toname := msgInfo.Get("toname").String()

		wcMsgItem := message.WcMsgItem{
			ServerPort: serverPort,
			SelfWxid:   selfwxid,
			CommonMsg: message.CommonMsg{
				MsgSvrid: msgsvrid,
				Time:     time,
				Msg:      msgContent,
				MsgType:  msgType,  // 消息类型代码
				FromType: fromType, // 个人消息=1 群消息=2
				FromId:   fromid,   // 发送方微信ID
				FromName: fromname, // 发送方昵称
				Index:    msgInfo.Get("index").String(),
			},
			ToCommonMsg: message.ToCommonMsg{
				ToId:   toid,
				ToName: toname,
			},
			EventType: message.UnknownEvent,
		}

		switch message.MsgType(msgType) {
		case message.MsgTypeText:

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

		case message.MsgTypeImage:
			// PC收到图片消息
			image := message.Image{
				Info:      msgInfo.Get("info").String(),                     // 消息源内容
				ImgLen:    stringToFloat64(msgInfo.Get("img_len").String()), // 消息源内容
				ImgPath:   msgInfo.Get("img_path").String(),                 // 消息源内容
				ImgBase64: msgInfo.Get("img_base64").String(),               // 消息源内容
			}

			wcMsgItem.MsgItem = image

			// fromtype: 1=个人消息,2=群消息
			if fromType == "1" {
				wcMsgItem.EventType = message.PCRecvImgMsgEvent
			} else if fromType == "2" {
				wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
					FromGname: msgInfo.Get("fromgname").String(), // 群名称
					FromGid:   msgInfo.Get("fromgid").String(),   // 群ID
				}
				wcMsgItem.EventType = message.PCRecvGroupImgMsgEvent
			}
		case message.MsgTypeFileOrAppShareLinkFile:
			// "msgtype":"49"

			msgContent := msgInfo.Get("msg").String()

			var appMsgXml message.AppMsgXml
			// 字符串转换为xml
			err := xml.Unmarshal([]byte(msgContent), &appMsgXml)
			if err != nil {
				fmt.Printf("xml.Unmarshal err: %v\n", err)
				continue
			}

			if appMsgXml.AppMsg.Type == "57" {
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

			} else if appMsgXml.AppMsg.Type == "51" {
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
			}
		case message.MsgTypeGif:
			// 自定义表情消息 "msgtype":"47",
			gifMsg := message.Gif{
				GifPath: msgInfo.Get("gif_path").String(),
			}

			if emojiMsg, err := getEmojiMsg(msgContent); err == nil {
				base64String := emojiMsg.Emoji.Desc
				if base64String != "" {
					if desc, err := decodeString(base64String); err == nil {
						desc = strings.TrimSpace(desc)
						desc = strings.ReplaceAll(desc, "\n", "")

						if desc2, err := getDesc(desc); err == nil {
							gifMsg.Desc = desc2
						}

					}
				}
			}

			wcMsgItem.MsgItem = gifMsg

			// fromtype: 1=个人消息,2=群消息
			if fromType == "1" {
				wcMsgItem.EventType = message.PCRecvGifImgMsgEvent
			} else if fromType == "2" {
				wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
					FromGname: msgInfo.Get("fromgname").String(), // 群名称
					FromGid:   msgInfo.Get("fromgid").String(),   // 群ID
				}
				wcMsgItem.EventType = message.PCRecvGroupGifImgMsgEvent
			}
		case message.MsgTypeVideo:
			// "msgtype":"43", PC收到视频消息
			wcMsgItem.MsgItem = message.Video{
				Info:      msgInfo.Get("info").String(), // 消息源内容
				VideoPath: msgInfo.Get("video_path").String(),
			}

			// fromtype: 1=个人消息,2=群消息
			if fromType == "1" {
				wcMsgItem.EventType = message.PCRecvVideoMsgEvent
			} else if fromType == "2" {
				wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
					FromGname: msgInfo.Get("fromgname").String(), // 群名称
					FromGid:   msgInfo.Get("fromgid").String(),   // 群ID
				}
				wcMsgItem.EventType = message.PCRecvGroupVideoMsgEvent
			}
		case message.MsgTypeVoice:
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
		case message.MsgTypeShareCard: // "msgtype":"42",
			// PC收到名片消息
			wcMsgItem.EventType = message.PCRecvShareCardMsgEvent
		case message.MsgTypeLocation: // "msgtype":"48",
			// PC收到位置消息
			wcMsgItem.EventType = message.PCRecvLocationMsgEvent
		case message.MsgTypeShareLocation: // "msgtype":"58",
			// PC收到共享位置消息
			wcMsgItem.EventType = message.PCRecvShareLocationMsgEvent
		case message.MsgTypeSystem0: // "msgtype":"1000", 系统消息
			// if msgItem.Msg == "位置共享结束" {
			// PC收到共享位置结束消息
			// eventType = message.PCRecvEndShareLocationMsgEvent
			// }
		case message.MsgTypeSystem2:

			// "msgtype":"10002", 撤回消息
			// msgContent := msgInfo.Get("msg").String()

			sysMsgXml := message.SysMsgXml{}
			err := xml.Unmarshal([]byte(msgContent), &sysMsgXml)
			if err != nil {
				fmt.Printf("xml.Unmarshal err: %v\n", err)
				continue
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

		case message.MsgTypeLoginQRCodeRefreshEvent:
			//表示登陆二维码刷新事件
			wcMsgItem.EventType = message.PCLoginQrcodeRefreshEvent
		case message.MsgTypeLoginWeChatEvent:
			//表示登陆微信事件
			wcMsgItem.EventType = message.PCLoginWxEvent
		case message.MsgTypeLogoutWeChatEvent:
			//表示退出登陆微信
			wcMsgItem.EventType = message.PCLogoutWxEvent
		case message.MsgTypeSwitchChatObject:
			//表示切换聊天对象
			wcMsgItem.EventType = message.PCSwitchChatObjectEvent
		case message.MsgTypeSwitchContact:
			//表示切换联系人
			wcMsgItem.EventType = message.PCSwitchContactEvent
		}

		msgItemList = append(msgItemList, wcMsgItem)

	}

	// // 设置事件类型
	// srv.SetEventType(jsonText)

	// // fromtype: 1=个人消息,2=群消息
	// fromType := srv.getFromType(jsonText)

	// // fromtype := gjson.Get(jsonText, "msglist.0.fromtype").String()
	// // msgsvrid := gjson.Get(jsonText, "msglist.0.msgsvrid").String()
	// // msgsvrid := gjson.Get(jsonText, "msglist.0.msgsvrid").String()

	// switch srv.EventType {
	// case message.PCRecvTextMsgEvent:
	// case message.PCRecvGroupTextMsgEvent:
	// case message.PCRecvImgMsgEvent: // PC收到图片消息
	// case message.MsgTypeFileOrAppShareLinkFile: // "msgtype":"49",
	// 	if msgContent == "PC发app/文件消息成功" {
	// 		continue
	// 	}

	// case message.MsgTypeGif: // "msgtype":"47",
	// 	if msgContent == "PC发动态图片消息成功" {
	// 		continue
	// 	}
	// 	gif := message.Gif{
	// 		GifPath: interfaceToString(msgItem["gif_path"]),
	// 	}
	// 	wcMsgItem.MsgItem = gif

	// 	if fromtype == "1" {
	// 		wcMsgItem.EventType = message.PCRecvGifImgMsgEvent
	// 	} else if fromtype == "2" {
	// 		wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
	// 			FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 			FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		}
	// 		// wcMsg := message.GroupGif{
	// 		// 	Gif: gif,
	// 		// 	CommonGroupMsg: message.CommonGroupMsg{
	// 		// 		FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 		// 		FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		// 	},
	// 		// }
	// 		// wcMsgItem.MsgItem = wcMsg
	// 		wcMsgItem.EventType = message.PCRecvGroupGifImgMsgEvent
	// 	}
	// case message.MsgTypeVideo: // "msgtype":"43", PC收到视频消息

	// 	video := message.Video{
	// 		Info:      interfaceToString(msgItem["info"]), // 消息源内容
	// 		VideoPath: interfaceToString(msgItem["video_path"]),
	// 	}
	// 	wcMsgItem.MsgItem = video

	// 	if fromtype == "1" {
	// 		wcMsgItem.EventType = message.PCRecvVideoMsgEvent
	// 	} else if fromtype == "2" {
	// 		// wcMsg := message.GroupVideo{
	// 		// 	Video: video,
	// 		// 	CommonGroupMsg: message.CommonGroupMsg{
	// 		// 		FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 		// 		FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		// 	},
	// 		// }
	// 		// wcMsgItem.MsgItem = wcMsg
	// 		wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
	// 			FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 			FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		}
	// 		wcMsgItem.EventType = message.PCRecvGroupVideoMsgEvent
	// 	}
	// case message.MsgTypeVoice: // "msgtype":"34", PC收到语音消息
	// 	voice := message.Voice{
	// 		VoiceLen:  interfaceToString(msgItem["voice_len"]),
	// 		VoiceData: interfaceToString(msgItem["voice_data"]),
	// 		VoiceHex:  interfaceToString(msgItem["voice_hex"]),
	// 	}
	// 	wcMsgItem.MsgItem = voice

	// 	if fromtype == "1" {
	// 		wcMsgItem.EventType = message.PCRecvVoiceMsgEvent
	// 	} else if fromtype == "2" {
	// 		// wcMsg := message.GroupVoice{
	// 		// 	Voice: voice,
	// 		// 	CommonGroupMsg: message.CommonGroupMsg{
	// 		// 		FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 		// 		FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		// 	},
	// 		// }
	// 		// wcMsgItem.MsgItem = wcMsg
	// 		wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
	// 			FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 			FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		}
	// 		wcMsgItem.EventType = message.PCRecvGroupVoiceMsgEvent
	// 	}
	// case message.MsgTypeShareCard: // "msgtype":"42",
	// 	// PC收到名片消息
	// 	// eventType = message.PCRecvShareCardMsgEvent
	// case message.MsgTypeLocation: // "msgtype":"48",
	// 	// PC收到位置消息
	// 	// eventType = message.PCRecvLocationMsgEvent
	// case message.MsgTypeShareLocation: // "msgtype":"58",
	// 	// PC收到共享位置消息
	// 	// eventType = message.PCRecvShareLocationMsgEvent
	// case message.MsgTypeSystem: // "msgtype":"1000", 系统消息
	// 	// if msgItem.Msg == "位置共享结束" {
	// 	// PC收到共享位置结束消息
	// 	// eventType = message.PCRecvEndShareLocationMsgEvent
	// 	// }
	// case message.MsgTypeRecall: // "msgtype":"10002", 撤回消息
	// 	sysMsgXml := message.SysMsgXml{}
	// 	err = xml.Unmarshal([]byte(msgContent), &sysMsgXml)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if sysMsgXml.Type == "revokemsg" { // 撤回消息
	// 		revoke := message.Revoke{
	// 			RevokeMsg:  interfaceToString(msgItem["revoke_msg"]),
	// 			Session:    sysMsgXml.RevokeMsg.Session,
	// 			MsgId:      sysMsgXml.RevokeMsg.MsgId,
	// 			NewMsgId:   sysMsgXml.RevokeMsg.NewMsgId,
	// 			ReplaceMsg: sysMsgXml.RevokeMsg.ReplaceMsg,
	// 		}
	// 		wcMsgItem.MsgItem = revoke

	// 		if fromtype == "1" {
	// 			wcMsgItem.EventType = message.PCRecvRevokeMsgEvent
	// 		} else if fromtype == "2" {
	// 			// wcMsg := message.GroupRevoke{
	// 			// 	Revoke: revoke,
	// 			// 	CommonGroupMsg: message.CommonGroupMsg{
	// 			// 		FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 			// 		FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 			// 	},
	// 			// }
	// 			// wcMsgItem.MsgItem = wcMsg
	// 			wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
	// 				FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 				FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 			}
	// 			wcMsgItem.EventType = message.PCRecvGroupRevokeMsgEvent
	// 		}
	// 	}
	// }

	// var msg interface{}
	// msg, err := srv.getMessage()
	// if err != nil {
	// 	return err
	// }
	// msgBody, success := msg.(*message.MsgBody)
	// if !success {
	// 	err = errors.New("消息类型转换失败")
	// 	return err
	// }
	// srv.RequestMsg = msgBody

	// var msgItemList []message.WcMsgItem

	// // sendorrecv: 1=收到的消息,2=发送的消息
	// sendorrecv := msgBody.Sendorrecv

	// if sendorrecv != "2" {
	// 	return nil
	// }

	// msgList := msgBody.Msglist

	// if len(msgList) == 0 {
	// 	return errors.New("消息列表为空")
	// }

	// for _, msgItem := range msgList {

	// 	fromid := interfaceToString(msgItem["fromid"])

	// 	// log.Println(" ---------------- msgBody.SelfWxid:", msgBody.SelfWxid, " ---------------- ")
	// 	// log.Println(" ---------------- fromid:", fromid, " ---------------- ")

	// 	msgType := message.MsgType(interfaceToString(msgItem["msgtype"]))

	// 	fromtype := interfaceToString(msgItem["fromtype"])

	// 	// if msgBody.SelfWxid == fromid {
	// 	// 	log.Println(" ---------------- 自己发送的消息 ---------------- ")

	// 	// }

	// 	toid := interfaceToString(msgItem["toid"])

	// 	if fromtype == "1" && srv.isGroupMsg(toid) {
	// 		fromtype = "2"
	// 	}

	// 	msgContent := interfaceToString(msgItem["msg"])

	// 	wcMsgItem := message.WcMsgItem{
	// 		ServerPort: msgBody.ServerPort,
	// 		SelfWxid:   msgBody.SelfWxid,
	// 		CommonMsg: message.CommonMsg{
	// 			MsgSvrid: interfaceToString(msgItem["msgsvrid"]),
	// 			Time:     interfaceToString(msgItem["time"]),
	// 			Msg:      msgContent,
	// 			MsgType:  interfaceToString(msgItem["msgtype"]),  // 消息类型代码
	// 			FromType: fromtype,                               // 个人消息=1 群消息=2
	// 			FromId:   fromid,                                 // 发送方微信ID
	// 			FromName: interfaceToString(msgItem["fromname"]), // 发送方昵称
	// 			Index:    interfaceToString(msgItem["index"]),
	// 		},
	// 		ToCommonMsg: message.ToCommonMsg{
	// 			ToId:   toid,
	// 			ToName: interfaceToString(msgItem["toname"]),
	// 		},
	// 	}

	// switch msgType {
	// case message.MsgTypeText:
	// 	if msgContent == "PC发文本消息成功" {
	// 		continue
	// 	}

	// 	text := message.Text{
	// 		MsgSource: interfaceToString(msgItem["msgsource"]), // 消息源内容
	// 	}

	// 	wcMsgItem.MsgItem = text

	// 	if fromtype == "1" {
	// 		wcMsgItem.EventType = message.PCRecvTextMsgEvent
	// 	} else if fromtype == "2" {
	// 		wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
	// 			FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 			FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		}

	// 		// groupText := message.GroupText{
	// 		// 	Text: text, // 消息源内容
	// 		// 	CommonGroupMsg: message.CommonGroupMsg{
	// 		// 		FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 		// 		FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		// 	},
	// 		// }

	// 		msgSource := text.MsgSource

	// 		var msgSourceXml message.MsgSourceXml
	// 		// 字符串转换为xml
	// 		err = xml.Unmarshal([]byte(msgSource), &msgSourceXml)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		if msgSourceXml.AtUserList != "" {
	// 			wcMsgItem.AtText = message.AtText{
	// 				AtUserList: strings.Split(msgSourceXml.AtUserList, ","),
	// 			}
	// 			wcMsgItem.EventType = message.PCRecvAtTextMsgEvent
	// 		} else {
	// 			wcMsgItem.EventType = message.PCRecvGroupTextMsgEvent
	// 		}

	// 	}
	// case message.MsgTypeImage: // PC收到图片消息
	// 	if msgContent == "PC发图片消息成功" {
	// 		continue
	// 	}

	// 	image := message.Image{
	// 		Info:      interfaceToString(msgItem["info"]),       // 消息源内容
	// 		ImgLen:    interfaceToFloat64(msgItem["img_len"]),   // 消息源内容
	// 		ImgPath:   interfaceToString(msgItem["img_path"]),   // 消息源内容
	// 		ImgBase64: interfaceToString(msgItem["img_base64"]), // 消息源内容
	// 	}

	// 	wcMsgItem.MsgItem = image

	// 	if fromtype == "1" {
	// 		wcMsgItem.EventType = message.PCRecvImgMsgEvent
	// 	} else if fromtype == "2" {
	// 		// wcMsg := message.GroupImage{
	// 		// 	Image: image,
	// 		// 	CommonGroupMsg: message.CommonGroupMsg{
	// 		// 		FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 		// 		FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		// 	},
	// 		// }
	// 		// wcMsgItem.MsgItem = wcMsg
	// 		wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
	// 			FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 			FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		}
	// 		wcMsgItem.EventType = message.PCRecvGroupImgMsgEvent
	// 	}
	// 	// log.Println(" ---------------- PC收到图片消息处理结束 ---------------- ")
	// case message.MsgTypeFileOrAppShareLinkFile: // "msgtype":"49",
	// 	if msgContent == "PC发app/文件消息成功" {
	// 		continue
	// 	}

	// 	var appMsgXml message.AppMsgXml
	// 	// 字符串转换为xml
	// 	err = xml.Unmarshal([]byte(msgContent), &appMsgXml)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if appMsgXml.AppMsg.Type == "57" { // 引用消息

	// 		var quoteMsgXml message.QuoteMsgXml
	// 		// 字符串转换为xml
	// 		err = xml.Unmarshal([]byte(msgContent), &appMsgXml)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		quote := message.Quote{
	// 			MsgSource:    quoteMsgXml.AppMsg.ReferMsg.MsgSource, // 消息源内容
	// 			QuoteMsg:     quoteMsgXml.AppMsg.ReferMsg.Content,   // 引用的消息内容
	// 			QuoteMsgType: quoteMsgXml.AppMsg.ReferMsg.Type,      // 引用的消息类型
	// 			QuoteMsgId:   quoteMsgXml.AppMsg.ReferMsg.Svrid,     // 引用的消息id
	// 			ReplyMsg:     quoteMsgXml.AppMsg.Title,              // 回复的消息内容
	// 		}
	// 		wcMsgItem.MsgItem = quote

	// 		if fromtype == "1" {
	// 			wcMsgItem.EventType = message.PCRecvQuoteMsgEvent
	// 		} else if fromtype == "2" {
	// 			wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
	// 				FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 				FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 			}
	// 			// wcMsgItem.MsgItem = wcMsg
	// 			wcMsgItem.EventType = message.PCRecvGroupQuoteMsgEvent
	// 		}
	// 	} else if appMsgXml.AppMsg.Type == "51" { // 视频号消息

	// 		var channelsMsgXml message.ChannelsMsgXml
	// 		// 字符串转换为xml
	// 		err = xml.Unmarshal([]byte(msgContent), &channelsMsgXml)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		channelsMsg := message.ChannelsMsg{
	// 			ObjectId:      channelsMsgXml.AppMsg.FinderFeed.ObjectId,
	// 			ObjectNonceId: channelsMsgXml.AppMsg.FinderFeed.ObjectNonceId,
	// 		}
	// 		wcMsgItem.MsgItem = channelsMsg

	// 		if fromtype == "1" {
	// 			wcMsgItem.EventType = message.PCRecvChannelsMsgEvent
	// 		} else if fromtype == "2" {
	// 			wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
	// 				FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 				FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 			}
	// 			// wcMsgItem.MsgItem = wcMsg
	// 			wcMsgItem.EventType = message.PCRecvGroupChannelsMsgEvent
	// 		}

	// 	}
	// case message.MsgTypeGif: // "msgtype":"47",
	// 	if msgContent == "PC发动态图片消息成功" {
	// 		continue
	// 	}

	// case message.MsgTypeVideo: // "msgtype":"43", PC收到视频消息

	// 	video := message.Video{
	// 		Info:      interfaceToString(msgItem["info"]), // 消息源内容
	// 		VideoPath: interfaceToString(msgItem["video_path"]),
	// 	}
	// 	wcMsgItem.MsgItem = video

	// 	if fromtype == "1" {
	// 		wcMsgItem.EventType = message.PCRecvVideoMsgEvent
	// 	} else if fromtype == "2" {
	// 		// wcMsg := message.GroupVideo{
	// 		// 	Video: video,
	// 		// 	CommonGroupMsg: message.CommonGroupMsg{
	// 		// 		FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 		// 		FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		// 	},
	// 		// }
	// 		// wcMsgItem.MsgItem = wcMsg
	// 		wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
	// 			FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 			FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		}
	// 		wcMsgItem.EventType = message.PCRecvGroupVideoMsgEvent
	// 	}
	// case message.MsgTypeVoice: // "msgtype":"34", PC收到语音消息

	// 	if fromtype == "1" {
	// 		wcMsgItem.EventType = message.PCRecvVoiceMsgEvent
	// 	} else if fromtype == "2" {
	// 		// wcMsg := message.GroupVoice{
	// 		// 	Voice: voice,
	// 		// 	CommonGroupMsg: message.CommonGroupMsg{
	// 		// 		FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 		// 		FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		// 	},
	// 		// }
	// 		// wcMsgItem.MsgItem = wcMsg
	// 		wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
	// 			FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 			FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 		}
	// 		wcMsgItem.EventType = message.PCRecvGroupVoiceMsgEvent
	// 	}
	// case message.MsgTypeShareCard: // "msgtype":"42",
	// 	// PC收到名片消息
	// 	// eventType = message.PCRecvShareCardMsgEvent
	// case message.MsgTypeLocation: // "msgtype":"48",
	// 	// PC收到位置消息
	// 	// eventType = message.PCRecvLocationMsgEvent
	// case message.MsgTypeShareLocation: // "msgtype":"58",
	// 	// PC收到共享位置消息
	// 	// eventType = message.PCRecvShareLocationMsgEvent
	// case message.MsgTypeSystem: // "msgtype":"1000", 系统消息
	// 	// if msgItem.Msg == "位置共享结束" {
	// 	// PC收到共享位置结束消息
	// 	// eventType = message.PCRecvEndShareLocationMsgEvent
	// 	// }
	// case message.MsgTypeRecall: // "msgtype":"10002", 撤回消息
	// 	sysMsgXml := message.SysMsgXml{}
	// 	err = xml.Unmarshal([]byte(msgContent), &sysMsgXml)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if sysMsgXml.Type == "revokemsg" { // 撤回消息
	// 		revoke := message.Revoke{
	// 			RevokeMsg:  interfaceToString(msgItem["revoke_msg"]),
	// 			Session:    sysMsgXml.RevokeMsg.Session,
	// 			MsgId:      sysMsgXml.RevokeMsg.MsgId,
	// 			NewMsgId:   sysMsgXml.RevokeMsg.NewMsgId,
	// 			ReplaceMsg: sysMsgXml.RevokeMsg.ReplaceMsg,
	// 		}
	// 		wcMsgItem.MsgItem = revoke

	// 		if fromtype == "1" {
	// 			wcMsgItem.EventType = message.PCRecvRevokeMsgEvent
	// 		} else if fromtype == "2" {
	// 			// wcMsg := message.GroupRevoke{
	// 			// 	Revoke: revoke,
	// 			// 	CommonGroupMsg: message.CommonGroupMsg{
	// 			// 		FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 			// 		FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 			// 	},
	// 			// }
	// 			// wcMsgItem.MsgItem = wcMsg
	// 			wcMsgItem.CommonGroupMsg = message.CommonGroupMsg{
	// 				FromGname: interfaceToString(msgItem["fromgname"]), // 群名称
	// 				FromGid:   interfaceToString(msgItem["fromgid"]),   // 群ID
	// 			}
	// 			wcMsgItem.EventType = message.PCRecvGroupRevokeMsgEvent
	// 		}
	// 	}
	// }
	// 	msgItemList = append(msgItemList, wcMsgItem)

	// }

	return srv.messageHandler(msgItemList)
}

// 过去Emoji消息
func getEmojiMsg(emojiXmlStr string) (*message.EmojiMsg, error) {

	var msg message.EmojiMsg
	err := xml.Unmarshal([]byte(emojiXmlStr), &msg)
	if err != nil {
		fmt.Println("Error unmarshalling XML: ", err)
		return nil, err
	}

	fmt.Printf("Parsed Struct: %+v\n", msg)

	return &msg, nil
}

// 解码 Base64 字符串
func decodeString(base64String string) (string, error) {

	// 解码 Base64 字符串
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		fmt.Println("Error decoding string: ", err.Error())
		return "", err
	}

	return string(data), nil
}

func getDesc(desc string) (string, error) {

	// 正则表达式匹配 'zh_cn' 后的任意字符，直到 'zh_tw' 之前
	re := regexp.MustCompile(`zh_cn\s*(.*?)\s*zh_tw`)

	match := re.FindStringSubmatch(desc)
	// && len(match) > 1
	if match != nil {
		// 去除换行符和额外的空格
		cleanedText := strings.ReplaceAll(match[1], "\n", "")
		cleanedText = strings.TrimSpace(cleanedText)
		cleanedText = compressStr(cleanedText)

		// fmt.Println("Matched Text:", cleanedText)

		return cleanedText, nil

	} else {
		fmt.Println("No match found")

		return "", nil
	}
}

// 利用正则表达式压缩字符串，去除空格或制表符
func compressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(str, "")
}
