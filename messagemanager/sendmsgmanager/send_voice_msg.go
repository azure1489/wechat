package sendmsgmanager

import (
	"encoding/json"
	"fmt"
)

type SendVoiceMsgReq struct {
	Wxid      string `json:"wxid"`
	VoicePath string `json:"voicepath"`
}

type SendVoiceMsgResult struct {
	SendVoiceMsg string `json:"SendVoiceMsg"`
}

// SendVoiceMsg 发送语音消息 https://www.showdoc.com.cn/WeChatProject/892913222925935
func (l *SendMsgManagerServiceImpl) SendVoiceMsg(wxid, voicePath string) error {

	req := SendVoiceMsgReq{
		Wxid:      wxid,
		VoicePath: voicePath,
	}

	resultBody, err := l.http.DoPost("/SendVoiceMsg", req)
	if err != nil {
		return err
	}

	commonResult := SendVoiceMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendVoiceMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
