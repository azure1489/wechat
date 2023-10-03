package contentmsgmanager

import "encoding/json"

// 参数名	必选	类型	说明
// msgtype	是	string	消息类型
// towxid	是	string	消息接收人
// content	是	string	消息内容
// msglocID	是	string	消息本地ID
// Lib_ID	是	string	消息分库ID
type MsgStructReust struct {
	Msgtype  string `json:"msgtype"`
	Towxid   string `json:"towxid"`
	Content  string `json:"content"`
	MsglocID string `json:"msglocID"`
	LibID    string `json:"Lib_ID"`
}

// MsgStruct 获取消息内容 https://www.showdoc.com.cn/WeChatProject/9364487555891575
func (l *ContentMsgManagerServiceImpl) MsgStruct(msgId string) (MsgStructReust, error) {

	resultBody, err := l.http.DoPost("/MsgStruct", map[string]string{"msgid": msgId})
	if err != nil {
		return MsgStructReust{}, err
	}

	commonResult := MsgStructReust{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return MsgStructReust{}, err
	}

	return commonResult, nil
}
