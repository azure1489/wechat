package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type TurnOffDisplayChatRoomMemberNickNameReq struct {
	Gid string `json:"gid"`
}

type TurnOffDisplayChatRoomMemberNickNameResult struct {
	Success string `json:"success"`
}

// TurnOffDisplayChatRoomMemberNickName 不显示群成员昵称 https://www.showdoc.com.cn/WeChatProject/9063350057038457
func (l *ChatRoomManagerServiceImpl) TurnOffDisplayChatRoomMemberNickName(gid string) error {

	req := TurnOffDisplayChatRoomMemberNickNameReq{
		Gid: gid,
	}

	resultBody, err := l.http.DoPost("/TurnOffDisplayChatRoomMemberNickName", req)
	if err != nil {
		return err
	}

	commonResult := TurnOffDisplayChatRoomMemberNickNameResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
