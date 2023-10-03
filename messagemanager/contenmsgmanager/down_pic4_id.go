package contentmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// topath	是	string	图片下载到的位置(本地路径)
// aeskey	是	string	收到的图片消息中的参数 img aeskey 或 cdnthumbaeskey
// fileid	是	string	收到的图片消息中的参数 cdnthumburl 或 cdnmidimgurl
type DownPic4Req struct {
	Topath string `json:"topath"`
	Aeskey string `json:"aeskey"`
	Fileid string `json:"fileid"`
}

// 参数名	必选	类型	说明
// DownPic4ID	是	string	返回1,下载后的图片在微信缓存目录下
type DownPic4Result struct {
	DownPic4ID string `json:"DownPic4ID"`
}

// DownPic4ID 下载图片_使用文件id和key https://www.showdoc.com.cn/WeChatProject/9745105666381249
func (l *ContentMsgManagerServiceImpl) DownPic4ID(req DownPic4Req) error {

	resultBody, err := l.http.DoPost("/DownPic4ID", req)
	if err != nil {
		return err
	}

	commonResult := DownPic4Result{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.DownPic4ID != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
