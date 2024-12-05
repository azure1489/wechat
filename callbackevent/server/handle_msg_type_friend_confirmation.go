package server

import (
	"encoding/xml"
	"fmt"

	"github.com/azure1489/wechat/callbackevent/message"
	"github.com/tidwall/gjson"
)

// handleMsgTypeFriendConfirmation 处理好友确认消息
func (srv *Server) handleMsgTypeFriendConfirmation(wcMsgItem *message.WcMsgItem, msgInfo gjson.Result) error {
	// "msgtype":"37", PC收到好友确认消息
	wcMsgItem.EventType = message.PCRecvFriendConfirmationMsgEvent

	msgContent := wcMsgItem.Msg

	var friendConfirmationMsgXml message.FriendConfirmationMsgXml
	// 字符串转换为xml
	err := xml.Unmarshal([]byte(msgContent), &friendConfirmationMsgXml)
	if err != nil {
		fmt.Printf("xml.Unmarshal err: %v\n", err)
		return err
	}

	friendConfirmation := message.FriendConfirmation{
		V3:           friendConfirmationMsgXml.EncryptUserName,
		V4:           friendConfirmationMsgXml.Ticket,
		Content:      friendConfirmationMsgXml.Content,
		FromUserName: friendConfirmationMsgXml.FromUserName,
		FromNickName: friendConfirmationMsgXml.FromNickName,
		HeadImgURL:   friendConfirmationMsgXml.BigHeadImgURL,
	}

	wcMsgItem.MsgItem = friendConfirmation
	return nil
}
