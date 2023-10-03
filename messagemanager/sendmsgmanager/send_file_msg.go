package sendmsgmanager

import (
	"encoding/json"
	"fmt"
)

type SendFileMsgReq struct {
	Wxid        string `json:"wxid"`
	Filepath    string `json:"filepath"`
	DiyFilename string `json:"diyfilename,omitempty"`
}

type SendFileMsgResult struct {
	SendFileMsg string `json:"SendFileMsg"`
}

// SendFileMsg 发送文件消息 https://www.showdoc.com.cn/WeChatProject/8929125290388797
func (l *SendMsgManagerServiceImpl) SendFileMsg(wxid, filepath, diyFilename string) error {

	req := SendFileMsgReq{
		Wxid:        wxid,
		Filepath:    filepath,
		DiyFilename: diyFilename,
	}

	resultBody, err := l.http.DoPost("/SendFileMsg", req)
	if err != nil {
		return err
	}

	commonResult := SendFileMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.SendFileMsg != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
