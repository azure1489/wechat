package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type UncollapseChatReq struct {
	Gid string `json:"gid"`
}

type UncollapseChatResult struct {
	Success string `json:"success"`
}

// UncollapseChat 展开群聊 https://www.showdoc.com.cn/WeChatProject/9063366996691400
func (l *ChatRoomManagerServiceImpl) UncollapseChat(gid string) error {

	req := UncollapseChatReq{
		Gid: gid,
	}

	resultBody, err := l.http.DoPost("/UncollapseChat", req)
	if err != nil {
		return err
	}

	commonResult := UncollapseChatResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
