package sendmsgmanager

import (
	"encoding/json"
	"fmt"
)

type SendTextMsgNoSrcReq struct {
	WxidOrGid string `json:"wxidorgid"`
	Msg       string `json:"msg"`
}

type SendTextMsgNoSrcResult struct {
	SendTextMsgNoSrc string `json:"SendTextMsg_NoSrc"`
}

// SendTextMsgNoSrc 发送文本消息_无源 https://www.showdoc.com.cn/WeChatProject/9859692078921113
func (l *SendMsgManagerServiceImpl) SendTextMsgNoSrc(wxid, msg string) error {

	req := SendTextMsgNoSrcReq{
		WxidOrGid: wxid,
		Msg:       msg,
	}

	resultBody, err := l.http.DoPost("/SendTextMsg_NoSrc", req)
	if err != nil {
		return err
	}

	commonResult := SendTextMsgNoSrcResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendTextMsgNoSrc != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
