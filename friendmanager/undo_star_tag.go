package friendmanager

import (
	"encoding/json"
	"fmt"
)

type UndoStarTagReq struct {
	Wxid string `json:"wxid"`
}

type UndoStarTagResult struct {
	UndoStarTag string `json:"UndoStarTag"`
}

// UndoStarTag 取消星标好友 https://www.showdoc.com.cn/WeChatProject/9078594588816258
func (l *FriendManagerServiceImpl) UndoStarTag(wxid string) error {

	req := UndoStarTagReq{
		Wxid: wxid,
	}

	resultBody, err := l.http.DoPost("/UndoStarTag", req)
	if err != nil {
		return err
	}

	commonResult := UndoStarTagResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.UndoStarTag != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
