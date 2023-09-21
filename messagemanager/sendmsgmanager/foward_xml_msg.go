package sendmsmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// towxid	是	string	接收方微信ID/群ID
// imgpath	是	string	封面图片路径
// xml	是	string	xml内容，只能是一行，不能有换行符
type FowardXMLMsgReq struct {
	Towxid  string `json:"towxid"`
	Imgpath string `json:"imgpath"`
	Xml     string `json:"xml"`
}

//	{
//	    "success": "1",
//	  }
type FowardXMLMsgResult struct {
	Success string `json:"success"`
}

// FowardXMLMsg 转发自定义XML消息 https://www.showdoc.com.cn/WeChatProject/9167364405161045
func (l *SendMsgManagerServiceImpl) FowardXMLMsg(req FowardXMLMsgReq) error {

	resultBody, err := l.http.DoPost("/FowardXMLMsg", req)
	if err != nil {
		return err
	}
	commonResult := FowardXMLMsgResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}
	if commonResult.Success != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}
	return nil
}
