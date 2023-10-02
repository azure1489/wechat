package friendmanager

import (
	"encoding/json"
)

// role 朋友权限	说明
// 0	聊天、朋友圈、微信运动等
// 1	不让他（她）看
// 2	不看他（她）
// 3	不让他（她）看+不看他（她）
// 8	仅聊天

// type 加人类型	说明
// 1	搜索QQ号
// 2	邮箱
// 3	微信号
// 6	单向添加,飞非单项则无法添加
// 10	朋友圈
// 12	QQ好友
// 15	手机号
// 30	扫一扫
// 31	Facebook

// 参数名	必选	类型	说明
// v3_wxid	是	string	v3 或 wxid
// v4	否	string	v4
// desc	是	string	加人要发送的说明
// type	是	string	加人类型,如果采用微信ID加人这里填写30
// role	是	string	朋友能访问的权限
type AddNewFriendReq struct {
	V3WxId string `json:"v3_wxid,omitempty"`
	V4     string `json:"v4,omitempty"`
	Desc   string `json:"desc,omitempty"`
	Type   string `json:"type,omitempty"`
	Role   string `json:"role,omitempty"`
}

// {"Status":"0","RetTxt":"","ret_v3":"v3_020b3826fd0301000000000094db0d44d9963e000000501ea9a3dba12f95f6b60a0536a1adb6c8c23507a72ca7927890723deb898fe58ba3103714cc40f8a585e1e410422e39ca5249051fd958fe7b51e0adbb42bbbbc0ebf7d42186f76ca455f07b73@stranger"}
type AddNewFriendResult struct {
	Status string `json:"Status"`
	RetTxt string `json:"RetTxt"`
	RetVe  string `json:"ret_v3"`
}

// AddNewFriend 添加好友 https://www.showdoc.com.cn/WeChatProject/8961754683625021
func (l *FriendManagerServiceImpl) AddNewFriend(req AddNewFriendReq) (*AddNewFriendResult, error) {

	resultBody, err := l.http.DoPost("/AddNewFriend", req)
	if err != nil {
		return nil, err
	}

	commonResult := AddNewFriendResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
