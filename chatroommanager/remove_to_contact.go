package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type RemoveToContactReq struct {
	GidOrWxid string `json:"gidorwxid"`
}

type RemoveToContactResult struct {
	Success string `json:"success"`
}

// RemoveToContact 移出通讯录 https://www.showdoc.com.cn/WeChatProject/9078598797893485
func (l *ChatRoomManagerServiceImpl) RemoveToContact(gidorwxid string) error {

	req := RemoveToContactReq{
		GidOrWxid: gidorwxid,
	}

	resultBody, err := l.http.DoPost("/RemoveToContact", req)
	if err != nil {
		return err
	}

	commonResult := RemoveToContactResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
