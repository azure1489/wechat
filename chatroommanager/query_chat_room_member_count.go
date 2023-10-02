package chatroommanager

import (
	"encoding/json"
	"strconv"
)

type QueryChatRoomMemberCountReq struct {
	Gid string `json:"gid"`
}

type QueryChatRoomMemberCountResult struct {
	Count string `json:"count"`
}

// QueryChatRoomMemberCount 本地查询群成员总数(如果查询不到则需要先点一下群聊或者调用网络获取好友或群详细信息) https://www.showdoc.com.cn/WeChatProject/9019566491411017
func (l *ChatRoomManagerServiceImpl) QueryChatRoomMemberCount(gid string) (int, error) {

	req := QueryChatRoomMemberCountReq{
		Gid: gid,
	}

	resultBody, err := l.http.DoPost("/QueryChatRoomMemberCount", req)
	if err != nil {
		return 0, err
	}

	commonResult := QueryChatRoomMemberCountResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return 0, err
	}

	count, err := strconv.Atoi(commonResult.Count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
