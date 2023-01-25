package model

type AgainStartWeChatReq struct {
	StartPort string `json:"StartPort"`
	WeChatPath string `json:"WeChatPath"`
	ProxyIP string `json:"Proxy_IP"`
	ProxyProt string `json:"Proxy_Prot"`
	ProxyUsr string `json:"Proxy_Usr"`
	ProxyPwd string `json:"Proxy_Pwd"`
}

type AgainStartWeChatResult struct {
	AgainStartWeChat string `json:"AgainStartWeChat"`
}
