package loginmanager

import (
	"encoding/json"
	"fmt"
)

type ClickLoginButtonResult struct {
	ClickLoginButton string `json:"ClickLoginButton"`
}

// ClickLoginButton 点击登陆微信 https://www.showdoc.com.cn/WeChatProject/10373442896174744
func (l *LoginManagerServiceImpl) ClickLoginButton() error {

	resultBody, err := l.http.DoPost("/ClickLoginButton", nil)
	if err != nil {
		return err
	}

	commonResult := ClickLoginButtonResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ClickLoginButton != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
