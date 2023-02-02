package wechat

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/azure1489/wechat/model"
	"github.com/azure1489/wechat/util"
)

// OpenInlineBrowser 打开内置浏览器 https://www.showdoc.com.cn/WeChatProject/8968098266210910
func (w *Wechat) OpenInlineBrowser(urlPath, toWxid string) error {
	timeout := time.Second * 60
	client, err := util.NewClient(w.Ip, w.Port, w.Url, timeout)
	if err != nil {
		return err
	}

	req := model.OpenInlineBrowserReq{
		URL:    urlPath,
		ToWxid: toWxid,
	}

	resultBody, err := client.DoPost("/OpenInlineBrowser", req)
	if err != nil {
		return err
	}

	//fmt.Println("resultBody:", string(resultBody))

	commonResult := model.OpenInlineBrowserResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.OpenInlineBrowser != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
