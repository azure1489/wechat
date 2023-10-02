package contactmanager

import (
	"encoding/json"
	"fmt"
)

type RemoveToContactReq struct {
	WxidOrGid string `json:"gidorwxid"`
}

type RemoveToContactResult struct {
	Success string `json:"success"`
}

// RemoveToContact 从通讯录移除  https://www.showdoc.com.cn/WeChatProject/9063365712246544
func (c *ContactManagerServiceImpl) RemoveToContact(wxidOrGid string) error {
	req := &RemoveToContactReq{
		WxidOrGid: wxidOrGid,
	}

	resultBody, err := c.http.DoPost("/RemoveToContact", req)
	if err != nil {
		return err
	}

	commonResult := RemoveToContactResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
