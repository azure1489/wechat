package loginmanager

import (
	"encoding/json"
	"fmt"
)

//	{
//	    "code": "0",
//	    "msg": "退出登录成功"
//	}
type LogoutResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// Logout 退出微信 https://www.showdoc.com.cn/WeChatProject/9008665345790557
func (l *LoginManagerServiceImpl) Logout() error {

	resultBody, err := l.http.DoPost("/Logout", nil)
	if err != nil {
		return err
	}

	commonResult := LogoutResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Code != "0" {
		return fmt.Errorf(commonResult.Msg+", body=%s", string(resultBody))
	}

	return nil
}
