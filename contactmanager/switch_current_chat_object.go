package contactmanager

import (
	"encoding/json"
	"fmt"
)

type SwitchCurrentChatObjectReq struct {
	WxidOrGid string `json:"wxidorgid"`
}

type SwitchCurrentChatObjectResult struct {
	Success string `json:"success"`
}

// SwitchCurrentChatObject 切换当前聊天对象  https://www.showdoc.com.cn/WeChatProject/9019589526757527
func (c *ContactManagerServiceImpl) SwitchCurrentChatObject(wxidOrGid string) error {
	req := &SwitchCurrentChatObjectReq{
		WxidOrGid: wxidOrGid,
	}

	resultBody, err := c.http.DoPost("/SwitchCurrentChatObject", req)
	if err != nil {
		return err
	}

	commonResult := SwitchCurrentChatObjectResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
