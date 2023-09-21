package sendmsmanager

import "encoding/json"

// 参数名	必选	类型	说明
// toid	是	string	对方微信ID/群ID
// aeskey	是	string	来源于接收到的语音消息内容中
// voiceurl	是	string	来源于接收到的语音消息内容中
// voicelength	是	string	语音内容长度
// length	是	string	语音显示长度。最长60秒
type ForwardVoiceMsgReq struct {
	Toid        string `json:"toid"`
	Aeskey      string `json:"aeskey"`
	Voiceurl    string `json:"voiceurl"`
	Voicelength string `json:"voicelength"`
	Length      string `json:"length"`
}

//	{
//	    "msgid": "2636943258839319581"
//	}
type ForwardVoiceMsgResult struct {
	Msgid string `json:"msgid"`
}

// ForwardVoiceMsg 转发语音消息 https://www.showdoc.com.cn/WeChatProject/10042262554089499
func (l *SendMsgManagerServiceImpl) ForwardVoiceMsg(req ForwardVoiceMsgReq) (string, error) {

	resultBody, err := l.http.DoPost("/ForwardVoiceMsg", req)
	if err != nil {
		return "", err
	}

	commonResult := ForwardVoiceMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.Msgid, nil
}
