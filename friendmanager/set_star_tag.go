package friendmanager

import (
	"encoding/json"
	"fmt"
)

type SetStarTagReq struct {
	Wxid string `json:"wxid"`
}

type SetStarTagResult struct {
	SetStarTag string `json:"SetStarTag"`
}

// SetStarTag 设置星标好友 https://www.showdoc.com.cn/WeChatProject/9078593012203121
func (l *FriendManagerServiceImpl) SetStarTag(wxid string) error {

	req := SetStarTagReq{
		Wxid: wxid,
	}

	resultBody, err := l.http.DoPost("/SetStarTag", req)
	if err != nil {
		return err
	}

	commonResult := SetStarTagResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SetStarTag != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
