package sendmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// msgsvrid	是	string	要转发的服务器消息ID，服务器消息ID可以从接收消息回调里拿到
// wxid	是	string	对方微信ID/群ID
type FowardAllMsgReq struct {
	Msgsvrid string `json:"msgsvrid"`
	Wxid     string `json:"wxid"`
}

//	{
//	    "success": "1"
//	  }
type FowardAllMsgResult struct {
	Success string `json:"success"`
}

// FowardAllMsg 转发任意消息（注意：无法转发语音） https://www.showdoc.com.cn/WeChatProject/9090147365509163
func (l *SendMsgManagerServiceImpl) FowardAllMsg(req FowardAllMsgReq) error {

	resultBody, err := l.http.DoPost("/FowardAllMsg", req)
	if err != nil {
		return err
	}

	commonResult := FowardAllMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
