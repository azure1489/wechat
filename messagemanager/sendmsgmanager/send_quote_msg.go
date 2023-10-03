package sendmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// towxid	是	string	对方微信ID/群ID
// title	是	string	要发送的文本消息内容，可以传USC2 例:\uXXXX
// svrid	否	string	源消息 消息ID
// fromusr	否	string	源消息 群ID
// chatusr	是	string	源消息 发送者WXID
// content	是	string	源消息内容
type SendQuoteMsgReq struct {
	Towxid  string `json:"towxid"`
	Title   string `json:"title"`
	Svrid   string `json:"svrid"`
	Fromusr string `json:"fromusr"`
	Chatusr string `json:"chatusr"`
	Content string `json:"content"`
}

//	{
//	    "SendQuoteMsg": "1"
//	}
type SendQuoteMsgResult struct {
	SendQuoteMsg string `json:"SendQuoteMsg"`
}

// SendQuoteMsg 发送引用消息 https://www.showdoc.com.cn/WeChatProject/10435701593451609
func (l *SendMsgManagerServiceImpl) SendQuoteMsg(req SendQuoteMsgReq) error {

	resultBody, err := l.http.DoPost("/SendQuoteMsg", req)
	if err != nil {
		return err
	}
	commonResult := SendQuoteMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}
	if commonResult.SendQuoteMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}
	return nil
}
