package wechat

type Wechat struct {
	Ip     string
	Port   string
	Url    string
	Secret string
}

func NewWechat(ip, port, url, secret string) *Wechat {
	return &Wechat{
		Url:    url,
		Ip:     ip,
		Port:   port,
		Secret: secret,
	}
}
