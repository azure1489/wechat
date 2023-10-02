package chatroommanager

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ChatRoomMemberBatchDeleteReq struct {
	Gid      string `json:"gid"`
	WxidList string `json:"wxidlist"`
}

type ChatRoomMemberBatchDeleteResult struct {
	Success bool `json:"success"`
}

// ChatRoomMemberBatchDelete 批量删除群成员 https://www.showdoc.com.cn/WeChatProject/8993612184500147
func (l *ChatRoomManagerServiceImpl) ChatRoomMemberBatchDelete(gid string, wxids []string) error {

	req := ChatRoomMemberBatchDeleteReq{
		Gid:      gid,
		WxidList: strings.Join(wxids, ","),
	}

	resultBody, err := l.http.DoPost("/ChatRoomMemberBatchDelete", req)
	if err != nil {
		return err
	}

	commonResult := ChatRoomMemberBatchDeleteResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if !commonResult.Success {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
