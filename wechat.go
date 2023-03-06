package wechat

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
