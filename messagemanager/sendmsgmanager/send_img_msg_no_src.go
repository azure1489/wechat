package sendmsmanager

import (
	"encoding/json"
	"fmt"
)

type SendImgMsgNoSrcReq struct {
	Wxidorgid string `json:"wxidorgid"`
	Fileid    string `json:"fileid"`
	Authkey   string `json:"authkey"`
	Filemd5   string `json:"filemd5"`
	Filesize  string `json:"filesize"`
	Filecrc32 string `json:"filecrc32"`
}

type SendImgMsgNoSrcResult struct {
	SendImgMsgNoSrc string `json:"SendImgMsg_NoSrc"`
}

// SendImgMsgNoSrc 发送图片消息_无源,无需上传图片直接发送图片消息给对方 https://www.showdoc.com.cn/WeChatProject/9859693859851287
func (l *SendMsgManagerServiceImpl) SendImgMsgNoSrc(req SendImgMsgNoSrcReq) error {

	resultBody, err := l.http.DoPost("/SendImgMsg_NoSrc", req)
	if err != nil {
		return err
	}

	commonResult := SendImgMsgNoSrcResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendImgMsgNoSrc != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
