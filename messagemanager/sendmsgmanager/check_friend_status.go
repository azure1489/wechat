package sendmsgmanager

import "encoding/json"

// 参数名	必选	类型	说明
// wxidlist	是	string	要检测的好友微信ID清单
// delay	否	string	每个微信请求延迟
type CheckFriendStatusReq struct {
	Wxidlist string `json:"wxidlist"`
	Delay    string `json:"delay,omitempty"`
}

//	{
//	    "code": "0",
//	    "list": [
//	        {
//	            "wxid": "qq180..",
//	            "status": "1"
//	        },
//	        {
//	            "wxid": "wxid_kq..2",
//	            "status": "1"
//	        }
//	    ]
//	}
//
// 参数名	类型	说明
// status	string	1=好友
// 2=被对方删除
// 3=被对方拉黑
// 4=主动拉黑了对方
type CheckFriendStatusResult struct {
	Code string `json:"code"`
	List []struct {
		Wxid   string `json:"wxid"`
		Status string `json:"status"`
	} `json:"list"`
}

// CheckFriendStatus 检测好友状态 https://www.showdoc.com.cn/WeChatProject/9063410601712380
func (l *SendMsgManagerServiceImpl) CheckFriendStatus(req CheckFriendStatusReq) (CheckFriendStatusResult, error) {

	resultBody, err := l.http.DoPost("/CheckFriendStatus", req)
	if err != nil {
		return CheckFriendStatusResult{}, err
	}

	commonResult := CheckFriendStatusResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return CheckFriendStatusResult{}, err
	}

	return commonResult, nil
}
