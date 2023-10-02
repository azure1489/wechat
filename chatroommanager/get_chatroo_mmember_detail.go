package chatroommanager

import "encoding/json"

type GetChatrooMmemberDetailReq struct {
	Gid string `json:"gid"`
}

//	{
//	    "type": "chatroom",
//	    "gid": "18475851833@chatroom",
//	    "member_count": "4",
//	    "master_wxid": "wxid_abcsm7kmr62s22",
//	    "master_nickname": "Visions",
//	    "chatroom_name": "学习交流群",
//	    "member": [
//	        {
//	            "wxid": "wxid_abcsm7kmr62s22",
//	            "nickname": "Visions",
//	            "user_head_big": "https://wx.qlogo.cn/mmhead/ver_1/ZHDxayhgh4WVgPdj3x4wLJhiaZXDCz...",
//	            "user_head_small": "https://wx.qlogo.cn/mmhead/ver_1/ZHDxayhgh4WVgPdj3x4wLJhiaZXDCzq6DicuG7dYHujMutJ...",
//	            "user_flag": "9",
//	            "inviter_wxid": "themid"
//	        },
//	        {
//	            "wxid": "wxid_71oa2aij4w2n22",
//	            "nickname": "水流",
//	            "user_head_big": "https://wx.qlogo.cn/mmhead/ver_1/gia6wFichwD...",
//	            "user_head_small": "https://wx.qlogo.cn/mmhead/ver_1/gia6wFichwD6r0nj1...",
//	            "user_flag": "1",
//	            "inviter_wxid": "wxid_abcsm7kmr62s22"
//	        },
//	        {
//	            "wxid": "xaixin1314",
//	            "nickname": "郭文海[加油]",
//	            "user_head_big": "https://wx.qlogo.cn/mmhead/ver_1/R422pohCA...",
//	            "user_head_small": "https://wx.qlogo.cn/mmhead/ver_1/R422...",
//	            "user_flag": "1",
//	            "inviter_wxid": "wxid_abcsm7kmr62s22"
//	        },
//	        {
//	            "wxid": "themid",
//	            "nickname": "Hahaa",
//	            "user_head_big": "https://wx.qlogo.cn/mmhead/ver_1/wvrTlK...",
//	            "user_head_small": "https://wx.qlogo.cn/mmhead/ver_1/wvrT...",
//	            "user_flag": "1",
//	            "inviter_wxid": "wxid_abcsm7kmr62s22"
//	        }
//	    ]
//	}
type GetChatrooMmemberDetailResult struct {
	Type           string `json:"type"`
	Gid            string `json:"gid"`
	MemberCount    string `json:"member_count"`
	MasterWxid     string `json:"master_wxid"`
	MasterNickname string `json:"master_nickname"`
	ChatroomName   string `json:"chatroom_name"`
	Member         []struct {
		Wxid          string `json:"wxid"`
		Nickname      string `json:"nickname"`
		UserHeadBig   string `json:"user_head_big"`
		UserHeadSmall string `json:"user_head_small"`
		UserFlag      string `json:"user_flag"`
		InviterWxid   string `json:"inviter_wxid"`
	} `json:"member"`
}

// GetChatrooMmemberDetail 批量获取群成员详细邀请信息,可以获取群成员是由谁邀请进来的 https://www.showdoc.com.cn/WeChatProject/9745666070753468
func (l *ChatRoomManagerServiceImpl) GetChatrooMmemberDetail(gid string) (*GetChatrooMmemberDetailResult, error) {

	req := GetChatrooMmemberDetailReq{
		Gid: gid,
	}

	resultBody, err := l.http.DoPost("/GetChatrooMmemberDetail", req)
	if err != nil {
		return nil, err
	}

	commonResult := GetChatrooMmemberDetailResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
