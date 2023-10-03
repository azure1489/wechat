package sendmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// towxid	是	string	接收方微信ID/群ID
// imgpath	是	string	封面图片路径
// xml	是	string	可以自己根据内容构建,或者先转发一条视频号消息，打开dbgview看内容是什么然后复制进去.
type ForwardVideoMsgReq struct {
	Towxid  string `json:"towxid"`
	Imgpath string `json:"imgpath"`
	Xml     string `json:"xml"`
}

//	{
//	    "success": "1",
//	  }
type ForwardVideoMsgResult struct {
	Success string `json:"success"`
}

// ForwardVideoMsg 转发视频号消息 https://www.showdoc.com.cn/WeChatProject/9252827740062809
func (l *SendMsgManagerServiceImpl) ForwardVideoMsg(req ForwardVideoMsgReq) error {

	resultBody, err := l.http.DoPost("/FowardXMLMsg", req)
	if err != nil {
		return err
	}
	commonResult := ForwardVideoMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}
	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}
	return nil
}
