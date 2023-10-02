package contactmanager

import "encoding/json"

type BeforeTransferReq struct {
	Wxid string `json:"wxid"`
}

type BeforeTransferResult struct {
	Name string `json:"name"`
}

// BeforeTransfer 查询好友实名 https://www.showdoc.com.cn/WeChatProject/10171243195617836
func (c *ContactManagerServiceImpl) BeforeTransfer(wxid string) (string, error) {
	req := &BeforeTransferReq{
		Wxid: wxid,
	}

	resultBody, err := c.http.DoPost("/BeforeTransfer", req)
	if err != nil {
		return "", err
	}

	commonResult := BeforeTransferResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.Name, nil
}
