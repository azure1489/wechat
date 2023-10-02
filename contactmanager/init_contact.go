package contactmanager

import "encoding/json"

// InitContactResult 初始化通讯录返回结果
type InitContactResult struct {
	InitContact string `json:"InitContact"` // 返回状态码
	Batch       []struct {
		Number string `json:"number"` // 批次号，每个批次号最多显示100个id
		List   string `json:"list"`   // 通讯录联系人的ID清单
	} `json:"batch"` // 批次号，每个批次号最多显示100个id
}

// 初始化通讯录 https://www.showdoc.com.cn/WeChatProject/9730605278628502
func (c *ContactManagerServiceImpl) InitContact() (*InitContactResult, error) {
	resultBody, err := c.http.DoPost("/InitContact", nil)
	if err != nil {
		return nil, err
	}

	commonResult := InitContactResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
