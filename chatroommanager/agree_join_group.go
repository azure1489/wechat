package chatroommanager

import (
	"encoding/json"
	"fmt"
)

type AgreeJoinGroupReq struct {
	GroupAccessUrl string `json:"GroupAccessUrl"`
}

type AgreeJoinGroupResult struct {
	Success string `json:"success"`
}

// AgreeJoinGroup 同意入群 https://www.showdoc.com.cn/WeChatProject/9032882631445023
func (l *ChatRoomManagerServiceImpl) AgreeJoinGroup(groupAccessUrl string) error {

	req := AgreeJoinGroupReq{
		GroupAccessUrl: groupAccessUrl,
	}

	resultBody, err := l.http.DoPost("/AgreeJoinGroup", req)
	if err != nil {
		return err
	}

	commonResult := AgreeJoinGroupResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
