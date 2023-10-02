package friendmanager

import (
	"encoding/json"
	"fmt"
)

type VoiceCallFriendReq struct {
	Wxid string `json:"wxid"`
}

type VoiceCallFriendResult struct {
	Success string `json:"success"`
}

// VoiceCallFriend 语音通话好友 https://www.showdoc.com.cn/WeChatProject/9019318364312888
func (l *FriendManagerServiceImpl) VoiceCallFriend(wxid string) error {

	req := VoiceCallFriendReq{
		Wxid: wxid,
	}

	resultBody, err := l.http.DoPost("/VoiceCallFriend", req)
	if err != nil {
		return err
	}

	commonResult := VoiceCallFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
