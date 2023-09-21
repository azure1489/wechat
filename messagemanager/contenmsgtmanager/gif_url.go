package contentmsgmanager

import "encoding/json"

// 参数名	必选	类型	说明
// msg_xml	是	string	收到的gif消息完整xml内容
type GetGIFURLReq struct {
	MsgXml string `json:"msg_xml"`
}

// 参数名	必选	类型	说明
// GetGIFURL	是	string	返回GIF的访问URL地址
type GetGIFURLResult struct {
	GetGIFURL string `json:"GetGIFURL"`
}

// GetGIFURL 获取GIF的访问URL https://www.showdoc.com.cn/WeChatProject/9690133784680867
func (l *ContentMsgManagerServiceImpl) GetGIFURL(msgXml string) (string, error) {

	req := GetGIFURLReq{
		MsgXml: msgXml,
	}

	resultBody, err := l.http.DoPost("/GetGIFURL", req)
	if err != nil {
		return "", err
	}

	commonResult := GetGIFURLResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.GetGIFURL, nil
}
