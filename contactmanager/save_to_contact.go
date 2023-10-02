package contactmanager

import (
	"encoding/json"
	"fmt"
)

type SaveToContactReq struct {
	WxidOrGid string `json:"gidorwxid"`
}

type SaveToContactResult struct {
	Success string `json:"success"`
}

// SaveToContact 保存到通讯录  https://www.showdoc.com.cn/WeChatProject/9063365453994960
func (c *ContactManagerServiceImpl) SaveToContact(wxidOrGid string) error {
	req := &SaveToContactReq{
		WxidOrGid: wxidOrGid,
	}

	resultBody, err := c.http.DoPost("/SaveToContact", req)
	if err != nil {
		return err
	}

	commonResult := SaveToContactResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
