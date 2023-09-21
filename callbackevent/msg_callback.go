package callbackevent

const (
	// ------------- PC端发送消息事件 -------------
	// PC端发送文本消息事件
	PCSendTextMsgEvent = "PCSendTextMsgEvent"
	// PC端发送文本消息成功事件
	PCSendTextMsgSuccessEvent = "PCSendTextMsgSuccessEvent"
	// PC端发送图片消息事件
	PCSendImgMsgEvent = "PCSendImgMsgEvent"
	// PC端发送图片消息成功事件
	PCSendImgMsgSuccessEvent = "PCSendImgMsgSuccessEvent"
	// PC发出文件/app消息事件
	PCSendFileOrAppMsgEvent = "PCSendFileOrAppMsgEvent"
	// PC发出文件/app消息成功事件
	PCSendFileOrAppMsgSuccessEvent = "PCSendFileOrAppMsgSuccessEvent"
	// PC发出视频消息事件
	PCSendVideoMsgEvent = "PCSendVideoMsgEvent"
	// PC发出视频消息成功事件
	PCSendVideoMsgSuccessEvent = "PCSendVideoMsgSuccessEvent"

	// ------------- PC端接收消息事件 -------------
	// PC端接收文本消息事件
	PCRecvTextMsgEvent = "PCRecvTextMsgEvent"
	// PC端接收图片消息事件
	PCRecvImgMsgEvent = "PCRecvImgMsgEvent"
	// PC端收到的消息(GIF动图)事件
	PCRecvGifImgMsgEvent = "PCRecvGifImgMsgEvent"
	// PC端收到的消息(语音消息)事件
	PCRecvVoiceMsgEvent = "PCRecvVoiceMsgEvent"
	// PC端收到的消息(视频消息)事件
	PCRecvVideoMsgEvent = "PCRecvVideoMsgEvent"
	// PC端拉人进群通知
	PCInviteInGroupEvent = "PCInviteInGroupEvent"
	// 登录二维码刷新事件
	PCLoginQrcodeRefreshEvent = "PCLoginQrcodeRefreshEvent"
	// 登陆微信事件
	PCLoginWxEvent = "PCLoginWxEvent"
	// 退出登陆微信事件
	PCLogoutWxEvent = "PCLogoutWxEvent"
	// 切换聊天对象事件
	PCSwitchChatObjectEvent = "PCSwitchChatObjectEvent"
	// 切换联系人事件
	PCSwitchContactEvent = "PCSwitchContactEvent"
	// 收款码到账事件
	PCRecvTransferMoneyEvent = "PCRecvTransferMoneyEvent"
	// 撤回消息事件
	PCRecallMsgEvent = "PCRecallMsgEvent"
	// 群成员增加事件
	PCGroupMemberAddEvent = "PCGroupMemberAddEvent"
	// 朋友圈被点赞事件
	PCFriendCircleLikeEvent = "PCFriendCircleLikeEvent"
)

