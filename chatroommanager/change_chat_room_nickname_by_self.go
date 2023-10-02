package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type ChangeChatRoomNicknameBySelfReq struct {
	Gid      string `json:"gid"`
	Wxid     string `json:"wxid"`
	Nickname string `json:"nickname"`
}

type ChangeChatRoomNicknameBySelfResult struct {
	Success string `json:"success"`
}

// ChangeChatRoomNicknameBySelf 修改自己在群里的昵称 https://www.showdoc.com.cn/WeChatProject/9019580625786730
func (l *ChatRoomManagerServiceImpl) ChangeChatRoomNicknameBySelf(gid string, wxid string, nickname string) error {

	req := ChangeChatRoomNicknameBySelfReq{
		Gid:      gid,
		Wxid:     wxid,
		Nickname: nickname,
	}

	resultBody, err := l.http.DoPost("/ChangeChatRoomNicknameBySelf", req)
	if err != nil {
		return err
	}

	commonResult := ChangeChatRoomNicknameBySelfResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
