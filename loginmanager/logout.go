package loginmanager

import (
	"encoding/json"
	"fmt"
)

type LogoutResult struct {
	Success string `json:"success"`
}

// Logout 退出微信 https://www.showdoc.com.cn/WeChatProject/9008665345790557
func (l *LoginManagerServiceImpl) Logout(url string) error {

	resultBody, err := l.http.DoPost("/Logout", nil)
	if err != nil {
		return err
	}

	commonResult := LogoutResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
