package sendmsmanager

import (
	"encoding/json"
	"fmt"
)

type SendPicMsgReq struct {
	Wxid        string `json:"wxid"`
	PicPath     string `json:"picpath"`
	DiyFilename string `json:"diyfilename,omitempty"`
}

type SendPicMsgResult struct {
	SendPicMsg string `json:"SendPicMsg"`
}

// SendPicMsg 发送图片消息 https://www.showdoc.com.cn/WeChatProject/8929124214624404
func (l *SendMsgManagerServiceImpl) SendPicMsg(wxid, picPath, diyFilename string) error {

	req := SendPicMsgReq{
		Wxid:        wxid,
		PicPath:     picPath,
		DiyFilename: diyFilename,
	}

	resultBody, err := l.http.DoPost("/SendPicMsg", req)
	if err != nil {
		return err
	}

	commonResult := SendPicMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendPicMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
