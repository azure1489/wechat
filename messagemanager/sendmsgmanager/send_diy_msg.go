package sendmsmanager

import (
	"encoding/json"
	"fmt"
)

// type	是	string	类型
// wxid	是	string	对方微信ID/群ID
// msg	是	string	要发送的位置消息内容(xml内容,不能有换行，注意中间的引号用\转译！)
type SendDiyMsgReq struct {
	Type string `json:"type"`
	Wxid string `json:"wxid"`
	Msg  string `json:"msg"`
}

// type 二级代码	必选	类型	说明
// 1	是	string	文本消息
// 3	是	string	图片消息
// 34	是	string	语音消息
// 37	是	string	好友确认消息
// 40	是	string	POSSIBLEFRIEND_MSG
// 42	是	string	共享名片
// 43	是	string	视频消息
// 47	是	string	动画表情
// 48	是	string	位置消息
// 49	是	string	分享链接/文件
// 50	是	string	VOIPMSG
// 51	是	string	微信初始化消息
// 52	是	string	VOIPNOTIFY
// 53	是	string	VOIPINVITE
// 62	是	string	小视频
// 9999	是	string	SYSNOTICE
// 10000	是	string	系统消息
// 10002	是	string	撤回消息
type MessageType int

const (
	TextMessage          MessageType = 1
	ImageMessage         MessageType = 3
	VoiceMessage         MessageType = 34
	FriendConfirmation   MessageType = 37
	POSSIBLEFRIEND_MSG   MessageType = 40
	SharedContact        MessageType = 42
	VideoMessage         MessageType = 43
	AnimatedEmoji        MessageType = 47
	LocationMessage      MessageType = 48
	SharedLinkOrFile     MessageType = 49
	VOIPMSG              MessageType = 50
	WeChatInitialization MessageType = 51
	VOIPNOTIFY           MessageType = 52
	VOIPINVITE           MessageType = 53
	SmallVideo           MessageType = 62
	SYSNOTICE            MessageType = 9999
	SystemMessage        MessageType = 10000
	RevokeMessage        MessageType = 10002
)

//	{
//	    "success": "1"
//	  }
type SendDiyMsgResult struct {
	Success string `json:"success"`
}

// SendDiyMsg 发送自定义消息 https://www.showdoc.com.cn/WeChatProject/8929124214624404
func (l *SendMsgManagerServiceImpl) SendDiyMsg(req SendDiyMsgReq) error {

	resultBody, err := l.http.DoPost("/SendDiyMsg", req)
	if err != nil {
		return err
	}

	commonResult := SendDiyMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
