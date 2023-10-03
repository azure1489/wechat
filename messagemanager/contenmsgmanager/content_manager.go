package contentmsgmanager

import (
	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/common"
)

type ContentMsgManagerService interface {
	// AddToEmoji 方法_添加到自己表情 https://www.showdoc.com.cn/WeChatProject/9966833153857757
	AddToEmoji(filemd5 string) error
	// ChatRoomTopMsg 群置顶消息 https://www.showdoc.com.cn/WeChatProject/10273567687375653
	ChatRoomTopMsg(req ChatRoomTopMsgReq) error
	// ChatRoomUnTopMsg 取消置顶群聊消息 https://www.showdoc.com.cn/WeChatProject/10273569091488904
	ChatRoomUnTopMsg(req ChatRoomUnTopMsgReq) error
	// ClearAllChatMsg 清空聊天记录 https://www.showdoc.com.cn/WeChatProject/9584482852044275
	ClearAllChatMsg() error
	// DownFileorPic 下载图片_使用服务器消息ID https://www.showdoc.com.cn/WeChatProject/9364507487384330
	DownFileorPic(msgId string) error
	// DownPic 下载图片_使用消息内容 https://www.showdoc.com.cn/WeChatProject/9682774769136221
	DownPic(topath, msgXml string) error
	// DownPic4ID 下载图片_使用文件id和key https://www.showdoc.com.cn/WeChatProject/9745105666381249
	DownPic4ID(req DownPic4Req) error
	// DownVideo4ID 下载视频_使用文件id和key https://www.showdoc.com.cn/WeChatProject/9819195637839857
	DownVideo4ID(req DownVideo4IDReq) error
	// GetGIFURL 获取GIF的访问URL https://www.showdoc.com.cn/WeChatProject/9690133784680867
	GetGIFURL(msgXml string) (string, error)
	// MsgStruct 获取消息内容 https://www.showdoc.com.cn/WeChatProject/9364487555891575
	MsgStruct(msgId string) (MsgStructReust, error)
	// Translate 文本翻译 https://www.showdoc.com.cn/WeChatProject/9690133784680867
	Translate(text string) (string, error)
	// GetUnReadMsgNum 获取未读消息总数 https://www.showdoc.com.cn/WeChatProject/9690385454945788
	GetUnReadMsgNum() (string, error)
	// VoiceToText 语音转文本 https://www.showdoc.com.cn/WeChatProject/10074768335116363
	VoiceToText(req VoiceToTextReq) (string, error)
}

type ContentMsgManagerServiceImpl struct {
	http common.HttpClientService
}

func NewContentMsgManagerService(config *wechat.WechatConfig) ContentMsgManagerService {

	httpClientService := common.NewHttpClientService(config.Ip, config.Port, config.Url, config.Secret, config.Timeout)

	return &ContentMsgManagerServiceImpl{
		// config: config,
		http: httpClientService,
	}
}
