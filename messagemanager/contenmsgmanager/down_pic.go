package contentmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// topath	是	string	图片下载到的位置(本地路径)
// msg_xml	是	string	图片消息的完整xml内容
type DownPicReq struct {
	Topath string `json:"topath"`
	MsgXml string `json:"msg_xml"`
}

// 参数名	必选	类型	说明
// DownPic	是	string	返回1,下载后的图片在微信缓存目录下
type DownPicResult struct {
	DownPic string `json:"DownPic"`
}

// DownPic 下载图片_使用消息内容 https://www.showdoc.com.cn/WeChatProject/9682774769136221
func (l *ContentMsgManagerServiceImpl) DownPic(topath, msgXml string) error {

	req := DownPicReq{
		Topath: topath,
		MsgXml: msgXml,
	}

	resultBody, err := l.http.DoPost("/DownPic", req)
	if err != nil {
		return err
	}

	commonResult := DownPicResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.DownPic != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
