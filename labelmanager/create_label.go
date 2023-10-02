package labelmanager

import (
	"encoding/json"
	"fmt"
)

type CreateLabelReq struct {
	LabelName string `json:"labelname"`
}

type CreateLabelResult struct {
	Success string `json:"success"`
}

// CreateLabel 创建标签 https://www.showdoc.com.cn/WeChatProject/9078585940900417
func (l *LabelManagerServiceImpl) CreateLabel(labelname string) error {

	req := CreateLabelReq{
		LabelName: labelname,
	}

	resultBody, err := l.http.DoPost("/CreateLabel", req)
	if err != nil {
		return err
	}

	commonResult := CreateLabelResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
