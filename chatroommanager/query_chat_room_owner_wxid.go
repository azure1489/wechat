package chatroommanager

import "encoding/json"

type QueryChatRoomOwnerWxidReq struct {
	Gid string `json:"gid"`
}

type QueryChatRoomOwnerWxidResult struct {
	OwnerWxid string `json:"OwnerWxid"`
}

// QueryChatRoomOwnerWxid 查询群主微信ID https://www.showdoc.com.cn/WeChatProject/9019577323953766
func (l *ChatRoomManagerServiceImpl) QueryChatRoomOwnerWxid(gid string) (string, error) {

	req := QueryChatRoomOwnerWxidReq{
		Gid: gid,
	}

	resultBody, err := l.http.DoPost("/QueryChatRoomOwnerWxid", req)
	if err != nil {
		return "", err
	}

	commonResult := QueryChatRoomOwnerWxidResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.OwnerWxid, nil
}