type MsgResult struct {
	ServerPort string    `json:"ServerPort"`
	Selfwxid   string    `json:"selfwxid"`
	Sendorrecv string    `json:"sendorrecv"`
	Msgnumber  string    `json:"msgnumber"`
	Msglist    []MsgItem `json:"msglist"`
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
type MsgItem struct {
	Time      string `json:"time"`
	Msgtype   string `json:"msgtype"`
	Fromtype  string `json:"fromtype"`
	Fromwxid  string `json:"fromwxid"`
	Toid      string `json:"toid"`
	Msg       string `json:"msg"`
	Msgsvrid  string `json:"msgsvrid"`
	ImgPath   string `json:"img_path"`
	GifPath   string `json:"gif_path"`
	FilePath  string `json:"file_path"`
	VideoPath string `json:"video_path"`
	VoipData  string `json:"voip_data"`
	// ----------------- 其他 -----------------
	Device       string `json:"device"`
	Index        string `json:"index"`
	RevokeMsg    string `json:"revoke_msg"`
	Appletcode   string `json:"appletcode" `
	Businesstype string `json:"businesstype"`
	Businessdata string `json:"businessdata"`
}

// PC端 发送的消息

// PC发出文本消息
// 第一次 发送的内容
//
//	{
//	    "ServerPort":"30001",
//	    "selfwxid":"themid",
//	    "sendorrecv":"1",
//	    "msgnumber":"1",
//	    "msglist":[{
//	            "index":"1",
//	            "time":"2023-03-05 13:07:54",
//	            "msgtype":"1",
//	            "msgsvrid":"",
//	            "msg":"大佬666~",
//	            "fromwxid":"themid",
//	            "toid":"filehelper"
//	        }]
//	}
type PCSendTextMsgResult struct {
	ServerPort string `json:"ServerPort"`
	Selfwxid   string `json:"selfwxid"`
	Sendorrecv string `json:"sendorrecv"`
	Msgnumber  string `json:"msgnumber"`
	Msglist    []struct {
		Index    string `json:"index"`
		Time     string `json:"time"`
		Msgtype  string `json:"msgtype"`
		Msgsvrid string `json:"msgsvrid"`
		Msg      string `json:"msg"`
		Fromwxid string `json:"fromwxid"`
		Toid     string `json:"toid"`
	} `json:"msglist"`
}

// 第二次 发送成功后 收到 服务器返回的消息ID (msgid)
//
//	{
//	    "ServerPort":"30001",
//	    "selfwxid":"themid",
//	    "sendorrecv":"2",
//	    "msgnumber":"1",
//	    "msglist":[{
//	            "index":"1",
//	            "time":"2023-03-05 13:07:54",
//	            "msgtype":"1",
//	            "msgsvrid":"7029814402633277009",
//	            "msg":"PC发文本消息成功",
//	            "fromwxid":"themid",
//	            "toid":"filehelper"
//	        }]
//	}
type PCSendTextMsgSuccessResult struct {
	ServerPort string `json:"ServerPort"`
	Selfwxid   string `json:"selfwxid"`
	Sendorrecv string `json:"sendorrecv"`
	Msgnumber  string `json:"msgnumber"`
	Msglist    []struct {
		Index    string `json:"index"`
		Time     string `json:"time"`
		Msgtype  string `json:"msgtype"`
		Msgsvrid string `json:"msgsvrid"`
		Msg      string `json:"msg"`
		Fromwxid string `json:"fromwxid"`
		Toid     string `json:"toid"`
	} `json:"msglist"`
}

// PC发出图片消息
// 第一次 发送的内容
//
//	{
//	    "ServerPort":"30001",
//	    "selfwxid":"themid",
//	    "sendorrecv":"1",
//	    "msgnumber":"0",
//	    "msglist":[{
//	            "index":"1",
//	            "time":"2023-03-05 13:09:19",
//	            "msgtype":"3",
//	            "msgsvrid":"",
//	            "msg":"8bb306d0ad78a2aa8cafe229dc304f38",
//	            "fromwxid":"themid",
//	            "toid":"filehelper",
//	            "img_path":"C:\\\\Users\\\\Administrator\\\\Documents\\\\WeChat Files\\\\thexed\\\\FileStorage\\\\Temp\\\\1677992958058.jpg"
//	        }]
//	}
type PCSendImgMsgResult struct {
	ServerPort string `json:"ServerPort"`
	Selfwxid   string `json:"selfwxid"`
	Sendorrecv string `json:"sendorrecv"`
	Msgnumber  string `json:"msgnumber"`
	Msglist    []struct {
		Index    string `json:"index"`
		Time     string `json:"time"`
		Msgtype  string `json:"msgtype"`
		Msgsvrid string `json:"msgsvrid"`
		Msg      string `json:"msg"`
		Fromwxid string `json:"fromwxid"`
		Toid     string `json:"toid"`
		ImgPath  string `json:"img_path"`
	} `json:"msglist"`
}

// 第二次 发送成功后 收到 服务器返回的消息ID (msgid)
//
//	{
//	    "ServerPort":"30001",
//	    "selfwxid":"themid",
//	    "sendorrecv":"2",
//	    "msgnumber":"1",
//	    "msglist":[{
//	            "index":"1",
//	            "time":"2023-03-05 13:09:20",
//	            "msgtype":"3",
//	            "msgsvrid":"8676825236427018921",
//	            "msg":"PC发图片消息成功",
//	            "fromwxid":"themid",
//	            "toid":"filehelper"
//	        }]
//	}
type PCSendImgMsgSuccessResult struct {
	ServerPort string `json:"ServerPort"`
	Selfwxid   string `json:"selfwxid"`
	Sendorrecv string `json:"sendorrecv"`
	Msgnumber  string `json:"msgnumber"`
	Msglist    []struct {
		Index    string `json:"index"`
		Time     string `json:"time"`
		Msgtype  string `json:"msgtype"`
		Msgsvrid string `json:"msgsvrid"`
		Msg      string `json:"msg"`
		Fromwxid string `json:"fromwxid"`
		Toid     string `json:"toid"`
	} `json:"msglist"`
}
