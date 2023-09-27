package loginmanager

import (
	"encoding/json"
)

//	{
//	    "onlinestatus": "0",
//	    "login_loading": "7%",
//	    "selfwxid": "wxid_123456",
//	    "nickname": "Heart"
//	}
type IsLoginStatusResult struct {
	OnlineStatus string `json:"onlinestatus"`
	LoginLoading string `json:"login_loading"`
	SelfWxid     string `json:"selfwxid"`
	NickName     string `json:"nickname"`
}

// IsLoginStatus 获取微信登陆状态 https://www.showdoc.com.cn/WeChatProject/8991915897471323
func (l *LoginManagerServiceImpl) IsLoginStatus() (*IsLoginStatusResult, error) {

	resultBody, err := l.http.DoPost("/IsLoginStatus", nil)
	if err != nil {
		return nil, err
	}

	commonResult := IsLoginStatusResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
