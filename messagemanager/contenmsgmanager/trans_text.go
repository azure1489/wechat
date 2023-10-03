package contentmsgmanager

import "encoding/json"

// 参数名	必选	类型	说明
// text	是	string	需要翻译的内容
type TranslateReq struct {
	Text string `json:"text"`
}

// 参数名	必选	类型	说明
// GetTransText	是	string	翻译后的结果
type TranslateResult struct {
	GetTransText string `json:"GetTransText"`
}

// Translate 文本翻译 https://www.showdoc.com.cn/WeChatProject/9690133784680867
func (l *ContentMsgManagerServiceImpl) Translate(text string) (string, error) {

	req := TranslateReq{
		Text: text,
	}

	resultBody, err := l.http.DoPost("/GetTransText", req)
	if err != nil {
		return "", err
	}

	commonResult := TranslateResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.GetTransText, nil
}
