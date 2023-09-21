package loginmanager

import (
	"encoding/json"
	"fmt"
)

type TerminateThisWeChatResult struct {
	Success string `json:"success"`
}

// TerminateThisWeChat 结束微信 https://www.showdoc.com.cn/WeChatProject/9214210657048561
func (l *LoginManagerServiceImpl) TerminateThisWeChat(url string) error {

	resultBody, err := l.http.DoPost("/TerminateThisWeChat", nil)
	if err != nil {
		return err
	}

	commonResult := TerminateThisWeChatResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
