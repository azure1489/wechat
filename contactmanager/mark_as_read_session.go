package contactmanager

import (
	"encoding/json"
	"fmt"
)

type MarkAsNoReadSessionReq struct {
	WxidOrGid string `json:"gidorwxid"`
}

type MarkAsNoReadSessionResult struct {
	Success string `json:"success"`
}

// MarkAsNoReadSession 标为未读会话  https://www.showdoc.com.cn/WeChatProject/9644356701774840
func (c *ContactManagerServiceImpl) MarkAsNoReadSession(wxidOrGid string) error {
	req := &MarkAsNoReadSessionReq{
		WxidOrGid: wxidOrGid,
	}

	resultBody, err := c.http.DoPost("/MarkAsNoReadSession", req)
	if err != nil {
		return err
	}

	commonResult := MarkAsNoReadSessionResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
