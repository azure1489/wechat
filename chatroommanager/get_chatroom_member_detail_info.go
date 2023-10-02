package chatroommanager

import "encoding/json"

type GetChatroomMemberDetailInfoReq struct {
	Gid  string `json:"gid"`
	Wxid string `json:"wxid"`
}

//	{
//	    "type": "friend",
//	    "wxid": "wxid_hs...",
//	    "nickname": "三哥",
//	    "province": "Hebei",
//	    "area": "Baoding",
//	    "signinfo": "",
//	    "wxaccount": "Yan_SAN",
//	    "timelinebgurl": "",
//	    "country": "CN",
//	    "headurl": "http://wx.qlogo.cn/mmhead/ver_1/hJSo...",
//	    "v3": "v3_020b3826fd03010000...c@stranger",
//	    "fromchatroom": "17627944@chatroom",
//	    "v4": ""
//	}
type GetChatroomMemberDetailInfoResult struct {
	Type          string `json:"type"`
	Wxid          string `json:"wxid"`
	Nickname      string `json:"nickname"`
	Province      string `json:"province"`
	Area          string `json:"area"`
	Signinfo      string `json:"signinfo"`
	Wxaccount     string `json:"wxaccount"`
	Timelinebgurl string `json:"timelinebgurl"`
	Country       string `json:"country"`
	Headurl       string `json:"headurl"`
	V3            string `json:"v3"`
	Fromchatroom  string `json:"fromchatroom"`
	V4            string `json:"v4"`
}

// GetChatroomMemberDetailInfo 网络获取群成员详细信息,可以获取群成员微信号 https://www.showdoc.com.cn/WeChatProject/9158592965427244
func (l *ChatRoomManagerServiceImpl) GetChatroomMemberDetailInfo(gid string, wxid string) (*GetChatroomMemberDetailInfoResult, error) {

	req := GetChatroomMemberDetailInfoReq{
		Gid:  gid,
		Wxid: wxid,
	}

	resultBody, err := l.http.DoPost("/GetChatroomMemberDetailInfo", req)
	if err != nil {
		return nil, err
	}

	commonResult := GetChatroomMemberDetailInfoResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
