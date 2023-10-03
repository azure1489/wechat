package sendmsgmanager

import (
	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/common"
)

type SendMsgManagerService interface {
	// SendTextMsg 发送文本消息 https://www.showdoc.com.cn/WeChatProject/8929112442643628
	SendTextMsg(wxid, msg string) error
	// SendTextMsgNoSrc 发送文本消息_无源 https://www.showdoc.com.cn/WeChatProject/9859692078921113
	SendTextMsgNoSrc(wxid, msg string) error
	// SendPicMsg 发送图片消息 https://www.showdoc.com.cn/WeChatProject/8929124214624404
	SendPicMsg(wxid, picPath, diyFilename string) error
	// SendImgMsgNoSrc 发送图片消息_无源,无需上传图片直接发送图片消息给对方 https://www.showdoc.com.cn/WeChatProject/9859693859851287
	SendImgMsgNoSrc(req SendImgMsgNoSrcReq) error
	// SendFileMsg 发送文件消息 https://www.showdoc.com.cn/WeChatProject/8929125290388797
	SendFileMsg(wxid, filepath, diyFilename string) error
	// SendVideoMsg 发送视频消息 https://www.showdoc.com.cn/WeChatProject/8929126402335432
	SendVideoMsg(wxid, videoPath, diyFilename string) error
	// SendGIFMsg 发送GIF动画表情消息 https://www.showdoc.com.cn/WeChatProject/8929127822471500
	SendGIFMsg(wxid, gifPath string) error
	// SendVoiceMsg 发送语音消息 https://www.showdoc.com.cn/WeChatProject/892913222925935
	SendVoiceMsg(wxid, voicePath string) error
	// SendDiyMsg 发送自定义消息 https://www.showdoc.com.cn/WeChatProject/8929124214624404
	SendDiyMsg(req SendDiyMsgReq) error
	// ForwardVoiceMsg 转发语音消息 https://www.showdoc.com.cn/WeChatProject/10042262554089499
	ForwardVoiceMsg(req ForwardVoiceMsgReq) (string, error)
	// SendPatMsg 发送拍一拍消息 https://www.showdoc.com.cn/WeChatProject/9644354859714189
	SendPatMsg(wxid, gid string) error
	// SendCardMsg 转发好友名片 https://www.showdoc.com.cn/WeChatProject/8929459103947241
	SendCardMsg(req SendCardMsgReq) error
	// SendQuoteMsg 发送引用消息 https://www.showdoc.com.cn/WeChatProject/10213975499481497
	SendQuoteMsg(req SendQuoteMsgReq) error
	// FowardXMLMsg 转发自定义XML消息 https://www.showdoc.com.cn/WeChatProject/9167364405161045
	FowardXMLMsg(req FowardXMLMsgReq) error
	// ForwardVideoMsg 转发视频号消息 https://www.showdoc.com.cn/WeChatProject/9252827740062809
	ForwardVideoMsg(req ForwardVideoMsgReq) error
	// SendAtMsg 发送群@消息 https://www.showdoc.com.cn/WeChatProject/8929229624237103
	SendAtMsg(req SendAtMsgReq) error
	// SendAtMsgNoSrc 发送群@消息_无源 https://www.showdoc.com.cn/WeChatProject/10214843464900880
	SendAtMsgNoSrc(req SendAtMsgNoSrcReq) (string, error)
	// SendAtAllMsg 发送群@全体成员消息 https://www.showdoc.com.cn/WeChatProject/8981408454846254
	SendAtAllMsg(req SendAtAllMsgReq) error
	// CheckFriendStatus 检测好友状态 https://www.showdoc.com.cn/WeChatProject/9063410601712380
	CheckFriendStatus(req CheckFriendStatusReq) (CheckFriendStatusResult, error)
}

type SendMsgManagerServiceImpl struct {
	http common.HttpClientService
}

func NewSendMsgManagerService(config *wechat.WechatConfig) SendMsgManagerService {

	httpClientService := common.NewHttpClientService(config.Ip, config.Port, config.Url, config.Secret, config.Timeout)

	return &SendMsgManagerServiceImpl{
		// config: config,
		http: httpClientService,
	}
}
