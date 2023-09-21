package collectionmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// fromwxid	是	string	对方微信ID
// transferid	是	string	转账消息ID
type UnCollectionReq struct {
	Fromwxid   string `json:"fromwxid"`
	Transferid string `json:"transferid"`
}

// 参数名	类型	说明
// success	string	成功返回1 失败返回 0
type UnCollectionResult struct {
	Success string `json:"success"`
}

// UnCollection 退还转账 https://www.showdoc.com.cn/WeChatProject/9652344490820665
func (l *CollectionManagerServiceImpl) UnCollection(req UnCollectionReq) error {

	resultBody, err := l.http.DoPost("/UnCollection", req)
	if err != nil {
		return err
	}
	commonResult := UnCollectionResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}
	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}
	return nil
}
