package chatroommanager

import "encoding/json"

type BatchGetChatRoomMemberWxidReq struct {
	Gid string `json:"gid"`
}

//	{
//	    "data": [
//	        {
//	            "wxid": "wxid_123456",
//	        },
//	        {
//	            "wxid": "wxid_123455",
//	        },
//	        {
//	            "wxid": "wxid_123454",
//	        }
//	    ]
//	}
type BatchGetChatRoomMemberWxidResult struct {
	Data []struct {
		Wxid string `json:"wxid"`
	} `json:"data"`
}

// BatchGetChatRoomMemberWxid 批量获取群成员微信ID https://www.showdoc.com.cn/WeChatProject/9021207508542836
func (l *ChatRoomManagerServiceImpl) BatchGetChatRoomMemberWxid(gid string) ([]string, error) {

	req := BatchGetChatRoomMemberWxidReq{
		Gid: gid,
	}

	resultBody, err := l.http.DoPost("/BatchGetChatRoomMemberWxid", req)
	if err != nil {
		return nil, err
	}

	commonResult := BatchGetChatRoomMemberWxidResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	wxids := make([]string, 0)
	for _, item := range commonResult.Data {
		wxids = append(wxids, item.Wxid)
	}

	return wxids, nil
}
