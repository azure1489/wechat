package sendmsmanager

import (
	"encoding/json"
	"fmt"
)

type SendTextMsgReq struct {
	Wxid string `json:"wxid"`
	Msg  string `json:"msg"`
}

type SendTextMsgResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SendTextMsg 发送文本消息 https://www.showdoc.com.cn/WeChatProject/8929112442643628
func (l *SendMsgManagerServiceImpl) SendTextMsg(wxid, msg string) error {

	req := SendTextMsgReq{
		Wxid: wxid,
		Msg:  msg,
	}

	resultBody, err := l.http.DoPost("/SendTextMsg", req)
	if err != nil {
		return err
	}

	commonResult := []SendTextMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult[0].Message != "success" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
