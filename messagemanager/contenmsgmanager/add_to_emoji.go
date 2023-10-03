package contentmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// filemd5	是	string	收到的gif表情md5
type AddToEmojiReq struct {
	Filemd5 string `json:"filemd5"`
}

//	参数名	必选	类型	说明
//
// AddtoEmoji	是	string	1
type AddToEmojiResult struct {
	AddtoEmoji string `json:"AddtoEmoji"`
}

// AddToEmoji 方法_添加到自己表情 https://www.showdoc.com.cn/WeChatProject/9966833153857757
func (l *ContentMsgManagerServiceImpl) AddToEmoji(filemd5 string) error {
	resultBody, err := l.http.DoPost("/AddtoEmoji", nil)
	if err != nil {
		return err
	}

	commonResult := AddToEmojiResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.AddtoEmoji != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
