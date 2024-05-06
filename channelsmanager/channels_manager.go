package channelsmanager

import (
	"github.com/azure1489/wechat"
	"github.com/azure1489/wechat/common"
)

type ChannelsManagerService interface {
	// 获取视频下载地址
	GetFinderDownloadAddress(objectId string, objectNonceId string) (*GetFinderDownloadAddressRes, error)
}

type ChannelsManagerServiceImpl struct {
	config *wechat.WechatConfig
	http   common.HttpClientService
}

func NewChannelsManagerService(config *wechat.WechatConfig) ChannelsManagerService {

	httpClientService := common.NewHttpClientService(config.Ip, config.Port, config.Url, config.PublicKeyPath, config.Timeout)

	return &ChannelsManagerServiceImpl{
		config: config,
		http:   httpClientService,
	}
}
