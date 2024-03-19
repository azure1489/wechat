package receivemsgmanager

import (
	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/common"
)

type ReceiveMsgManagerService interface {
	// ConfigureMsgRecive 配置消息接收 https://www.showdoc.com.cn/WeChatProject/9204125262722344
	ConfigureMsgRecive(isEnable, url string) error
}

type ReceiveMsgManagerServiceImpl struct {
	http common.HttpClientService
}

func NewReceiveMsgManagerService(config *wechat.WechatConfig) ReceiveMsgManagerService {

	httpClientService := common.NewHttpClientService(config.Ip, config.Port, config.Url, config.PublicKeyPath, config.Timeout)

	return &ReceiveMsgManagerServiceImpl{
		// config: config,
		http: httpClientService,
	}
}
