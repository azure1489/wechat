package contactmanager

import (
	"encoding/json"
	"fmt"
)

type TurnOnDoNotDisturbReq struct {
	GidOrWxid string `json:"gidorwxid"`
}

type TurnOnDoNotDisturbResult struct {
	Success string `json:"success"`
}

// TurnOnDoNotDisturb 开启免打扰 https://www.showdoc.com.cn/WeChatProject/9063347648143133
func (c *ContactManagerServiceImpl) TurnOnDoNotDisturb(gidOrWxid string) error {
	req := &TurnOnDoNotDisturbReq{
		GidOrWxid: gidOrWxid,
	}

	resultBody, err := c.http.DoPost("/TurnOnDoNotDisturb", req)
	if err != nil {
		return err
	}

	commonResult := TurnOnDoNotDisturbResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
