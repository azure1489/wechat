package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type ChatRoomEditMarkReq struct {
	Gid  string `json:"gid"`
	Mark string `json:"mark"`
}

type ChatRoomEditMarkResult struct {
	Success string `json:"success"`
}

// ChatRoomEditMark 修改群备注 https://www.showdoc.com.cn/WeChatProject/8994766712410900
func (l *ChatRoomManagerServiceImpl) ChatRoomEditMark(gid string, mark string) error {

	req := ChatRoomEditMarkReq{
		Gid:  gid,
		Mark: mark,
	}

	resultBody, err := l.http.DoPost("/ChatRoomEditMark", req)
	if err != nil {
		return err
	}

	commonResult := ChatRoomEditMarkResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
