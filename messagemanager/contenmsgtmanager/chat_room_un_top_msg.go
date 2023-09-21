package contentmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// gid	是	string	群ID
// msgsvrid	是	string	收到的语音消息 服务器消息ID
type ChatRoomUnTopMsgReq struct {
	Gid      string `json:"gid"`
	Msgsvrid string `json:"msgsvrid"`
}

// 参数名	必选	类型	说明
// ChatRoomUnTopMsg	是	string	默认 返回1
type ChatRoomUnTopMsgResult struct {
	ChatRoomUnTopMsg string `json:"ChatRoomUnTopMsg"`
}

// ChatRoomUnTopMsg 取消置顶群聊消息 https://www.showdoc.com.cn/WeChatProject/10273569091488904
func (l *ContentMsgManagerServiceImpl) ChatRoomUnTopMsg(req ChatRoomUnTopMsgReq) error {

	resultBody, err := l.http.DoPost("/ChatRoomUnTopMsg", req)
	if err != nil {
		return err
	}

	commonResult := ChatRoomUnTopMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ChatRoomUnTopMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
