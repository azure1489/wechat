package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type SaveToContactReq struct {
	GidOrWxid string `json:"gidorwxid"`
}

type SaveToContactResult struct {
	Success string `json:"success"`
}

// SaveToContact 保存到通讯录 https://www.showdoc.com.cn/WeChatProject/9078597995448349
func (l *ChatRoomManagerServiceImpl) SaveToContact(gidorwxid string) error {

	req := SaveToContactReq{
		GidOrWxid: gidorwxid,
	}

	resultBody, err := l.http.DoPost("/SaveToContact", req)
	if err != nil {
		return err
	}

	commonResult := SaveToContactResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
