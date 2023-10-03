package contentmsgmanager

import "encoding/json"

// 参数名	必选	类型	说明
// voice_hex	是	string	收到的语音消息内容
// msgsvrid	是	string	收到的语音消息 服务器消息ID
// voice_len	是	string	收到的语音长度
type VoiceToTextReq struct {
	VoiceHex string `json:"voice_hex"`
	Msgsvrid string `json:"msgsvrid"`
	VoiceLen string `json:"voice_len"`
}

// 参数名	必选	类型	说明
// text	是	string	语音转出的文本内容
type VoiceToTextResult struct {
	Text string `json:"text"`
}

// VoiceToText 语音转文本 https://www.showdoc.com.cn/WeChatProject/10074768335116363
func (l *ContentMsgManagerServiceImpl) VoiceToText(req VoiceToTextReq) (string, error) {

	resultBody, err := l.http.DoPost("/Voice2Text", req)
	if err != nil {
		return "", err
	}

	commonResult := VoiceToTextResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.Text, nil
}
