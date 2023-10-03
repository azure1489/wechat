package sendmsgmanager

import (
	"encoding/json"
	"fmt"
)

type SendVideoMsgReq struct {
	Wxid        string `json:"wxid"`
	VideoPath   string `json:"videopath"`
	DiyFilename string `json:"diyfilename,omitempty"`
}

type SendVideoMsgResult struct {
	SendVideoMsg string `json:"SendVideoMsg"`
}

// SendVideoMsg 发送视频消息 https://www.showdoc.com.cn/WeChatProject/8929126402335432
func (l *SendMsgManagerServiceImpl) SendVideoMsg(wxid, videoPath, diyFilename string) error {

	req := SendVideoMsgReq{
		Wxid:        wxid,
		VideoPath:   videoPath,
		DiyFilename: diyFilename,
	}

	resultBody, err := l.http.DoPost("/SendVideoMsg", req)
	if err != nil {
		return err
	}

	commonResult := SendVideoMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendVideoMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
