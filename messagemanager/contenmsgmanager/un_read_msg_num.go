package contentmsgmanager

import "encoding/json"

// 参数名	必选	类型	说明
// GetUnReadMsgNum	是	string	返回未读消息总数
type GetUnReadMsgNumResult struct {
	GetUnReadMsgNum string `json:"GetUnReadMsgNum"`
}

// GetUnReadMsgNum 获取未读消息总数 https://www.showdoc.com.cn/WeChatProject/9690385454945788
func (l *ContentMsgManagerServiceImpl) GetUnReadMsgNum() (string, error) {

	resultBody, err := l.http.DoPost("/GetUnReadMsgNum", nil)
	if err != nil {
		return "", err
	}

	commonResult := GetUnReadMsgNumResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.GetUnReadMsgNum, nil
}
