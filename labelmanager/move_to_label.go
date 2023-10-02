package labelmanager

import (
	"encoding/json"
	"fmt"
	"strings"
)

//	{
//	    "wxidlist": "themid,wxid_123456",
//	    "labelid": "1,2,3,4,5,6"
//	}
type MoveToLabelReq struct {
	WxIdList string `json:"wxidlist"`
	LabelId  string `json:"labelid"`
}

type MoveToLabelResult struct {
	Success string `json:"success"`
}

// MoveToLabel 移动联系人到标签 https://www.showdoc.com.cn/WeChatProject/9078583356534047
func (l *LabelManagerServiceImpl) MoveToLabel(wxidList []string, labelid string) error {

	req := MoveToLabelReq{
		WxIdList: strings.Join(wxidList, ","),
		LabelId:  labelid,
	}

	resultBody, err := l.http.DoPost("/MoveToLabel", req)
	if err != nil {
		return err
	}

	commonResult := MoveToLabelResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
