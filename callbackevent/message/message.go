package message

// MsgType 基本消息类型
type MsgType string

// EventType 事件类型
type EventType string

// InfoType 消息内容类型
type InfoType string

const (
	// ------------- PC端发送消息事件 -------------
	// PC端发送文本消息事件
	PCSendTextMsgEvent EventType = "PCSendTextMsgEvent"
	// PC端发送文本消息成功事件
	PCSendTextMsgSuccessEvent EventType = "PCSendTextMsgSuccessEvent"
	// PC端发送图片消息事件
	PCSendImgMsgEvent EventType = "PCSendImgMsgEvent"
	// PC端发送图片消息成功事件
	PCSendImgMsgSuccessEvent EventType = "PCSendImgMsgSuccessEvent"
	// PC发出文件/app消息事件
	PCSendFileOrAppMsgEvent EventType = "PCSendFileOrAppMsgEvent"
	// PC发出文件/app消息成功事件
	PCSendFileOrAppMsgSuccessEvent EventType = "PCSendFileOrAppMsgSuccessEvent"
	// PC发动态图片消息
	PCSendGifImgMsgEvent EventType = "PCSendGifImgMsgEvent"
	// PC发动态图片消息成功
	PCSendGifImgMsgSuccessEvent EventType = "PCSendGifImgMsgSuccessEvent"
	// PC发出视频消息事件
	PCSendVideoMsgEvent EventType = "PCSendVideoMsgEvent"
	// PC发出视频消息成功事件
	PCSendVideoMsgSuccessEvent EventType = "PCSendVideoMsgSuccessEvent"

	// ------------- PC端接收消息事件 -------------
	// PC端接收文本消息事件
	PCRecvTextMsgEvent EventType = "PCRecvTextMsgEvent"
	// PC端接收图片消息事件
	PCRecvImgMsgEvent EventType = "PCRecvImgMsgEvent"
	// PC端收到的消息(GIF动图)事件
	PCRecvGifImgMsgEvent EventType = "PCRecvGifImgMsgEvent"
	// PC端收到的消息(语音消息)事件
	PCRecvVoiceMsgEvent EventType = "PCRecvVoiceMsgEvent"
	// PC端收到的消息(视频消息)事件
	PCRecvVideoMsgEvent EventType = "PCRecvVideoMsgEvent"
	// PC发app/文件消息成功事件
	PCRecvFileOrAppMsgEvent EventType = "PCRecvFileOrAppMsgEvent"
	// PC端拉人进群通知
	PCInviteInGroupEvent EventType = "PCInviteInGroupEvent"
	// PC端收到名片消息
	PCRecvShareCardMsgEvent EventType = "PCRecvShareCardMsgEvent"
	// PC端收到位置消息
	PCRecvLocationMsgEvent EventType = "PCRecvLocationMsgEvent"

	// PC端收到开始共享位置消息
	// PCRecvStartShareLocationMsgEvent EventType = "PCRecvStartShareLocationMsgEvent"
	// PC端收到结束共享位置消息
	PCRecvEndShareLocationMsgEvent EventType = "PCRecvEndShareLocationMsgEvent"
	// PC端收到共享位置消息
	PCRecvShareLocationMsgEvent EventType = "PCRecvShareLocationMsgEvent"

	// 登录二维码刷新事件
	PCLoginQrcodeRefreshEvent EventType = "PCLoginQrcodeRefreshEvent"
	// 登陆微信事件
	PCLoginWxEvent EventType = "PCLoginWxEvent"
	// 退出登陆微信事件
	PCLogoutWxEvent EventType = "PCLogoutWxEvent"
	// 切换聊天对象事件
	PCSwitchChatObjectEvent EventType = "PCSwitchChatObjectEvent"
	// 切换联系人事件
	PCSwitchContactEvent EventType = "PCSwitchContactEvent"
	// 收款码到账事件
	PCRecvTransferMoneyEvent EventType = "PCRecvTransferMoneyEvent"
	// 撤回消息事件
	PCRecallMsgEvent EventType = "PCRecallMsgEvent"
	// 群成员增加事件
	PCGroupMemberAddEvent EventType = "PCGroupMemberAddEvent"
	// 朋友圈被点赞事件
	PCFriendCircleLikeEvent EventType = "PCFriendCircleLikeEvent"
	// 未知事件
	UnknownEvent EventType = "UnknownEvent"
)

