package model

type SendTextMsgReq struct {
	Wxid string `json:"wxid"`
	Msg  string `json:"msg"`
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

type ConfigureMsgReciveFullURLReq struct {
	Url string `json:"url"`
}

type ConfigureMsgReciveFullURLResult struct {
	ConfigureMsgReciveFullURL string `json:"ConfigureMsgReciveFullURL"`
}

type ConfigureMsgReciveReq struct {
	IsEnable string `json:"isEnable"`
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
