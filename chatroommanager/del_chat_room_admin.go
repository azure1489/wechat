package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type DelChatRoomAdminReq struct {
	Gid  string `json:"gid"`
	Wxid string `json:"wxid"`
}

//	{
//	    "status": 0,
//	    "rettxt": ""
//	}
type DelChatRoomAdminResult struct {
	Status int    `json:"status"`
	Rettxt string `json:"rettxt"`
}

// DelChatRoomAdmin 删除群管理员 https://www.showdoc.com.cn/WeChatProject/10518531857265007
func (l *ChatRoomManagerServiceImpl) DelChatRoomAdmin(gid string, wxid string) error {

	req := DelChatRoomAdminReq{
		Gid:  gid,
		Wxid: wxid,
	}

	resultBody, err := l.http.DoPost("/DelChatRoomAdmin", req)
	if err != nil {
		return err
	}

	commonResult := DelChatRoomAdminResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Status != 0 {
		return fmt.Errorf("提交失败, body=%s", commonResult.Rettxt)
	}

	return nil
}
