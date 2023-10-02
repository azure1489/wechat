package friendmanager

import (
	"encoding/json"
	"fmt"
)

type VideoCallFriendReq struct {
	Wxid string `json:"wxid"`
}

type VideoCallFriendResult struct {
	Success string `json:"success"`
}

// VideoCallFriend 视频通话好友 https://www.showdoc.com.cn/WeChatProject/9019318364312888
func (l *FriendManagerServiceImpl) VideoCallFriend(wxid string) error {

	req := VideoCallFriendReq{
		Wxid: wxid,
	}

	resultBody, err := l.http.DoPost("/VideoCallFriend", req)
	if err != nil {
		return err
	}

	commonResult := VideoCallFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
