package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type CollapseChatReq struct {
	Gid string `json:"gid"`
}

type CollapseChatResult struct {
	Success string `json:"success"`
}

// CollapseChat 收起群聊 https://www.showdoc.com.cn/WeChatProject/9063366035200767
func (l *ChatRoomManagerServiceImpl) CollapseChat(gid string) error {

	req := CollapseChatReq{
		Gid: gid,
	}

	resultBody, err := l.http.DoPost("/CollapseChat", req)
	if err != nil {
		return err
	}

	commonResult := CollapseChatResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
