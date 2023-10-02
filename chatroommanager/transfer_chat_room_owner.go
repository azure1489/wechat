package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type TransferChatRoomOwnerReq struct {
	Gid  string `json:"gid"`
	Wxid string `json:"wxid"`
}

type TransferChatRoomOwnerResult struct {
	Stuts  int    `json:"status"`
	Rettxt string `json:"rettxt"`
}

// TransferChatRoomOwner 转让群主 https://www.showdoc.com.cn/WeChatProject/10518532504405659
func (l *ChatRoomManagerServiceImpl) TransferChatRoomOwner(gid string, wxid string) error {

	req := TransferChatRoomOwnerReq{
		Gid:  gid,
		Wxid: wxid,
	}

	resultBody, err := l.http.DoPost("/TransferChatRoomOwner", req)
	if err != nil {
		return err
	}

	commonResult := TransferChatRoomOwnerResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Stuts != 0 {
		return fmt.Errorf("提交失败, body=%s", commonResult.Rettxt)
	}

	return nil
}
