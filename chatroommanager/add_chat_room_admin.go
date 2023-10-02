package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type AddChatRoomAdminReq struct {
	Gid  string `json:"gid"`
	Wxid string `json:"wxid"`
}

//	{
//	    "status": 0,
//	    "rettxt": ""
//	}
type AddChatRoomAdminResult struct {
	Stuts  int    `json:"status"`
	Rettxt string `json:"rettxt"`
}

// AddChatRoomAdmin 添加群管理员 https://www.showdoc.com.cn/WeChatProject/10518529776146702
func (l *ChatRoomManagerServiceImpl) AddChatRoomAdmin(gid string, wxid string) error {

	req := AddChatRoomAdminReq{
		Gid:  gid,
		Wxid: wxid,
	}

	resultBody, err := l.http.DoPost("/AddChatRoomAdmin", req)
	if err != nil {
		return err
	}

	commonResult := AddChatRoomAdminResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Stuts != 0 {
		return fmt.Errorf("提交失败, body=%s", commonResult.Rettxt)
	}

	return nil
}
