package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type ChatRoomEditNameReq struct {
	Gid  string `json:"gid"`
	Name string `json:"name"`
}

type ChatRoomEditNameResult struct {
	Success string `json:"success"`
}

// ChatRoomEditName 修改群名称 https://www.showdoc.com.cn/WeChatProject/8993745484621751
func (l *ChatRoomManagerServiceImpl) ChatRoomEditName(gid string, name string) error {

	req := ChatRoomEditNameReq{
		Gid:  gid,
		Name: name,
	}

	resultBody, err := l.http.DoPost("/ChatRoomEditName", req)
	if err != nil {
		return err
	}

	commonResult := ChatRoomEditNameResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
