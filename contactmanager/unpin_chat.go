package contactmanager

import (
	"encoding/json"
	"fmt"
)

type UnpinChatReq struct {
	WxidOrGid string `json:"wxidorgid"`
}

type UnpinChatResult struct {
	Success string `json:"success"`
}

// UnpinChat 取消置顶聊天  https://www.showdoc.com.cn/WeChatProject/9063351366305816
func (c *ContactManagerServiceImpl) UnpinChat(wxidOrGid string) error {
	req := &UnpinChatReq{
		WxidOrGid: wxidOrGid,
	}

	resultBody, err := c.http.DoPost("/UnpinChat", req)
	if err != nil {
		return err
	}

	commonResult := UnpinChatResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
