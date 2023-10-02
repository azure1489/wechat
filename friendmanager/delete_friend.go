package friendmanager

import (
	"encoding/json"
	"fmt"
)

type DeleteFriendReq struct {
	Wxid string `json:"wxid"`
}

type DeleteFriendResult struct {
	Success string `json:"success"`
}

// DeleteFriend 删除好友 https://www.showdoc.com.cn/WeChatProject/8981501348113066
func (l *FriendManagerServiceImpl) DeleteFriend(wxid string) error {

	req := DeleteFriendReq{
		Wxid: wxid,
	}

	resultBody, err := l.http.DoPost("/DeleteFriend", req)
	if err != nil {
		return err
	}

	commonResult := DeleteFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
