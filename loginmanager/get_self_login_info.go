package loginmanager

import "encoding/json"

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

// GetSelfLoginInfo 获取个人详细信息 https://www.showdoc.com.cn/WeChatProject/8929111706614173
func (l *LoginManagerServiceImpl) GetSelfLoginInfo(url string) (*GetSelfLoginInfoResult, error) {

	resultBody, err := l.http.DoPost("/GetSelfLoginInfo", nil)
	if err != nil {
		return nil, err
	}

	commonResult := GetSelfLoginInfoResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
