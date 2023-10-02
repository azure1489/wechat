package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type DisplayChatRoomMemberNickNameReq struct {
	Gid string `json:"gid"`
}

type DisplayChatRoomMemberNickNameResult struct {
	Success string `json:"success"`
}

// DisplayChatRoomMemberNickName 显示群成员昵称 https://www.showdoc.com.cn/WeChatProject/9063349491429947
func (l *ChatRoomManagerServiceImpl) DisplayChatRoomMemberNickName(gid string) error {

	req := DisplayChatRoomMemberNickNameReq{
		Gid: gid,
	}

	resultBody, err := l.http.DoPost("/DisplayChatRoomMemberNickName", req)
	if err != nil {
		return err
	}

	commonResult := DisplayChatRoomMemberNickNameResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
