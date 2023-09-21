package sendmsmanager

import (
	"encoding/json"
	"fmt"
)

type SendPatMsgReq struct {
	Wxid string `json:"wxid"`
	Gid  string `json:"gid"`
}

type SendPatMsgResult struct {
	SendPatMsg string `json:"SendPatMsg"`
}

// SendPatMsg 发送拍一拍消息 https://www.showdoc.com.cn/WeChatProject/9644354859714189
func (l *SendMsgManagerServiceImpl) SendPatMsg(wxid, gid string) error {

	req := SendPatMsgReq{
		Wxid: wxid,
		Gid:  gid,
	}
	resultBody, err := l.http.DoPost("/SendLocationMsg", req)
	if err != nil {
		return err
	}
	commonResult := SendPatMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}
	if commonResult.SendPatMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}
	return nil
}
