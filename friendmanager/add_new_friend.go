package friendmanager

import (
	"encoding/json"
	"fmt"
)

//	{
//	    "v3_wxid": "v3_xxx",
//	    "v4": "v4_xxx",
//	    "desc": "我的你的好朋友~",
//	    "type": "30",
//	    "role": "0"
//	  }
type AddNewFriendReq struct {
	V3WxId string `json:"v3_wxid,omitempty"`
	V4     string `json:"v4,omitempty"`
	Desc   string `json:"desc,omitempty"`
	Type   string `json:"type,omitempty"`
	Role   string `json:"role,omitempty"`
}

type AddNewFriendResult struct {
	Success string `json:"success"`
}

// AddNewFriend 添加好友 https://www.showdoc.com.cn/WeChatProject/8961754683625021
func (l *FriendManagerServiceImpl) AddNewFriend(req *AddNewFriendReq) error {

	resultBody, err := l.http.DoPost("/AddNewFriend", req)
	if err != nil {
		return err
	}

	commonResult := AddNewFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
