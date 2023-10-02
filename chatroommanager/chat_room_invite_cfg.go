package chatroommanager

import (
	"encoding/json"
)

type ChatRoomInviteCfgReq struct {
	Gid string `json:"gid"`
	Opt string `json:"opt"`
}

type ChatRoomInviteCfgResult struct {
	ChatRoomInviteCfg string `json:"ChatRoomInviteCfg"`
}

// ChatRoomInviteCfg 群聊邀请确认 https://www.showdoc.com.cn/WeChatProject/9859658923614486
func (l *ChatRoomManagerServiceImpl) ChatRoomInviteCfg(gid string, opt string) (string, error) {

	req := ChatRoomInviteCfgReq{
		Gid: gid,
		Opt: opt,
	}

	resultBody, err := l.http.DoPost("/ChatRoomInviteCfg", req)
	if err != nil {
		return "", err
	}

	commonResult := ChatRoomInviteCfgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.ChatRoomInviteCfg, nil
}
