package wechat

import "time"

// Wechat is a struct.
type Wechat struct {
	Ip     string
	Port   string
	Url    string
	Secret string
}

// NewWechat is a function that returns a pointer to a Wechat struct.
func NewWechat(ip, port, url, secret string) *Wechat {
	return &Wechat{
		Url:    url,
		Ip:     ip,
		Port:   port,
		Secret: secret,
	}
}

// NewWechat is a function that returns a pointer to a Wechat struct.
func NewWechatEncryption(url string) *Wechat {
	return &Wechat{
		Url: url,
	}
}

type WechatConfig struct {
	Ip      string
	Port    string
	Url     string
	Secret  string
	Timeout time.Duration
}

// NewWechat is a function that returns a pointer to a Wechat struct.
func NewWechatConfig(ip, port, url, secret string, timeout time.Duration) *WechatConfig {
	return &WechatConfig{
		Url:     url,
		Ip:      ip,
		Port:    port,
		Secret:  secret,
		Timeout: timeout,
	}
}

// NewWechat is a function that returns a pointer to a Wechat struct.
func NewWechatConfigEncryption(url string) *WechatConfig {
	return &WechatConfig{
		Url: url,
	}
}
