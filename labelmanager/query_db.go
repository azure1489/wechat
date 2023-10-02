package labelmanager

import (
	"encoding/json"
	"strings"
)

type QueryDBReq struct {
	WxidList string `json:"wxidlist"`
}

type QueryDBResult struct {
	Count string `json:"count"`
	Data  []struct {
		UserName    string `json:"UserName"`
		LabelIDList string `json:"LabelIDList"`
	} `json:"data"`
}

// QueryDB 查询用户所属标签ID清单 https://www.showdoc.com.cn/WeChatProject/9241846291057416
func (l *LabelManagerServiceImpl) QueryDB(wxidList []string) (*QueryDBResult, error) {

	req := QueryDBReq{
		WxidList: strings.Join(wxidList, ","),
	}

	resultBody, err := l.http.DoPost("/QueryDB", req)
	if err != nil {
		return nil, err
	}

	commonResult := QueryDBResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return nil, err
	}

	return &commonResult, nil
}
