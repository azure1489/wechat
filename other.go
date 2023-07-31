package wechat

import (
	"encoding/json"
	"time"

	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
)

// GetFileCrc32 取文件crc32 https://www.showdoc.com.cn/WeChatProject/10393196987937526
func (w *Wechat) GetFileCrc32(req model.GetFileCrc32Req) (string, error) {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, w.Secret, timeout)
	if err != nil {
		return "", err
	}

	resultBody, err := client.DoPost("/Get_File_crc32", req)
	if err != nil {
		return "", err
	}

	commonResult := model.GetFileCrc32Result{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return "", err
	}

	return commonResult.GetFileCrc32, nil
}
