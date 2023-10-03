package sendmsgmanager

import (
	"encoding/json"
)

type SendVoiceMsgReq struct {
	Wxid      string `json:"wxid"`
	VideoFile string `json:"voice_file"`
	TimeMs    int    `json:"time_ms"`
}

type SendVoiceMsgResult struct {
	MsgSvrID string `json:"MsgSvrID"`
}

// SendVoiceMsg 发送语音消息 https://www.showdoc.com.cn/WeChatProject/892913222925935
func (l *SendMsgManagerServiceImpl) SendVoiceMsg(wxid, videoFile string, timeMs int) (string, error) {

	req := SendVoiceMsgReq{
		Wxid:      wxid,
		VideoFile: videoFile,
		TimeMs:    timeMs,
	}

	resultBody, err := l.http.DoPost("/SendVoiceMsg", req)
	if err != nil {
		return "", err
	}

	commonResult := SendVoiceMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.MsgSvrID, nil
}
