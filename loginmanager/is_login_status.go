package loginmanager

import (
	"encoding/json"
	"fmt"
)

type IsLoginStatusResult struct {
	OnlineStatus string `json:"onlinestatus"`
}

// IsLoginStatus 获取微信登陆状态 https://www.showdoc.com.cn/WeChatProject/8991915897471323
func (l *LoginManagerServiceImpl) IsLoginStatus(url string) error {

	resultBody, err := l.http.DoPost("/IsLoginStatus", nil)
	if err != nil {
		return err
	}

	commonResult := IsLoginStatusResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.OnlineStatus != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
