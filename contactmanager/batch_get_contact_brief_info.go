package contactmanager

import (
	"encoding/json"
	"strings"
)

// 批量获取联系人简明信息请求参数
type BatchGetContactBriefInfoReq struct {
	WxidList string `json:"wxidlist"`
}

// 批量获取联系人简明信息返回结果
type BatchGetContactBriefInfoReult struct {
	Count string `json:"count"`
	Info  []struct {
		Wxid      string `json:"wxid"`
		Nickname  string `json:"nickname"`
		Markname  string `json:"markname"`
		Country   string `json:"country"`
		City      string `json:"city"`
		Bighead   string `json:"bighead"`
		Smallhead string `json:"smallhead"`
	} `json:"info"`
}

// 批量获取联系人简明信息
func (c *ContactManagerServiceImpl) BatchGetContactBriefInfo(wxids []string) (*BatchGetContactBriefInfoReult, error) {

	req := BatchGetContactBriefInfoReq{
		WxidList: strings.Join(wxids, ","),
	}

	resultBody, err := c.http.DoPost("/BatchGetContactBriefInfo", req)
	if err != nil {
		return nil, err
	}

	commonResult := BatchGetContactBriefInfoReult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
