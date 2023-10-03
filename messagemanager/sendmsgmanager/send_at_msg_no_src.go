package sendmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// gid	是	string	群ID
// wxidlist	是	string	@的对方微信ID清单，多个wxid用逗号隔开
// msg	否	string	要发送的消息内容，可以不传该参数或留空
type SendAtMsgNoSrcReq struct {
	Gid      string `json:"gid"`
	Wxidlist string `json:"wxidlist"`
	Msg      string `json:"msg,omitempty"`
}

//	{
//	    "MsgSvrID": "6769153476747201109"
//	}
type SendAtMsgNoSrcResult struct {
	MsgSvrID string `json:"MsgSvrID"`
}

// SendAtMsgNoSrc 发送群@消息_无源 https://www.showdoc.com.cn/WeChatProject/10214843464900880
func (l *SendMsgManagerServiceImpl) SendAtMsgNoSrc(req SendAtMsgNoSrcReq) (string, error) {

	resultBody, err := l.http.DoPost("/SendAtMsg_NoSrc", req)
	if err != nil {
		return "", err
	}

	commonResult := SendAtMsgNoSrcResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	if commonResult.MsgSvrID == "" {
		return "", fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return commonResult.MsgSvrID, nil
}
