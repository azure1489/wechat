package contactmanager

import (
	"encoding/json"
	"fmt"
)

type MarkAsReadSessionReq struct {
	WxidOrGid string `json:"gidorwxid"`
}

type MarkAsReadSessionResult struct {
	Success string `json:"success"`
}

// MarkAsReadSession 标为已读会话  https://www.showdoc.com.cn/WeChatProject/9219474016553735
func (c *ContactManagerServiceImpl) MarkAsReadSession(wxidOrGid string) error {
	req := &MarkAsReadSessionReq{
		WxidOrGid: wxidOrGid,
	}

	resultBody, err := c.http.DoPost("/MarkAsReadSession", req)
	if err != nil {
		return err
	}

	commonResult := MarkAsReadSessionResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
