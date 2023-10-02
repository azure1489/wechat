package contactmanager

import (
	"encoding/json"
	"fmt"
)

type TurnOffDoNotDisturbReq struct {
	GidOrWxid string `json:"gidorwxid"`
}

type TurnOffDoNotDisturbResult struct {
	Success string `json:"success"`
}

// TurnOffDoNotDisturb 关闭消息免打扰 https://www.showdoc.com.cn/WeChatProject/9063348367251383
func (c *ContactManagerServiceImpl) TurnOffDoNotDisturb(gidOrWxid string) error {
	req := &TurnOffDoNotDisturbReq{
		GidOrWxid: gidOrWxid,
	}

	resultBody, err := c.http.DoPost("/TurnOffDoNotDisturb", req)
	if err != nil {
		return err
	}

	commonResult := TurnOffDoNotDisturbResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
