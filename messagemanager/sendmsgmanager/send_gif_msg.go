package sendmsgmanager

import (
	"encoding/json"
	"fmt"
)

type SendGIFMsgReq struct {
	Wxid    string `json:"wxid"`
	GifPath string `json:"gifpath"`
}

type SendGIFMsgResult struct {
	SendGIFMsg string `json:"SendGIFMsg"`
}

// SendGIFMsg 发送GIF动画表情消息 https://www.showdoc.com.cn/WeChatProject/8929127822471500
func (l *SendMsgManagerServiceImpl) SendGIFMsg(wxid, gifPath string) error {

	req := SendGIFMsgReq{
		Wxid:    wxid,
		GifPath: gifPath,
	}

	resultBody, err := l.http.DoPost("/SendGIFMsg", req)
	if err != nil {
		return err
	}

	commonResult := SendGIFMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendGIFMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
