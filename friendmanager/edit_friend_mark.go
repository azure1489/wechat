package friendmanager

import (
	"encoding/json"
	"fmt"
)

type EditFriendMarkReq struct {
	Wxid string `json:"wxid"`
	Mark string `json:"mark"`
}

type EditFriendMarkResult struct {
	Success string `json:"success"`
}

// EditFriendMark 修改好友备注 https://www.showdoc.com.cn/WeChatProject/9020497967635095
func (l *FriendManagerServiceImpl) EditFriendMark(wxid string, mark string) error {

	req := EditFriendMarkReq{
		Wxid: wxid,
		Mark: mark,
	}

	resultBody, err := l.http.DoPost("/EditFriendMark", req)
	if err != nil {
		return err
	}

	commonResult := EditFriendMarkResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
