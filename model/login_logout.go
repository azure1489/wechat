package model

type AgainStartWeChatReq struct {
	StartPort  string `json:"StartPort"`
	WeChatPath string `json:"WeChatPath"`
	ProxyIP    string `json:"Proxy_IP"`
	ProxyProt  string `json:"Proxy_Prot"`
	ProxyUsr   string `json:"Proxy_Usr"`
	ProxyPwd   string `json:"Proxy_Pwd"`
}

type AgainStartWeChatResult struct {
	AgainStartWeChat string `json:"AgainStartWeChat"`
}

type GetLoginQRCodeResult struct {
	HexData string `json:"hex_data"`
}

type IsLoginStatusResult struct {
	OnlineStatus string `json:"onlinestatus"`
}

type GetSelfLoginInfoResult struct {
	ProcessID  string `json:"ProcessID"`
	Wxid       string `json:"wxid"`
	Account    string `json:"account"`
	Nickname   string `json:"nickname"`
	Qq         string `json:"qq"`
	Tel        string `json:"tel"`
	Email      string `json:"email"`
	Country    string `json:"country"`
	Province   string `json:"province"`
	City       string `json:"city"`
	HeadSmall  string `json:"head_small"`
	HeadBig    string `json:"head_big"`
	DiySign    string `json:"diy_sign"`
	TimelineBg string `json:"timeline_bg"`
}

type LogoutResult struct {
	Success string `json:"success"`
}

type TerminateThisWeChatResult struct {
	Success string `json:"success"`
}
