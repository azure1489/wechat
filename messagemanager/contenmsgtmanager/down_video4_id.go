package contentmsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// topath	是	string	视频下载到的位置(本地路径)，下载后为mp4文件
// aeskey	是	string	收到的视频消息中的参数 videomsg aeskey 或 cdnthumbaeskey
// fileid	是	string	收到的视频消息中的参数 cdnvideourl 或 cdnthumburl
type DownVideo4IDReq struct {
	Topath string `json:"topath"`
	Aeskey string `json:"aeskey"`
	Fileid string `json:"fileid"`
}

// 参数名	必选	类型	说明
// DownVideo4ID	是	string	返回1
type DownVideo4IDResult struct {
	DownVideo4ID string `json:"DownVideo4ID"`
}

// DownVideo4ID 下载视频_使用文件id和key https://www.showdoc.com.cn/WeChatProject/9819195637839857
func (l *ContentMsgManagerServiceImpl) DownVideo4ID(req DownVideo4IDReq) error {

	resultBody, err := l.http.DoPost("/DownVideo4ID", req)
	if err != nil {
		return err
	}

	commonResult := DownVideo4IDResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.DownVideo4ID != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
