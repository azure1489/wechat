package chatroommanager

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ChatRoomVoipReq struct {
	Gid      string `json:"gid"`
	WxidList string `json:"wxidlist"`
}

type ChatRoomVoipResult struct {
	Success string `json:"success"`
}

// ChatRoomVoip 群聊多人语音 https://www.showdoc.com.cn/WeChatProject/9019559882256250
func (l *ChatRoomManagerServiceImpl) ChatRoomVoip(gid string, wxidList []string) error {

	req := ChatRoomVoipReq{
		Gid:      gid,
		WxidList: strings.Join(wxidList, ","),
	}

	resultBody, err := l.http.DoPost("/ChatRoomVoip", req)
	if err != nil {
		return err
	}

	commonResult := ChatRoomVoipResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
