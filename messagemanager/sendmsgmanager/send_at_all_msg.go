package sendmsmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// gid	是	string	群ID
// msg	否	string	要发送的消息内容
type SendAtAllMsgReq struct {
	Gid string `json:"gid"`
	Msg string `json:"msg,omitempty"`
}

//	{
//	    "success": "1",
//	  }
type SendAtAllMsgResult struct {
	Success string `json:"success"`
}

// SendAtAllMsg 发送群@全体成员消息 https://www.showdoc.com.cn/WeChatProject/8981408454846254
func (l *SendMsgManagerServiceImpl) SendAtAllMsg(req SendAtAllMsgReq) error {

	resultBody, err := l.http.DoPost("/SendAtAllMsg", req)
	if err != nil {
		return err
	}

	commonResult := SendAtAllMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
