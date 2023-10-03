package contentmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	类型	说明
// ClearAllChatMsg	string	1
type ClearAllChatMsgResult struct {
	ClearAllChatMsg string `json:"ClearAllChatMsg"`
}

// ClearAllChatMsg 清空聊天记录 https://www.showdoc.com.cn/WeChatProject/9584482852044275
func (l *ContentMsgManagerServiceImpl) ClearAllChatMsg() error {

	resultBody, err := l.http.DoPost("/ClearAllChatMsg", nil)
	if err != nil {
		return err
	}

	commonResult := ClearAllChatMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ClearAllChatMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
