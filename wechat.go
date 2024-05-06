package wechat

import (
	"time"

	"github.com/azure1489/wechat/callbackevent/server"
)

// Wechat is a struct.
// type Wechat struct {
// 	Ip            string
// 	Port          string
// 	Url           string
// 	PublicKeyPath string
// }

// NewWechat is a function that returns a pointer to a Wechat struct.
// func NewWechat(ip, port, url, publicKeyPath string) *Wechat {
// 	return &Wechat{
// 		Url:           url,
// 		Ip:            ip,
// 		Port:          port,
// 		PublicKeyPath: publicKeyPath,
// 	}
// }

// // NewWechat is a function that returns a pointer to a Wechat struct.
// func NewWechatEncryption(url string) *Wechat {
// 	return &Wechat{
// 		Url: url,
// 	}
// }

type WechatConfig struct {
	Ip            string
	Port          string
	Url           string
	PublicKeyPath string
	// ControllerIp   string
	// ControllerPort string
	Timeout time.Duration
}

var wechatConfigInstance *WechatConfig

// NewWechat is a function that returns a pointer to a Wechat struct.
func NewWechatConfig(ip, port, url, publicKeyPath string, timeout time.Duration) *WechatConfig {
	if wechatConfigInstance == nil {
		wechatConfigInstance = &WechatConfig{
			Url:           url,
			Ip:            ip,
			Port:          port,
			PublicKeyPath: publicKeyPath,
			Timeout:       timeout,
		}
	}
	return wechatConfigInstance
}

// func NewControllerWechatConfig(controllerIp, controllerPort, url, publicKeyPath string, timeout time.Duration) *WechatConfig {
// 	if wechatConfigInstance == nil {
// 		wechatConfigInstance = &WechatConfig{
// 			Url:            url,
// 			ControllerIp:   controllerIp,
// 			ControllerPort: controllerPort,
// 			PublicKeyPath:  publicKeyPath,
// 			Timeout:        timeout,
// 		}
// 	}
// 	return wechatConfigInstance
// }

// NewWechat is a function that returns a pointer to a Wechat struct.
// func NewWechatConfigEncryption(url string) *WechatConfig {
// 	return &WechatConfig{
// 		Url: url,
// 	}
// }

// GetOfficialAccount 获取微信公众号实例
func GetCallBackServer(body []byte) *server.Server {
	return server.NewServer(body)
}
