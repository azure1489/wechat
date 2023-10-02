package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type QuitChatRoomReq struct {
	Gid string `json:"gid"`
}

type QuitChatRoomResult struct {
	Success string `json:"success"`
}

// QuitChatRoom 退出群聊 hhttps://www.showdoc.com.cn/WeChatProject/9002651226705549
func (l *ChatRoomManagerServiceImpl) QuitChatRoom(gid string) error {

	req := QuitChatRoomReq{
		Gid: gid,
	}

	resultBody, err := l.http.DoPost("/QuitChatRoom", req)
	if err != nil {
		return err
	}

	commonResult := QuitChatRoomResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
