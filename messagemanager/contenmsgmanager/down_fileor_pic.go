package contentmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// DownFileorPic	是	string	返回1,下载后的图片在微信缓存目录下
type DownFileorPicResult struct {
	DownFileorPic string `json:"DownFileorPic"`
}

// DownFileorPic 下载图片_使用服务器消息ID https://www.showdoc.com.cn/WeChatProject/9364507487384330
func (l *ContentMsgManagerServiceImpl) DownFileorPic(msgId string) error {

	req := map[string]string{"msg_id": msgId}

	resultBody, err := l.http.DoPost("/DownFileorPic", req)
	if err != nil {
		return err
	}

	commonResult := DownFileorPicResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.DownFileorPic != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
