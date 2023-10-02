package friendmanager

import (
	"encoding/json"
	"fmt"
)

type BlackFriendReq struct {
	Wxid string `json:"wxid"`
}

type BlackFriendResult struct {
	Success string `json:"success"`
}

// BlackFriend 拉黑好友 https://www.showdoc.com.cn/WeChatProject/9745604340432390
func (l *FriendManagerServiceImpl) BlackFriend(wxid string) error {

	req := BlackFriendReq{
		Wxid: wxid,
	}

	resultBody, err := l.http.DoPost("/BlackFriend", req)
	if err != nil {
		return err
	}

	commonResult := BlackFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
