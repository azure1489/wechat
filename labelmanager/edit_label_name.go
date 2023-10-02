package labelmanager

import (
	"encoding/json"
	"fmt"
)

//	{
//	    "labelname":"新标签名字",
//	    "labelid": "1"
//	}
type EditLabelNameReq struct {
	LabelName string `json:"labelname"`
	LabelId   string `json:"labelid"`
}

type EditLabelNameResult struct {
	Success string `json:"success"`
}

// EditLabelName 编辑标签名字 https://www.showdoc.com.cn/WeChatProject/9078586377107582
func (l *LabelManagerServiceImpl) EditLabelName(labelName string, labelId string) error {

	req := EditLabelNameReq{
		LabelName: labelName,
		LabelId:   labelId,
	}

	resultBody, err := l.http.DoPost("/EditLabelName", req)
	if err != nil {
		return err
	}

	commonResult := EditLabelNameResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
