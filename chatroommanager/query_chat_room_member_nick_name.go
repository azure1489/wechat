package chatroommanager

import "encoding/json"

type QueryChatRoomMemberNickNameReq struct {
	Gid  string `json:"gid"`
	Wxid string `json:"wxid"`
}

type QueryChatRoomMemberNickNameResult struct {
	NickName string `json:"nickname"`
}

// QueryChatRoomMemberNickName 查询群成员昵称 https://www.showdoc.com.cn/WeChatProject/9019565404784998
func (l *ChatRoomManagerServiceImpl) QueryChatRoomMemberNickName(gid string, wxid string) (string, error) {

	req := QueryChatRoomMemberNickNameReq{
		Gid:  gid,
		Wxid: wxid,
	}

	resultBody, err := l.http.DoPost("/QueryChatRoomMemberNickName", req)
	if err != nil {
		return "", err
	}

	commonResult := QueryChatRoomMemberNickNameResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.NickName, nil
}
