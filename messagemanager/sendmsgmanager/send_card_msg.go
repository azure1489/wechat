package sendmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// towxid	是	string	转发给谁的微信ID/群ID
// fromwxid	是	string	转发谁的微信ID
// nickname	否	string	怎么称呼转发的名片人
type SendCardMsgReq struct {
	Towxid   string `json:"towxid"`
	Fromwxid string `json:"fromwxid"`
	Nickname string `json:"nickname"`
}

//	{
//	    "success": "1",
//	  }
type SendCardMsgResult struct {
	Success string `json:"success"`
}

// SendCardMsg 转发好友名片 https://www.showdoc.com.cn/WeChatProject/8929459103947241
func (l *SendMsgManagerServiceImpl) SendCardMsg(req SendCardMsgReq) error {

	resultBody, err := l.http.DoPost("/SendCardMsg", req)
	if err != nil {
		return err
	}
	commonResult := SendCardMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}
	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}
	return nil
}
