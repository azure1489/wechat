package sendmsmanager

import (
	"encoding/json"
	"fmt"
)

type SendLocationMsgReq struct {
	Wxid string `json:"wxid"`
	Msg  string `json:"msg"`
}

type SendLocationMsgResult struct {
	SendLocationMsg string `json:"SendLocationMsg"`
}

// SendLocationMsg 发送位置消息 https://www.showdoc.com.cn/WeChatProject/9167445492364225
func (l *SendMsgManagerServiceImpl) SendLocationMsg(wxid, msg string) error {

	req := SendLocationMsgReq{
		Wxid: wxid,
		Msg:  msg,
	}

	resultBody, err := l.http.DoPost("/SendLocationMsg", req)
	if err != nil {
		return err
	}

	commonResult := SendLocationMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendLocationMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