// 1 文本消息
// 3 图片消息
// 34 语音消息
// 37 好友确认消息
// 40 好友推荐消息
// 42 共享名片
// 43 视频消息
// 44 主动撤回
// 47 动画表情
// 48 位置消息
// 49 APP分享链接/文件
// 50 VOIP消息
// 51 微信初始化消息
// 52 VOIP结束消息
// 53 VOIP邀请
// 62 小视频
// 9999 SYSNOTICE
// 10000 系统消息
// 10002 撤回消息
// 9995 登陆二维码刷新事件
// 9996 登陆微信事件
// 9997 退出登陆微信
// 9998 切换聊天对象
// 9994 切换联系人
const (
	// MsgTypeText 表示文本消息
	MsgTypeText MsgType = "1"
	// MsgTypeImage 表示图片消息
	MsgTypeImage MsgType = "3"
	// MsgTypeVoice 表示语音消息
	MsgTypeVoice MsgType = "34"
	// MsgTypeFriendConfirmation 表示好友确认消息
	MsgTypeFriendConfirmation MsgType = "37"
	// MsgTypeFriendRecommendation 表示好友推荐消息
	MsgTypeFriendRecommendation MsgType = "40"
	// MsgTypeShareCard 表示共享名片
	MsgTypeShareCard MsgType = "42"
	// MsgTypeVideo 表示视频消息
	MsgTypeVideo MsgType = "43"
	// MsgTypeActiveRecall 表示主动撤回
	MsgTypeActiveRecall MsgType = "44"
	// MsgTypeGif 表示动画表情
	MsgTypeGif MsgType = "47"
	// MsgTypeLocation 表示位置消息
	MsgTypeLocation MsgType = "48"
	//  表示共享位置消息
	MsgTypeShareLocation MsgType = "56"
	// MsgTypeFileOrAppShareLinkFile 表示APP分享链接/文件
	MsgTypeFileOrAppShareLinkFile MsgType = "49"
	// MsgTypeVOIP 表示VOIP消息
	MsgTypeVOIP MsgType = "50"
	// MsgTypeWeChatInitialization 表示微信初始化消息
	MsgTypeWeChatInitialization MsgType = "51"
	// MsgTypeVOIPEnd 表示VOIP结束消息
	MsgTypeVOIPEnd MsgType = "52"
	// MsgTypeVOIPInvitation 表示VOIP邀请
	MsgTypeVOIPInvitation MsgType = "53"
	// MsgTypeSmallVideo 表示小视频
	MsgTypeSmallVideo MsgType = "62"
	// MsgTypeSYSNOTICE 表示SYSNOTICE
	MsgTypeSYSNOTICE MsgType = "9999"
	// MsgTypeSystem 表示系统消息
	MsgTypeSystem MsgType = "10000"
	// MsgTypeRecall 表示撤回消息
	MsgTypeRecall MsgType = "10002"
	// MsgTypeLoginQRCodeRefreshEvent 表示登陆二维码刷新事件
	MsgTypeLoginQRCodeRefreshEvent MsgType = "9995"
	// MsgTypeLoginWeChatEvent 表示登陆微信事件
	MsgTypeLoginWeChatEvent MsgType = "9996"
	// MsgTypeLogoutWeChatEvent 表示退出登陆微信
	MsgTypeLogoutWeChatEvent MsgType = "9997"
	// MsgTypeSwitchChatObject 表示切换聊天对象
	MsgTypeSwitchChatObject MsgType = "9998"
	// MsgTypeSwitchContact 表示切换联系人
	MsgTypeSwitchContact MsgType = "9994"
)

type MsgBody struct {
	ServerPort string      `json:"ServerPort"`
	SelfWxid   string      `json:"selfwxid"`
	Sendorrecv string      `json:"sendorrecv"`
	MsgNumber  string      `json:"msgnumber"`
	Msglist    []CommonMsg `json:"msglist"`
	CommonMsg
}

// 参数名	必选	类型	说明
// businesstype	是	string	业务类型代码
// businessdata	是	string	业务数据(16进制protobuf文本)
// time	是	string	收到消息的时间
// msgtype	是	string	消息类型代码
// device	是	string	设备代码 手机=1 其他=0
// fromtype	是	string	个人消息=1 群消息=2
// fromwxid	是	string	发送方微信ID
// toid	是	string	接收人微信ID/群ID
// msg	是	string	消息内容
// msgsvrid	是	string	服务器消息ID,用于撤回，或者下载图片/视频/文件
// img_path	是	string	如果当前收到的是图片，则会显示图片路径图片为.dat格式
// gif_path	是	string	如果当前收到的是gif表情，则会显示gif路径
// file_path	是	string	如果当前收到的是文件，则会显示文件路径
// video_path	是	string	如果当前收到的是视频，则会显示视频路径
// voip_data	是	string	如果当前收到的是语音，则会显示语音数据。格式为16进制文本的.silk文件
// revoke_msg	是	string	如果当前收到的是撤回的消息，则会显示撤回的内容
// appletcode	否	string	返回小程序code，如果调用了GetMiniAppCode_Sync接口则会返回该数据

