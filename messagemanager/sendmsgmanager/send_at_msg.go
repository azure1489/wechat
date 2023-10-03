package sendmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// gid	是	string	群ID
// nicknamelist	否	string	可空,@的对方微信昵称清单，多个自定义昵称用逗号隔开，数量必须和wxidlist一致
// wxidlist	是	string	@的对方微信ID清单，多个wxid用逗号隔开
// msg	否	string	要发送的消息内容，可以不传该参数或留空
type SendAtMsgReq struct {
	Gid          string `json:"gid"`
	Nicknamelist string `json:"nicknamelist,omitempty"`
	Wxidlist     string `json:"wxidlist"`
	Msg          string `json:"msg,omitempty"`
}

//	{
//	    "SendAtAllMsg": "1",
//	  }
type SendAtMsgResult struct {
	SendAtAllMsg string `json:"SendAtAllMsg"`
}

// SendAtMsg 发送群@消息 https://www.showdoc.com.cn/WeChatProject/8929229624237103
func (l *SendMsgManagerServiceImpl) SendAtMsg(req SendAtMsgReq) error {

	resultBody, err := l.http.DoPost("/SendAtMsg", req)
	if err != nil {
		return err
	}

	commonResult := SendAtMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendAtAllMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
