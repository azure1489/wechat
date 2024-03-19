package collectionmanager

import (
	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/common"
)

type CollectionManagerService interface {
	// UnCollection 退还转账 https://www.showdoc.com.cn/WeChatProject/9652344490820665
	UnCollection(req UnCollectionReq) error
	// Collection 确认收款 https://www.showdoc.com.cn/WeChatProject/8929685205004781
	Collection(req CollectionReq) error
}

type CollectionManagerServiceImpl struct {
	http common.HttpClientService
}

func NewCollectionManagerService(config *wechat.WechatConfig) CollectionManagerService {

	httpClientService := common.NewHttpClientService(config.Ip, config.Port, config.Url, config.PublicKeyPath, config.Timeout)

	return &CollectionManagerServiceImpl{
		// config: config,
		http: httpClientService,
	}
}
