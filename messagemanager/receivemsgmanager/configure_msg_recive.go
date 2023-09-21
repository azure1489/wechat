package receivemsgmanager

import (
	"encoding/json"
	"fmt"
)

// 参数名	必选	类型	说明
// isEnable	是	string	0=关
// 1=开(接收处理后的PB)
// 2=开(接收原始PB)
// 3=指定好友朋友圈异步数据
// 4=指定好友朋友圈异步PB数据
// url	是	string	消息回调URL地址
type ConfigureMsgReciveReq struct {
	IsEnable string `json:"isEnable"`
	URL      string `json:"url"`
}

// 参数名	必选	类型	说明
// ConfigureMsgRecive	是	string	返回配置情况状态
// 1=配置成功
// 0=配置失败
type ConfigureMsgReciveResult struct {
	ConfigureMsgRecive string `json:"ConfigureMsgRecive"`
}

// ConfigureMsgRecive 配置消息接收 https://www.showdoc.com.cn/WeChatProject/10273567687375653
func (l *ReceiveMsgManagerServiceImpl) ConfigureMsgRecive(req ConfigureMsgReciveReq) error {

	resultBody, err := l.http.DoPost("/ConfigureMsgRecive", req)
	if err != nil {
		return err
	}

	commonResult := ConfigureMsgReciveResult{}
	err = json.Unmarshal(resultBody, &commonResult)
	if err != nil {
		return err
	}

	if commonResult.ConfigureMsgRecive != "1" {
		return fmt.Errorf("提交失败, body=%s", string(resultBody))
	}

	return nil
}
