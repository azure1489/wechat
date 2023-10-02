package labelmanager

import (
	"encoding/json"
	"fmt"
)

type DeleteLabelReq struct {
	LabelId string `json:"labelid"`
}

type DeleteLabelResult struct {
	Success string `json:"success"`
}

// DeleteLabel 删除标签 https://www.showdoc.com.cn/WeChatProject/9078586560978544
func (l *LabelManagerServiceImpl) DeleteLabel(labelId string) error {

	req := DeleteLabelReq{
		LabelId: labelId,
	}

	resultBody, err := l.http.DoPost("/DeleteLabel", req)
	if err != nil {
		return err
	}

	commonResult := DeleteLabelResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
