package contactmanager

import (
	"encoding/json"
	"fmt"
)

type StickyChatReq struct {
	WxidOrGid string `json:"wxidorgid"`
}

type StickyChatResult struct {
	Success string `json:"success"`
}

// StickyChat 置顶聊天  https://www.showdoc.com.cn/WeChatProject/9019589526757527
func (c *ContactManagerServiceImpl) StickyChat(wxidOrGid string) error {
	req := &StickyChatReq{
		WxidOrGid: wxidOrGid,
	}

	resultBody, err := c.http.DoPost("/StickyChat", req)
	if err != nil {
		return err
	}

	commonResult := StickyChatResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
