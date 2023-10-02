package chatroommanager

import (
	"encoding/json"
	"fmt"
	"strings"
)

type AddFriendToChatRoomReq struct {
	Gid      string `json:"gid"`
	WxidList string `json:"wxidlist"`
}

type AddFriendToChatRoomResult struct {
	Success string `json:"success"`
}

// AddFriendToChatRoom 添加好友进群 https://www.showdoc.com.cn/WeChatProject/9020971162414036
func (l *ChatRoomManagerServiceImpl) AddFriendToChatRoom(gid string, wxidList []string) error {

	req := AddFriendToChatRoomReq{
		Gid:      gid,
		WxidList: strings.Join(wxidList, ","),
	}

	resultBody, err := l.http.DoPost("/AddFriendToChatRoom", req)
	if err != nil {
		return err
	}

	commonResult := AddFriendToChatRoomResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
