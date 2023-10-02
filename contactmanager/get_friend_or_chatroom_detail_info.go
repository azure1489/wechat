package contactmanager

import "encoding/json"

//	{
//	    "type": "friend",
//	    "wxid": "wxid_9r8qlbuy438y21",
//	    "nickname": "大吕同学",
//	    "sex": "1",
//	    "province": "",
//	    "area": "",
//	    "signinfo": "学习新思想 争做新青年",
//	    "wxaccount": "",
//	    "timelinebgurl": "http://szmmsns.qpic.cn/mmsns/Hq0GuDd2VlCibmINoSAOBfsLFxQib3icMIy2ZLdsBSsx3CuxZJUP8ZWNhBlMibMSbRic2zjZ7V4yLaek/0",
//	    "country": "CN",
//	    "headurl": "https://wx.qlogo.cn/mmhead/ver_1/Ziaekj2XhBdrJ8JgdzHoI9FZlB3gASuQAnsfejnd8yekmRs98icsEVBxPeCA4UibbnNUeCHINlqwItiaYM3GfMZqDRBrtuANfAvyPEsnyC7gIxc/0",
//	    "v3": "",
//	    "fromchatroom": "",
//	    "v4": "v4_000b708f0b040000010000000000cabed9f7d61ce9d253393decf4631000000050ded0b020927e3c97896a09d47e6e9ee4763e3a8a2df76e276dbe01519db2bf874f8f075eaa3d1acec47358eff0a3d21a0baee2746a945d087418285aa7a41062b07a10272d3832b0dde88ec0df8d89f4b36a4d329e32688bab61f8f7502d17a44c9915410014412c@stranger"
//	}
type GetFriendOrChatroomDetailInfoResult struct {
	Type          string `json:"type"`
	Wxid          string `json:"wxid"`
	Nickname      string `json:"nickname"`
	Sex           string `json:"sex"`
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

type GetFriendOrChatroomDetailInfoReq struct {
	Wxid string `json:"wxid"`
}

// 获取好友或群聊详情 https://www.showdoc.com.cn/WeChatProject/9157629212860224
func (c *ContactManagerServiceImpl) GetFriendOrChatroomDetailInfo(wxid string) (*GetFriendOrChatroomDetailInfoResult, error) {

	req := GetFriendOrChatroomDetailInfoReq{
		Wxid: wxid,
	}

	resultBody, err := c.http.DoPost("/GetFriendOrChatroomDetailInfo", req)
	if err != nil {
		return nil, err
	}

	commonResult := GetFriendOrChatroomDetailInfoResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
