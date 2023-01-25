package model

type SendTextMsgReq struct {
	Wxid string `json:"wxid"`
	Msg  string `json:"msg"`
}

type SendTextMsgResult struct {
	SendTextMsg string `json:"SendTextMsg"`
}

type SendPicMsgReq struct {
	Wxid        string `json:"wxid"`
	PicPath     string `json:"picpath"`
	DiyFilename string `json:"diyfilename,omitempty"`
}

type SendPicMsgResult struct {
	SendPicMsg string `json:"SendPicMsg"`
}

type SendFileMsgReq struct {
	Wxid        string `json:"wxid"`
	Filepath    string `json:"filepath"`
	DiyFilename string `json:"diyfilename,omitempty"`
}

type SendFileMsgResult struct {
	SendFileMsg string `json:"SendFileMsg"`
}

type SendVideoMsgReq struct {
	Wxid        string `json:"wxid"`
	VideoPath   string `json:"videopath"`
	DiyFilename string `json:"diyfilename,omitempty"`
}

type SendVideoMsgResult struct {
	SendVideoMsg string `json:"SendVideoMsg"`
}

type SendGIFMsgReq struct {
	Wxid    string `json:"wxid"`
	GifPath string `json:"gifpath"`
}

type SendGIFMsgResult struct {
	SendGIFMsg string `json:"SendGIFMsg"`
}

type SendLocationMsgReq struct {
	Wxid string `json:"wxid"`
	Msg  string `json:"msg"`
}

type SendLocationMsgResult struct {
	SendLocationMsg string `json:"SendLocationMsg"`
}

type SendPatMsgReq struct {
	Wxid string `json:"wxid"`
	Gid  string `json:"gid"`
}

type SendPatMsgResult struct {
	SendPatMsg string `json:"SendPatMsg"`
}

type SendVoiceMsgReq struct {
	Wxid     string `json:"wxid"`
	VoiceHex string `json:"voice_hex"`
}

type SendVoiceMsgResult struct {
	SendVoiceMsg string `json:"SendVoiceMsg"`
}

type ConfigureMsgReciveFullURLReq struct {
	Url string `json:"url"`
}

type ConfigureMsgReciveFullURLResult struct {
	ConfigureMsgReciveFullURL string `json:"ConfigureMsgReciveFullURL"`
}

type ConfigureMsgReciveReq struct {
	IsEnable string `json:"isEnable"`
	Url      string `json:"url"`
}

type ConfigureMsgReciveResult struct {
	ConfigureMsgRecive string `json:"ConfigureMsgRecive"`
}

type CheckFriendStatusReq struct {
	Wxid string `json:"wxid"`
}

type CheckFriendStatusResult struct {
	CheckFriendStatus string `json:"CheckFriendStatus"`
}

type GetGIFURLReq struct {
	MsgXml string `json:"msg_xml"`
}

type GetGIFURLResult struct {
	GetGIFURL string `json:"GetGIFURL"`
}

type DownPicReq struct {
	ToPath string `json:"topath"`
	MsgXml string `json:"msg_xml"`
}

type DownPicResult struct {
	DownPic string `json:"DownPic"`
}

type GetUnReadMsgNumResult struct {
	GetUnReadMsgNum string `json:"GetUnReadMsgNum"`
}

type SendDIYMsgReq struct {
	Type string `json:"type"`
	Wxid string `json:"wxid"`
	Msg  string `json:"msg"`
}

type SendDIYMsgResult struct {
	SendDIYMsg string `json:"SendDIYMsg"`
}

type SendCardMsgReq struct {
	Towxid   string `json:"towxid"`
	Fromwxid string `json:"fromwxid"`
	Nickname string `json:"nickname"`
}

type SendCardMsgResult struct {
	SendCardMsg string `json:"SendCardMsg"`
}

type FowardEssayMsgReq struct {
	Towxid  string `json:"towxid"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	URL     string `json:"url"`
	Imgpath string `json:"imgpath"`
}

type FowardEssayMsgResult struct {
	FowardEssayMsg string `json:"FowardEssayMsg"`
}

type FowardAppMsgReq struct {
	Towxid  string `json:"towxid"`
	Imgpath string `json:"imgpath"`
	XML     string `json:"xml"`
}

type FowardAppMsgResult struct {
	FowardAppMsg string `json:"FowardAppMsg"`
}

type FowardMusicMsgReq struct {
	Towxid  string `json:"towxid"`
	Imgpath string `json:"imgpath"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	URL     string `json:"url"`
	Dataurl string `json:"dataurl"`
}

type FowardMusicMsgResult struct {
	FowardMusicMsg string `json:"FowardMusicMsg"`
}

type FowardXMLMsgReq struct {
	Type    string `json:"type"`
	Towxid  string `json:"towxid"`
	Imgpath string `json:"imgpath"`
	XML     string `json:"xml"`
}

type FowardXMLMsgResult struct {
	FowardXMLMsg string `json:"FowardXMLMsg"`
}

type SendAtMsgReq struct {
	Gid  string `json:"gid"`
	Wxid string `json:"wxid"`
	Msg  string `json:"msg"`
	Desc string `json:"desc"`
}

type SendAtMsgResult struct {
	SendAtMsg string `json:"SendAtMsg"`
}

type SendAtAllMsgReq struct {
	Gid string `json:"gid"`
	Msg string `json:"msg"`
}

type SendAtAllMsgResult struct {
	SendAtAllMsg string `json:"SendAtAllMsg"`
}

type ForwardAllMsgReq struct {
	Msgid string `json:"msgid"`
	Wxid  string `json:"wxid"`
}

type ForwardAllMsgResult struct {
	ForwardAllMsg string `json:"ForwardAllMsg"`
}
