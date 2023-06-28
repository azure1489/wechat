package model

type SendTextMsgReq struct {
	Wxid string `json:"wxid"`
	Msg  string `json:"msg"`
}

type SendTextMsgResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SendTextMsgNoSrcReq struct {
	WxidOrGid string `json:"wxidorgid"`
	Msg       string `json:"msg"`
}

type SendTextMsgNoSrcResult struct {
	SendTextMsgNoSrc string `json:"SendTextMsg_NoSrc"`
}

type SendPicMsgReq struct {
	Wxid        string `json:"wxid"`
	PicPath     string `json:"picpath"`
	DiyFilename string `json:"diyfilename,omitempty"`
}

type SendPicMsgResult struct {
	SendPicMsg string `json:"SendPicMsg"`
}

type SendImgMsgNoSrcReq struct {
	Wxidorgid string `json:"wxidorgid"`
	Fileid    string `json:"fileid"`
	Authkey   string `json:"authkey"`
	Filemd5   string `json:"filemd5"`
	Filesize  string `json:"filesize"`
	Filecrc32 string `json:"filecrc32"`
}

type SendImgMsgNoSrcResult struct {
	SendImgMsgNoSrc string `json:"SendImgMsg_NoSrc"`
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
	Wxid      string `json:"wxid"`
	VoiceFile string `json:"voice_file"`
	TimeMs    int    `json:"time_ms"`
}

type SendVoiceMsgResult struct {
	MsgSvrID string `json:"MsgSvrID"`
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

type DecodePicReq struct {
	OriPath  string `json:"oripath"`
	SavePath string `json:"savepath"`
}

type DecodePicResult struct {
	DecodePic string `json:"DecodePic"`
}

type CollectionReq struct {
	Fromwxid   string `json:"fromwxid"`
	Transferid string `json:"transferid"`
}

type CollectionResult struct {
	Collection string `json:"Collection"`
}

type UnCollectionReq struct {
	Fromwxid   string `json:"fromwxid"`
	Transferid string `json:"transferid"`
}

type UnCollectionResult struct {
	UnCollection string `json:"UnCollection"`
}

type GetMsgStructReq struct {
	MsgId string `json:"msg_id"`
}

type GetMsgStructResult struct {
	Msgtype string `json:"msgtype"`
	Towxid  string `json:"towxid"`
	Content string `json:"content"`
}

type DownPic4IDReq struct {
	Topath string `json:"topath"`
	Aeskey string `json:"aeskey"`
	Fileid string `json:"fileid"`
}

type DownPic4IDResult struct {
	DownPic4ID string `json:"DownPic4ID"`
}

type DownFileorPicReq struct {
	MsgId string `json:"msg_id"`
}

type DownFileorPicResult struct {
	DownFileorPic string `json:"DownFileorPic"`
}

type ClearAllChatMsgResult struct {
	ClearAllChatMsg string `json:"ClearAllChatMsg"`
}

type CheckFriendStatusSyncReq struct {
	WxidList string `json:"wxidlist"`
}

type CheckFriendStatusSyncResult struct {
	Code   string `json:"code"`
	IsSync string `json:"is_sync"`
}
