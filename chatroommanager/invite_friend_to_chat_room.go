package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type InviteFriendToChatRoomReq struct {
	Gid  string `json:"gid"`
	Wxid string `json:"wxid"`
}

type InviteFriendToChatRoomResult struct {
	Success string `json:"success"`
}

// InviteFriendToChatRoom 邀请好友进群(40人以上) https://www.showdoc.com.cn/WeChatProject/9021493538921140
func (l *ChatRoomManagerServiceImpl) InviteFriendToChatRoom(gid string, wxid string) error {

	req := InviteFriendToChatRoomReq{
		Gid:  gid,
		Wxid: wxid,
	}

	resultBody, err := l.http.DoPost("/InviteFriendToChatRoom", req)
	if err != nil {
		return err
	}

	commonResult := InviteFriendToChatRoomResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