// "time":"2023-09-27 15:48:23",
// "msgtype":"1",
// "msgsvrid":"9168006379603397754",
// "msg":"C",
// "msgsource":"<msgsource>\\n\\t<silence>1</silence>\\n\\t<membercount>192</membercount>\\n\\t<signature>v1_vy+YqGvH</signature>\\n\\t<tmp_node>\\n\\t\\t<publisher-id></publisher-id>\\n\\t</tmp_node>\\n</msgsource>\\n",
// "fromtype":"2",
// "fromgid":"25009802945@chatroom",
// "fromgname":"",
// "fromid":"wxid_nuy8ayquaxox22",
// "fromname":"",
// "toid":"azure1489",
// "toname":""

// "msgtype":"9998",
// "msgsvrid":"",
// "msg":"切换聊天对象",
// "fromtype":"2",
// "fromid":"5872694777@chatroom",
// "fromgname":"test",

type MsgItem struct {
	Time      string `json:"time"`      // 收到消息的时间
	Msgtype   string `json:"msgtype"`   // 消息类型代码
	Msgsvrid  string `json:"msgsvrid"`  // 服务器消息ID,用于撤回，或者下载图片/视频/文件
	Msg       string `json:"msg"`       // 消息内容
	MsgSource string `json:"msgsource"` // 消息源内容
	Fromtype  string `json:"fromtype"`  // 个人消息=1 群消息=2
	FromId    string `json:"fromid"`    // 发送方微信ID
	FromName  string `json:"fromname"`  // 发送方昵称
	// ----------------- 群消息 -----------------
	FromGname string `json:"fromgname"` // 群名称
	FromGid   string `json:"fromgid"`   // 群ID
	// ----------------- 接收人 -----------------
	ToId   string `json:"toid"`   // 接收人微信ID/群ID
	ToName string `json:"toname"` // 接收人昵称
	// ----------------- 文件路径 -----------------
	ImgPath string `json:"img_path"`
	// ----------------- "msgtype":"47" GIF动图 -----------------
	GifPath string `json:"gif_path"`
	// ----------------- "msgtype":"49" 文件/app消息 -----------------
	AppName  string `json:"appname"` // 应用名称
	FilePath string `json:"file_path"`

	VideoPath string `json:"video_path"`
	VoipData  string `json:"voip_data"`
	// ----------------- 切换联系人事件 -----------------
	Nickname   string `json:"nickname"`
	V3         string `json:"v3"`
	Label      string `json:"label"`
	Head       string `json:"head"`
	Timelinebg string `json:"timelinebg"`
	Country    string `json:"country"`
	// ----------------- 其他 -----------------
	Index     string `json:"index"`
	RevokeMsg string `json:"revoke_msg"`

	Appletcode   string `json:"appletcode" `
	Businesstype string `json:"businesstype"`
	Businessdata string `json:"businessdata"`

	// ----------------- 其他事件 -----------------
	SelfAccount  string `json:"selfaccount"`
	QRCodeBase64 string `json:"qrcode_base64"`
	// ----------------- 其他 -----------------
	MsgType   MsgType   `json:"-"`
	EventType EventType `json:"-"`
}

type CommonMsg struct {
	Time     string `json:"time"`     // 收到消息的时间
	MsgType  string `json:"msgtype"`  // 消息类型代码
	MsgSvrid string `json:"msgsvrid"` // 服务器消息ID,用于撤回，或者下载图片/视频/文件
	Msg      string `json:"msg"`      // 消息内容
	FromType string `json:"fromtype"` // 个人消息=1 群消息=2
	FromId   string `json:"fromid"`   // 发送方微信ID
	FromName string `json:"fromname"` // 发送方昵称
	Index    string `json:"index"`
}

type CommonGroupMsg struct {
	FromGname string `json:"fromgname"` // 群名称
	FromGid   string `json:"fromgid"`   // 群ID
}

type ToCommonMsg struct {
	ToId   string `json:"toid"`   // 接收人微信ID/群ID
	ToName string `json:"toname"` // 接收人昵称
}
