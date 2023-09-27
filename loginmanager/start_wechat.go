package loginmanager

import (
	"encoding/json"
	"fmt"
)

type StartWechatReq struct {
	StartPort  string `json:"StartPort"`
	WeChatPath string `json:"WeChatPath"`
	ProxyIP    string `json:"Proxy_IP"`
	ProxyProt  string `json:"Proxy_Prot"`
	ProxyUsr   string `json:"Proxy_Usr"`
	ProxyPwd   string `json:"Proxy_Pwd"`
	PriCfgPath string `json:"Pri_Cfg_Path"`
}

type StartWechatResult struct {
	StartPort  string `json:"StartPort"`
	WeChatPath string `json:"WeChatPath"`
	ProxyIP    string `json:"Proxy_IP"`
	ProxyProt  string `json:"Proxy_Prot"`
	ProxyUsr   string `json:"Proxy_Usr"`
	ProxyPwd   string `json:"Proxy_Pwd"`
	PriCfgPath string `json:"Pri_Cfg_Path"`
	Success    string `json:"success"`
}

// AgainStartWeChat 启动更多微信 https://www.showdoc.com.cn/WeChatProject/9063540299207309
func (l *LoginManagerServiceImpl) StartWechat(req *StartWechatReq) error {
	resultBody, err := l.http.DoPost("/StartWechat", req)
	if err != nil {
		return err
	}

	commonResult := StartWechatResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
