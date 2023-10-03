package contentmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// gid	是	string	群ID
// msgsvrid	是	string	收到的语音消息 服务器消息ID
// content	是	string	要显示的置顶文本内容，可以自定义
type ChatRoomTopMsgReq struct {
	Gid      string `json:"gid"`
	Msgsvrid string `json:"msgsvrid"`
	Content  string `json:"content"`
}

// 参数名	必选	类型	说明
// ChatRoomTopMsg	是	string	默认 返回1
type ChatRoomTopMsgResult struct {
	ChatRoomTopMsg string `json:"ChatRoomTopMsg"`
}

// ChatRoomTopMsg 群置顶消息 https://www.showdoc.com.cn/WeChatProject/10273567687375653
func (l *ContentMsgManagerServiceImpl) ChatRoomTopMsg(req ChatRoomTopMsgReq) error {

	resultBody, err := l.http.DoPost("/ChatRoomTopMsg", req)
	if err != nil {
		return err
	}

	commonResult := ChatRoomTopMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ChatRoomTopMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
