package chatroommanager

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CreateChatRoomReq struct {
	WxidList string `json:"wxidlist"`
}

// 参数名	类型	说明
// retstr	string	返回操作状态文本
// gid	string	返回群ID
type CreateChatRoomResult struct {
	RetStr string `json:"retstr"` // 返回操作状态文本 Everything is OK
	Gid    string `json:"gid"`    // 返回群ID
}

// CreateChatRoom 创建群聊 https://www.showdoc.com.cn/WeChatProject/9019609319630104
func (l *ChatRoomManagerServiceImpl) CreateChatRoom(wxidList []string) (string, error) {

	req := CreateChatRoomReq{
		WxidList: strings.Join(wxidList, ","),
	}

	resultBody, err := l.http.DoPost("/CreateChatRoom", req)
	if err != nil {
		return "", err
	}

	commonResult := CreateChatRoomResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	if commonResult.RetStr != "Everything is OK" {
		return "", fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return commonResult.Gid, nil
}
